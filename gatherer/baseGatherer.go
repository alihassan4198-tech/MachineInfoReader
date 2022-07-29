package gatherer

import (
	"errors"
	"fmt"
	"machine_info_gatherer/common"
	"machine_info_gatherer/distro"
	"machine_info_gatherer/errorslist"
	"machine_info_gatherer/model"
	"os"
	"os/exec"
	"strings"

	"github.com/alihassan4198-tech/ghw"
	"github.com/coreos/go-iptables/iptables"
	"github.com/zcalusic/sysinfo"
)

type BaseGatherer struct {
}

//--------------------------
// Baseboard Implementation
//--------------------------

func (bg *BaseGatherer) GetComputerBaseboard() *model.ComputerBaseboardType {
	// Get Computer Softwares Installed Distro Wise
	currentDistro := distro.GetInstance()
	comBaseBoard, err := currentDistro.GetComputerBaseboard()
	if err != nil {
		fmt.Println(err)
	}

	return comBaseBoard
}

func (bg *BaseGatherer) GetComputerBios() *model.ComputerBiosType {
	cbios := model.ComputerBiosType{}
	currentDistro := distro.GetInstance()
	cbios, err := currentDistro.GetComputerBaseboard()
	if err != nil {
		fmt.Println(err)
	}

	return &cbios
}

func (bg *BaseGatherer) GetComputerCPU() *model.ComputerCPUType {
	compCpu := model.ComputerCPUType{}

	cpu, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}

	compCpu.Caption = cpuCaption

	for _, proc := range cpu.Processors {
		c := model.ComputerCPU{}

		c.Manufacturer = proc.Vendor
		c.Max_clock_speed = proc.CPUMhz
		c.Name = proc.Name
		c.Device_id = proc.ID

		compCpu.CPUCores = append(compCpu.CPUCores, c)
	}

	return &compCpu
}

func (bg *BaseGatherer) GetComputerEndpointProtectionSoftwares() *model.ComputerEndpointProtectionType {
	epsoft := model.ComputerEndpointProtectionType{}
	// soft , err := ghw.

	return &epsoft
}

func (bg *BaseGatherer) GetComputerFirewallRules() *model.ComputerFirewallRulesType {

	cfwRules := model.ComputerFirewallRulesType{}

	tables := []string{"filter", "mangle", "nat", "raw"}

	ipt, err := iptables.New()
	if err != nil {
		fmt.Println(err)
	}

	for _, table := range tables {
		chains, err := ipt.ListChains(table)
		if err != nil {
			fmt.Println(err)
		}
		for _, chain := range chains {
			structStat, err := ipt.StructuredStats(table, chain)
			if err != nil {
				fmt.Println("parsing error: ", err)
				continue
			}
			for _, s := range structStat {
				cfwRule := model.FirewallRuleType{}

				cfwRule.TableName = table
				cfwRule.ChainName = chain
				cfwRule.Action = s.Target
				cfwRule.Enabled = "enable"
				cfwRule.Local_ip = s.Source.IP.String()
				cfwRule.Local_port = s.Input
				cfwRule.Remote_ip = s.Destination.IP.String()
				cfwRule.Remote_port = s.Output
				cfwRule.Protocol = s.Protocol
				cfwRule.Direction = chain
				// fmt.Printf("%#v\n", cfwRule)
				cfwRules.FW_rules = append(cfwRules.FW_rules, cfwRule)
			}
		}
	}

	return &cfwRules
}

func (bg *BaseGatherer) GetComputerNIC() *[]model.ComputerNICType {

	comNic := []model.ComputerNICType{}
	net, err := ghw.Network()
	if err != nil {
		fmt.Printf("Error getting network info: %v", err)
	}

	for _, nic := range net.NICs {

		comNic = append(comNic, model.ComputerNICType{
			Caption:     nic.Name,
			Mac_address: nic.MacAddress,
		})

	}

	return &comNic
}

