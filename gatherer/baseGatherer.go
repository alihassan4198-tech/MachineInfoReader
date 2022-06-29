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
)

func (bg *BaseGatherer) GetComputerBaseboard() *model.ComputerBaseboard {
	cb := model.ComputerBaseboard{}

	var err error
	computerName, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	cb.Computer_name = computerName
	cb.Caption = baseboardCaption

	baseboard, err := ghw.Baseboard()
	if err != nil {
		fmt.Printf("baseboard error : %s", err)
	}

	cb.Serialnumber = common.RootNeeded(baseboard.SerialNumber)
	cb.Version = common.RootNeeded(baseboard.Version)
	cb.Product = common.RootNeeded(baseboard.Product)
	cb.Manufacturer = common.RootNeeded(baseboard.Vendor)
	cb.Tag = common.RootNeeded(baseboard.AssetTag)

	return &cb
}

func (bg *BaseGatherer) GetComputerBios() *model.ComputerBios {
	cbios := model.ComputerBios{}

	bios, err := ghw.BIOS()
	if err != nil {
		fmt.Printf("Error getting BIOS info: %v", err)
	}

	cbios.Biosversion = common.RootNeeded(bios.Version)
	cbios.Manufacturer = common.RootNeeded(bios.Vendor)
	cbios.Installdate = common.RootNeeded(bios.Date)

	return &cbios
}

func (bg *BaseGatherer) GetComputerCPU() *model.ComputerCPU {
	compCpu := model.ComputerCPU{}

	_, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}

	compCpu.Caption = cpuCaption

	return &compCpu
}

func (bg *BaseGatherer) GetComputerEndpointProtectionSoftwares() *model.ComputerEndpointProtection {

	epsoft := model.ComputerEndpointProtection{}

	return &epsoft
}

func (bg *BaseGatherer) GetComputerFirewallRules() *model.ComputerFirewallRules {

	cfwRules := model.ComputerFirewallRules{}

	return &cfwRules
}

func (bg *BaseGatherer) GetComputerNIC() *[]model.ComputerNIC {

	comNic := []model.ComputerNIC{}
	net, err := ghw.Network()
	if err != nil {
		fmt.Printf("Error getting network info: %v", err)
	}

	for _, nic := range net.NICs {

		comNic = append(comNic, model.ComputerNIC{
			Caption:     nic.Name,
			Mac_address: nic.MacAddress,
		})

	}

	return &comNic
}

func (bg *BaseGatherer) GetComputerOS() *model.ComputerOS {

	comOS := model.ComputerOS{}

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

func (bg *BaseGatherer) GetComputerServices() *model.ComputerServices {

	comServ := model.ComputerServices{}

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

func (bg *BaseGatherer) GetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalled, error) {
	comInsSoft := model.ComputerSoftwaresInstalled{}

	return &comInsSoft, errors.New(errorslist.ErrNotImplemented)
}

//  All Info
func (bg *BaseGatherer) GatherInfo() *model.ComputerInfo {

	m := model.ComputerInfo{}

	m.ComputerBaseboard = *(bg.GetComputerBaseboard())
	m.ComputerBios = *(bg.GetComputerBios())
	m.ComputerCPU = *(bg.GetComputerCPU())
	m.ComputerEndpointProtection = *(bg.GetComputerEndpointProtectionSoftwares())
	m.ComputerFirewallRules = *(bg.GetComputerFirewallRules())
	m.ComputerNICS = *(bg.GetComputerNIC())
	m.ComputerOS = *(bg.GetComputerOS())
	m.ComputerServices = *(bg.GetComputerServices())

	// Get Computer Softwares Installed Distro Wise
	currentDistro := distro.GetInstance()
	comSoftInstall, err := currentDistro.GetComputerSoftwaresInstalled()
	if err != nil {
		fmt.Println(err)
	} else {
		m.ComputerSoftwaresInstalled = *comSoftInstall
	}

	return &m
}
