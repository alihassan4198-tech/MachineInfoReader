//go:build darwin
// +build darwin

package distro

import (
	"fmt"
	"machine_info_gatherer/common"
	"machine_info_gatherer/distro/systemprofiler"
	"machine_info_gatherer/model"
	"os"
	"os/exec"
	"strings"

	"github.com/alihassan4198-tech/ghw"
	"github.com/coreos/go-iptables/iptables"
)

// Mac

const (
	find0 = "Name"
	find1 = "Version"
	find2 = "Architecture"
	find3 = "Description"
)

type MacBased struct {
	LinuxBase
}

var infoMap systemprofiler.DarwinSystemProfilerInfo

func MacGetAllInfoInMap() {
	infoMap = systemprofiler.DarwinSystemProfiler()
}

func (mb *MacBased) DistroGatherInfo() (*model.ComputerInfoType, error) {

	return &model.ComputerInfoType{}, nil
}

func (mb *MacBased) DistroGetComputerBaseboard() (*model.ComputerBaseboardType, error) {
	cb := model.ComputerBaseboardType{}

	hw := len(infoMap.SPHardware.SpHardwareDataType) > 0
	sw := len(infoMap.SPSoftware.SpSoftwareDataType) > 0

	var err error
	computerName, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}

	cb.Computer_name = computerName
	cb.Caption = baseboardCaption

	cb.Creationclassname = ""

	if hw {
		cb.Serialnumber = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].SerialNumber)
		cb.Manufacturer = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].MachineName)
		cb.Tag = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].PlatformUuid)
		cb.Model = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].MachineModel)
	}

	if sw {
		cb.Version = common.RootNeeded(infoMap.SPSoftware.SpSoftwareDataType[0].OSVersion)
		cb.Product = common.RootNeeded(infoMap.SPSoftware.SpSoftwareDataType[0].KernelVersion)
		cb.Name = common.RootNeeded(infoMap.SPSoftware.SpSoftwareDataType[0].LocalHostName)
		cb.Status = common.RootNeeded(infoMap.SPSoftware.SpSoftwareDataType[0].BootMode)
	}

	cb.Creationclassname = ""

	date, err := common.RunFullCommand("ls -l /var/db/.AppleSetupDone")
	if err != nil {
		date = ""
	} else {
		split := strings.Split(date, " ")
		date = strings.Join(split[8:12], " ")
	}

	cb.Installdate = date

	cb.Description = "Baseboard Info"

	cb.Poweredon = true

	return &cb, nil
}

func (mb *MacBased) DistroGetComputerBios() (*model.ComputerBiosType, error) {
	cbios := model.ComputerBiosType{}

	hw := len(infoMap.SPHardware.SpHardwareDataType) > 0
	sw := len(infoMap.SPSoftware.SpSoftwareDataType) > 0

	if hw {

	}

	if sw {

	}

	cbios.Name = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].Name)
	cbios.Biosversion = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].BootRomVersion)
	cbios.Version = cbios.Biosversion
	cbios.Manufacturer = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].MachineName)
	cbios.Installdate = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].Date)
	cbios.Serialnumber = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].SerialNumber)
	cbios.Primarybios = true
	cbios.Caption = biosCaption
	maj, min, found := strings.Cut(cbios.Biosversion, ".")

	if found {
		cbios.Smbiosmajorversion = maj
		cbios.Smbiosminorversion = min
	}

	if cbios.Biosversion != "" {
		cbios.Status = "Installed"
	}

	return &cbios, nil
}

func (mb *MacBased) DistroGetComputerCPU() (*model.ComputerCPUType, error) {
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

	return &compCpu, nil
}

func (mb *MacBased) DistroGetComputerEndpointProtectionSoftwares() (*model.ComputerEndpointProtectionType, error) {
	epsoft := model.ComputerEndpointProtectionType{}
	// soft , err := ghw.

	return &epsoft, nil
}

func (mb *MacBased) DistroGetComputerFirewallRules() (*model.ComputerFirewallRulesType, error) {

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
				cfwRule := model.FirewallRule{}

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

	return &cfwRules, nil
}

func (mb *MacBased) DistroGetComputerNIC() (*[]model.ComputerNICType, error) {

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

	return &comNic, nil
}

func (mb *MacBased) DistroGetComputerOS() (*model.ComputerOSType, error) {

	comOS := model.ComputerOSType{}
	// var si sysinfo.SysInfo
	// si.GetSysInfo()
	// Os := si.OS
	// cname, err := os.Hostname()
	// if err != nil {
	// 	fmt.Printf("error:%#v", err)
	// }

	// comOS.Computer_name = cname
	// m := *(common.ReadOSRelease())

	// comOS.Os_version = m["VERSION_CODENAME"] + " " + m["VERSION_ID"]
	// comOS.Os_name = Os.Name
	// comOS.Vendor = Os.Vendor
	// comOS.Caption = osCaption
	// comOS.Os_architecture = Os.Architecture
	// comOS.Os_version = Os.Version
	// comOS.Release = Os.Release
	// comOS.Lts = false

	return &comOS, nil
}

func (mb *MacBased) DistroGetComputerServices() (*model.ComputerServicesType, error) {

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

	return &comServ, nil
}

func (mb *MacBased) DistroGetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalledType, error) {
	// Get Computer Softwares Installed Distro Wise

	return &model.ComputerSoftwaresInstalledType{}, nil
}

func (mb *MacBased) DistroGetComputerSystem() (*model.ComputerSystemType, error) {

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

	return &comSys, nil
}

func (mb *MacBased) DistroGetComputerPatches() (*model.ComputerPatchesType, error) {
	comPatch := model.ComputerPatchesType{}

	return &comPatch, nil
}

func (mb *MacBased) GetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalledType, error) {
	comSoftInst := model.ComputerSoftwaresInstalledType{}

	installSoft, err := exec.Command("dpkg", "-l").Output()
	if err != nil {
		fmt.Println(err)
		return &comSoftInst, err
	}
	splittedInstallSoft := strings.Split(string(installSoft), "\n")

	for i, soft := range splittedInstallSoft {
		if strings.Contains(soft, find0) &&
			strings.Contains(soft, find1) &&
			strings.Contains(soft, find2) &&
			strings.Contains(soft, find3) {
			// Remove first 2 useless record strings
			splittedInstallSoft = splittedInstallSoft[i+2:]
			break
		}
	}

	// Remove last \n record string
	splittedInstallSoft = splittedInstallSoft[:len(splittedInstallSoft)-1]

	// Convert to Strcut
	for _, soft := range splittedInstallSoft {
		soft = strings.Join(strings.Fields(soft), " ")
		splittedSoft := strings.Split(soft, " ")
		softInstall := model.SoftwareInstalledType{}
		softInstall.Display_name = splittedSoft[1]
		softInstall.Version = splittedSoft[2]

		comSoftInst.SoftwaresInstalled = append(comSoftInst.SoftwaresInstalled, softInstall)
	}

	comSoftInst.Total_software = len(comSoftInst.SoftwaresInstalled)

	return &comSoftInst, nil
}