func (bg *BaseGatherer) GetComputerOS() *model.ComputerOSType {

	comOS := model.ComputerOSType{}
	var si sysinfo.SysInfo
	si.GetSysInfo()
	Os := si.OS
	cname, err := os.Hostname()
	if err != nil {
		fmt.Printf("error:%#v", err)
	}

	comOS.Computer_name = cname
	m := *(common.ReadOSRelease())

	comOS.Os_version = m["VERSION_CODENAME"] + " " + m["VERSION_ID"]
	comOS.Os_name = Os.Name
	comOS.Vendor = Os.Vendor
	comOS.Caption = osCaption
	comOS.Os_architecture = Os.Architecture
	comOS.Os_version = Os.Version
	comOS.Release = Os.Release
	comOS.Lts = false

	return &comOS
}

func (bg *BaseGatherer) GetComputerServices() *model.ComputerServicesType {

	comServ := model.ComputerServicesType{}

	cmd, err := exec.Command("systemctl", "--type=service").Output()
	if err != nil {
		fmt.Println(err)
	}
	cmdOutput := strings.Split(string(cmd), "\n")

	for _, svc := range cmdOutput {
		if common.IsService(svc) && common.IsServiceRunning(svc) {
			comServ.Services = append(comServ.Services, *common.ParseService(svc))
		}
	}

	comServ.TotalServciesRunning = len(comServ.Services)

	return &comServ
}

func (bg *BaseGatherer) GetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalledType, error) {
	comInsSoft := model.ComputerSoftwaresInstalledType{}
	return &comInsSoft, errors.New(errorslist.ErrNotImplemented)
}

func (bg *BaseGatherer) GetDistroBasedComputerSoftwareInstalled() *model.ComputerSoftwaresInstalledType {
	// Get Computer Softwares Installed Distro Wise
	currentDistro := distro.GetInstance()
	comSoftInstall, err := currentDistro.GetComputerSoftwaresInstalled()
	if err != nil {
		fmt.Println(err)
	}

	return comSoftInstall
}

func (bg *BaseGatherer) GetComputerSystem() *model.ComputerSystemType {

	comSys := model.ComputerSystemType{}

	domainName, err := exec.Command("domainname").Output()
	if err != nil {
		fmt.Println(err)
	}
	comSys.Domain = strings.TrimSpace(string(domainName))

	manufacturer, err := common.RunFullCommandWithSudo("dmidecode -s baseboard-manufacturer")
	if err != nil {
		fmt.Println(err)
		comSys.Manufacturer = common.NeedSudoPreviliges(err)
	} else {
		comSys.Manufacturer = manufacturer
	}

	model, err := common.RunFullCommandWithSudo("dmidecode -s baseboard-product-name")
	if err != nil {
		fmt.Println(err)
		comSys.Model = common.NeedSudoPreviliges(err)
	} else {
		comSys.Model = model
	}

	return &comSys
}

func (bg *BaseGatherer) GetComputerPatches() *model.ComputerPatchesType {
	comPatch := model.ComputerPatchesType{}

	return &comPatch
}

//  All Info
func (bg *BaseGatherer) GatherInfo() *model.ComputerInfoType {

	m := model.ComputerInfoType{}

	m.ComputerBaseboard = *(bg.GetComputerBaseboard())
	m.ComputerBios = *(bg.GetComputerBios())
	m.ComputerCPU = *(bg.GetComputerCPU())
	m.ComputerEndpointProtection = *(bg.GetComputerEndpointProtectionSoftwares())
	m.ComputerFirewallRules = *(bg.GetComputerFirewallRules())
	m.ComputerNICS = *(bg.GetComputerNIC())
	m.ComputerOS = *(bg.GetComputerOS())
	m.ComputerServices = *(bg.GetComputerServices())
	m.ComputerSoftwaresInstalled = *(bg.GetDistroBasedComputerSoftwareInstalled())
	m.ComputerSystem = *(bg.GetComputerSystem())
	m.ComputerPatches = *(bg.GetComputerPatches())

	return &m
}
