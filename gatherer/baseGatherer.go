package gatherer

import (
	"fmt"
	"machine_info_gatherer/distro"
	"machine_info_gatherer/model"
	"reflect"
	"strings"
)

type BaseGatherer struct {
}

//--------------------------
// Baseboard Implementation
//--------------------------

func (bg *BaseGatherer) GetComputerBaseboard() *model.ComputerBaseboardType {
	// Get Computer Softwares Installed Distro Wise
	currentDistro := distro.GetInstance()
	comBaseBoard, err := currentDistro.DistroGetComputerBaseboard()
	if err != nil {
		fmt.Println(err)
	}

	return comBaseBoard
}

func (bg *BaseGatherer) GetComputerBios() *model.ComputerBiosType {
	currentDistro := distro.GetInstance()
	cbios, err := currentDistro.DistroGetComputerBios()
	if err != nil {
		fmt.Println(err)
	}

	return cbios
}

func (bg *BaseGatherer) GetComputerCPU() *model.ComputerCPUType {
	currentDistro := distro.GetInstance()
	compCpu, err := currentDistro.DistroGetComputerCPU()
	if err != nil {
		fmt.Println(err)
	}

	return compCpu
}

func (bg *BaseGatherer) GetComputerEndpointProtectionSoftwares() *model.ComputerEndpointProtectionType {
	currentDistro := distro.GetInstance()
	epsoft, err := currentDistro.DistroGetComputerEndpointProtectionSoftwares()
	if err != nil {
		fmt.Println(err)
	}

	return epsoft
}

func (bg *BaseGatherer) GetComputerFirewallRules() *model.ComputerFirewallRulesType {

	currentDistro := distro.GetInstance()
	cfwRules, err := currentDistro.DistroGetComputerFirewallRules()
	if err != nil {
		fmt.Println(err)
	}

	return cfwRules
}

func (bg *BaseGatherer) GetComputerNIC() *[]model.ComputerNICType {

	currentDistro := distro.GetInstance()
	comNic, err := currentDistro.DistroGetComputerNIC()
	if err != nil {
		fmt.Println(err)
	}

	return comNic
}

func (bg *BaseGatherer) GetComputerOS() *model.ComputerOSType {

	currentDistro := distro.GetInstance()
	comOS, err := currentDistro.DistroGetComputerOS()
	if err != nil {
		fmt.Println(err)
	}

	return comOS
}

func (bg *BaseGatherer) GetComputerServices() *model.ComputerServicesType {

	currentDistro := distro.GetInstance()
	comServ, err := currentDistro.DistroGetComputerServices()
	if err != nil {
		fmt.Println(err)
	}

	return comServ
}

func (bg *BaseGatherer) GetComputerSoftwaresInstalled() *model.ComputerSoftwaresInstalledType {

	currentDistro := distro.GetInstance()
	comInsSoft, err := currentDistro.DistroGetComputerSoftwaresInstalled()
	if err != nil {
		fmt.Println(err)
	}

	return comInsSoft
}

func (bg *BaseGatherer) GetComputerSystem() *model.ComputerSystemType {

	currentDistro := distro.GetInstance()
	comSys, err := currentDistro.DistroGetComputerSystem()
	if err != nil {
		fmt.Println(err)
	}

	return comSys
}

func (bg *BaseGatherer) GetComputerPatches() *model.ComputerPatchesType {
	currentDistro := distro.GetInstance()
	comPatch, err := currentDistro.DistroGetComputerPatches()
	if err != nil {
		fmt.Println(err)
	}

	return comPatch
}

//  All Info
func (bg *BaseGatherer) GatherInfo() *model.ComputerInfoType {

	m := model.ComputerInfoType{}

	currentDistro := distro.GetInstance()
	if strings.Contains(reflect.TypeOf(currentDistro).String(), "MacBased") {
		distro.MacGetAllInfoInMap()
	}

	m.ComputerBaseboard = *(bg.GetComputerBaseboard())
	m.ComputerBios = *(bg.GetComputerBios())
	m.ComputerCPU = *(bg.GetComputerCPU())
	m.ComputerEndpointProtection = *(bg.GetComputerEndpointProtectionSoftwares())
	m.ComputerFirewallRules = *(bg.GetComputerFirewallRules())
	m.ComputerNICS = *(bg.GetComputerNIC())
	m.ComputerOS = *(bg.GetComputerOS())
	m.ComputerServices = *(bg.GetComputerServices())
	m.ComputerSoftwaresInstalled = *(bg.GetComputerSoftwaresInstalled())
	m.ComputerSystem = *(bg.GetComputerSystem())
	m.ComputerPatches = *(bg.GetComputerPatches())

	return &m
}
