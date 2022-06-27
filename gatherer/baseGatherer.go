package gatherer

import (
	"fmt"
	"machine_info_gatherer/model"
	"os"

	"github.com/jaypipes/ghw"
)

type BaseGatherer struct {
}

//--------------------------
// Baseboard Implementation
//--------------------------

const (
	baseboardCaption string = "Base Board"
	cpuCaption       string = "CPU"
)

func rootNeeded(arg string) string {
	if arg == "None" || arg == "unknown" {
		return arg + " (need root access)"
	} else {
		return arg
	}
}

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

	cb.Serialnumber = rootNeeded(baseboard.SerialNumber)
	cb.Version = rootNeeded(baseboard.Version)
	cb.Product = rootNeeded(baseboard.Product)
	cb.Manufacturer = rootNeeded(baseboard.Vendor)
	cb.Tag = rootNeeded(baseboard.AssetTag)

	return &cb
}

func (bg *BaseGatherer) GetComputerBios() *model.ComputerBios {
	cbios := model.ComputerBios{}

	bios, err := ghw.BIOS()
	if err != nil {
		fmt.Printf("Error getting BIOS info: %v", err)
	}

	cbios.Biosversion = rootNeeded(bios.Version)
	cbios.Manufacturer = rootNeeded(bios.Vendor)
	cbios.Installdate = rootNeeded(bios.Date)

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

//  All Info
func (bg *BaseGatherer) GatherInfo() *model.ComputerInfo {

	m := model.ComputerInfo{}

	m.ComputerBaseboard = *(bg.GetComputerBaseboard())
	m.ComputerBios = *(bg.GetComputerBios())
	m.ComputerCPU = *(bg.GetComputerCPU())

	return &m
}
