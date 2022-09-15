package gatherer

import (
	"fmt"
	"machine_info_gatherer/debug"
	"machine_info_gatherer/distro"
	"machine_info_gatherer/model"
	"reflect"
	"strings"
	"time"
)

type BaseGatherer struct {
}

//--------------------------
// Baseboard Implementation
//--------------------------

func (bg *BaseGatherer) GetComputerBaseboard() *model.ComputerBaseboardType {
	// Get Computer Softwares Installed Distro Wise
	defer debug.TimeTrack(time.Now(), debug.FileFunctionLine())

	currentDistro := distro.GetInstance()
	comBaseBoard, err := currentDistro.DistroGetComputerBaseboard()
	if err != nil {
		fmt.Println(err)
	}
	return comBaseBoard
}

func (bg *BaseGatherer) GetComputerBios() *model.ComputerBiosType {
	defer debug.TimeTrack(time.Now(), debug.FileFunctionLine())
	currentDistro := distro.GetInstance()
	cbios, err := currentDistro.DistroGetComputerBios()
	if err != nil {
		fmt.Println(err)
	}
	return cbios
}

func (bg *BaseGatherer) GetComputerCPU() *model.ComputerCPUType {
	defer debug.TimeTrack(time.Now(), debug.FileFunctionLine())
	currentDistro := distro.GetInstance()
	compCpu, err := currentDistro.DistroGetComputerCPU()
	if err != nil {
		fmt.Println(err)
	}
	return compCpu
}

func (bg *BaseGatherer) GetComputerEndpointProtectionSoftwares() *model.ComputerEndpointProtectionType {
	defer debug.TimeTrack(time.Now(), debug.FileFunctionLine())
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
	fmt.Println("We are in baseGather.go file")
	m.ComputerBaseboard = *(bg.GetComputerBaseboard())
	fmt.Println("GetComputerBaseboard Finished")
	m.ComputerBios = *(bg.GetComputerBios())
	fmt.Println("GetComputerBios Finished")
	m.ComputerCPU = *(bg.GetComputerCPU())
	fmt.Println("GetComputerCPU Finished")
	m.ComputerEndpointProtection = *(bg.GetComputerEndpointProtectionSoftwares())
	fmt.Println("GetComputerEndPoint Finished")
	m.ComputerFirewallRules = *(bg.GetComputerFirewallRules())
	fmt.Println("GetComputerFirewall Finished")
	m.ComputerNICS = *(bg.GetComputerNIC())
	fmt.Println("GetComputerNIC Finished")
	m.ComputerOS = *(bg.GetComputerOS())
	fmt.Println("GetComputerOS Finished")
	m.ComputerServices = *(bg.GetComputerServices())
	fmt.Println("GetComputerServices Finished")
	m.ComputerSoftwaresInstalled = *(bg.GetComputerSoftwaresInstalled())
	fmt.Println("GetComputerSoftwareInstalled Finished")
	m.ComputerSystem = *(bg.GetComputerSystem())
	fmt.Println("GetComputerSystem Finished")
	m.ComputerPatches = *(bg.GetComputerPatches())
	fmt.Println("GetComputerPatches Finished")

	fmt.Println("baseGather.go file Finished")
	return &m
}
