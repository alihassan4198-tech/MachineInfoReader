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
)

type BaseGatherer struct {
}

//--------------------------
// Baseboard Implementation
//--------------------------

const (
	baseboardCaption string = "Base Board"
	cpuCaption       string = "CPU"
	osCaption        string = "Computer OS"
	biosCaption      string = "Computer Bios"
)

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

	bios, err := ghw.BIOS()
	if err != nil {
		fmt.Printf("Error getting BIOS info: %v", err)
	}

	cbios.Name = bios.Name
	cbios.Biosversion = common.RootNeeded(bios.Version)
	cbios.Version = cbios.Biosversion
	cbios.Manufacturer = common.RootNeeded(bios.Vendor)
	cbios.Installdate = common.RootNeeded(bios.Date)
	cbios.Serialnumber = common.RootNeeded(bios.Serialnumber)
	cbios.Installdate = common.RootNeeded(bios.Installdate)
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

	return &epsoft
}

func (bg *BaseGatherer) GetComputerFirewallRules() *model.ComputerFirewallRulesType {

	cfwRules := model.ComputerFirewallRulesType{}

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

	cname, err := os.Hostname()
	if err != nil {
		fmt.Printf("error:%#v", err)
	}

	comOS.Computer_name = cname
	comOS.Caption = osCaption

	m := *(common.ReadOSRelease())

	comOS.Os_version = m["VERSION_CODENAME"] + " " + m["VERSION_ID"]
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
