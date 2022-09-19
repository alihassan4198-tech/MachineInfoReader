package gatherer

import (
	"fmt"
	// "machine_info_gatherer/debug"
	"machine_info_gatherer/distro"
	"machine_info_gatherer/model"
	"reflect"
	"strings"
	// "time"
)

type BaseGatherer struct {
}

//--------------------------
// Baseboard Implementation
//--------------------------

func (bg *BaseGatherer) GetComputerBaseboard() *model.ComputerBaseboardType {
	// Get Computer Softwares Installed Distro Wise
	// defer debug.TimeTrack(time.Now(), debug.FileFunctionLine())

	// fmt.Println("in GetComputerBaseboard start")//ok

	currentDistro := distro.GetInstance()
	fmt.Println(" in GetComputerBaseboard check data : ", currentDistro) //ok

	comBaseBoard, err := currentDistro.DistroGetComputerBaseboard()
	fmt.Println(" in GetComputerBaseboard check data : ", comBaseBoard) //ok
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("in GetComputerBaseboard end")

	return comBaseBoard
}

func (bg *BaseGatherer) GetComputerBios() *model.ComputerBiosType {
	// defer debug.TimeTrack(time.Now(), debug.FileFunctionLine())

	fmt.Println("in GetComputerBios start")
	currentDistro := distro.GetInstance()
	fmt.Println(" in GetComputerBios  check data current distro : ", currentDistro) //ok
	cbios, err := currentDistro.DistroGetComputerBios()
	fmt.Println(" in GetComputerBios  check data : ", cbios)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("in GetComputerBios end")
	return cbios
}

func (bg *BaseGatherer) GetComputerCPU() *model.ComputerCPUType {
	// defer debug.TimeTrack(time.Now(), debug.FileFunctionLine())
	fmt.Println("in GetComputerCPU start")
	currentDistro := distro.GetInstance()
	compCpu, err := currentDistro.DistroGetComputerCPU()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("in GetComputerCPU end")
	return compCpu
}

func (bg *BaseGatherer) GetComputerEndpointProtectionSoftwares() *model.ComputerEndpointProtectionType {
	// defer debug.TimeTrack(time.Now(), debug.FileFunctionLine())
	fmt.Println("in GetComputerEndpointProtectionSoftwares start")
	currentDistro := distro.GetInstance()
	epsoft, err := currentDistro.DistroGetComputerEndpointProtectionSoftwares()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("in GetComputerEndpointProtectionSoftwares end")
	return epsoft
}

func (bg *BaseGatherer) GetComputerFirewallRules() *model.ComputerFirewallRulesType {
	fmt.Println("in GetComputerFirewallRules start")
	currentDistro := distro.GetInstance()
	cfwRules, err := currentDistro.DistroGetComputerFirewallRules()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("in GetComputerFirewallRules end")
	return cfwRules
}

func (bg *BaseGatherer) GetComputerNIC() *[]model.ComputerNICType {
	fmt.Println("in GetComputerNIC start")
	currentDistro := distro.GetInstance()
	comNic, err := currentDistro.DistroGetComputerNIC()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("in GetComputerNIC end")
	return comNic
}

func (bg *BaseGatherer) GetComputerOS() *model.ComputerOSType {
	fmt.Println("in GetComputerOS start")
	currentDistro := distro.GetInstance()
	comOS, err := currentDistro.DistroGetComputerOS()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("in GetComputerOS end")
	return comOS
}

func (bg *BaseGatherer) GetComputerServices() *model.ComputerServicesType {
	fmt.Println("in GetComputerServices start")
	currentDistro := distro.GetInstance()
	comServ, err := currentDistro.DistroGetComputerServices()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("in GetComputerServices end")
	return comServ
}

func (bg *BaseGatherer) GetComputerSoftwaresInstalled() *model.ComputerSoftwaresInstalledType {
	fmt.Println("in GetComputerSoftwaresInstalled start")
	currentDistro := distro.GetInstance()
	comInsSoft, err := currentDistro.DistroGetComputerSoftwaresInstalled()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("in GetComputerSoftwaresInstalled end")
	return comInsSoft
}

func (bg *BaseGatherer) GetComputerSystem() *model.ComputerSystemType {
	fmt.Println("in GetComputerSystem start")
	currentDistro := distro.GetInstance()
	comSys, err := currentDistro.DistroGetComputerSystem()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("in GetComputerSystem end")
	return comSys
}

func (bg *BaseGatherer) GetComputerPatches() *model.ComputerPatchesType {
	fmt.Println("in GetComputerPatches start")
	currentDistro := distro.GetInstance()
	comPatch, err := currentDistro.DistroGetComputerPatches()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("in GetComputerPatches end")
	return comPatch
}

//  All Info
func (bg *BaseGatherer) GatherInfo() *model.ComputerInfoType {

	m := model.ComputerInfoType{}

	currentDistro := distro.GetInstance()
	if strings.Contains(reflect.TypeOf(currentDistro).String(), "MacBased") {
		distro.MacGetAllInfoInMap()
	}
	fmt.Println("We are in Gather Info ")

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
