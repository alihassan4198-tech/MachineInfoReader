package gatherer

import (
	"fmt"
	"machine_info_gatherer/debug"
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
	currentDistro := distro.GetInstance()

	comBaseBoard, err := currentDistro.DistroGetComputerBaseboard()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	ownerInfo, err := currentDistro.DistroGetComputerOwner()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	comBaseBoard.OwnerInfo = *ownerInfo

	return comBaseBoard
}

func (bg *BaseGatherer) GetComputerBios() *model.ComputerBiosType {

	currentDistro := distro.GetInstance()
	cbios, err := currentDistro.DistroGetComputerBios()
	if err != nil {
		fmt.Println(err)
	}
	ownerInfo, err := currentDistro.DistroGetComputerOwner()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cbios.OwnerInfo = *ownerInfo

	return cbios
}

func (bg *BaseGatherer) GetComputerCPU() *model.ComputerCPUType {

	currentDistro := distro.GetInstance()
	compCpu, err := currentDistro.DistroGetComputerCPU()
	if err != nil {
		fmt.Println(err)
	}
	ownerInfo, err := currentDistro.DistroGetComputerOwner()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	compCpu.OwnerInfo = *ownerInfo

	return compCpu
}

func (bg *BaseGatherer) GetComputerEndpointProtectionSoftwares() *model.ComputerEndpointProtectionType {

	currentDistro := distro.GetInstance()
	epsoft, err := currentDistro.DistroGetComputerEndpointProtectionSoftwares()
	if err != nil {
		fmt.Println(err)
	}
	ownerInfo, err := currentDistro.DistroGetComputerOwner()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	epsoft.OwnerInfo = *ownerInfo
	return epsoft
}

func (bg *BaseGatherer) GetComputerFirewallRules() *model.ComputerFirewallRulesType {

	currentDistro := distro.GetInstance()
	cfwRules, err := currentDistro.DistroGetComputerFirewallRules()
	if err != nil {
		fmt.Println(err)
	}
	ownerInfo, err := currentDistro.DistroGetComputerOwner()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cfwRules.OwnerInfo = *ownerInfo
	return cfwRules
}

func (bg *BaseGatherer) GetComputerNIC() *[]model.ComputerNICType {

	currentDistro := distro.GetInstance()

	comNicList, err := currentDistro.DistroGetComputerNIC()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	ownerInfo, err := currentDistro.DistroGetComputerOwner()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for i := range *comNicList {
		(*comNicList)[i].OwnerInfo = *ownerInfo
	}

	return comNicList
}

func (bg *BaseGatherer) GetComputerOS() *model.ComputerOSType {

	currentDistro := distro.GetInstance()
	comOS, err := currentDistro.DistroGetComputerOS()
	if err != nil {
		fmt.Println(err)
	}
	ownerInfo, err := currentDistro.DistroGetComputerOwner()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	comOS.OwnerInfo = *ownerInfo
	return comOS
}

func (bg *BaseGatherer) GetComputerServices() *model.ComputerServicesType {

	currentDistro := distro.GetInstance()
	comServ, err := currentDistro.DistroGetComputerServices()
	if err != nil {
		fmt.Println(err)
	}
	ownerInfo, err := currentDistro.DistroGetComputerOwner()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	comServ.OwnerInfo = *ownerInfo
	return comServ
}

func (bg *BaseGatherer) GetComputerSoftwaresInstalled() *model.ComputerSoftwaresInstalledType {

	currentDistro := distro.GetInstance()
	comInsSoft, err := currentDistro.DistroGetComputerSoftwaresInstalled()
	if err != nil {
		fmt.Println(err)
	}
	ownerInfo, err := currentDistro.DistroGetComputerOwner()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	comInsSoft.OwnerInfo = *ownerInfo
	return comInsSoft
}
func (bg *BaseGatherer) GetComputerOwner() *model.ComputerOwnerType {
	currentDistro := distro.GetInstance()
	ownerDetails, err := currentDistro.DistroGetComputerOwner()
	if err != nil {
		fmt.Println(err)
	}
	return ownerDetails
}
func (bg *BaseGatherer) GetComputerSystem() *model.ComputerSystemType {

	currentDistro := distro.GetInstance()
	comSys, err := currentDistro.DistroGetComputerSystem()
	if err != nil {
		fmt.Println(err)
	}
	ownerInfo, err := currentDistro.DistroGetComputerOwner()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	comSys.OwnerInfo = *ownerInfo
	return comSys
}

func (bg *BaseGatherer) GetComputerPatches() *model.ComputerPatchesType {

	currentDistro := distro.GetInstance()
	comPatch, err := currentDistro.DistroGetComputerPatches()
	if err != nil {
		fmt.Println(err)
	}
	ownerInfo, err := currentDistro.DistroGetComputerOwner()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	comPatch.OwnerInfo = *ownerInfo
	return comPatch
}

// All Info
func (bg *BaseGatherer) GatherInfo() *model.ComputerInfoType {

	m := model.ComputerInfoType{}
	currentDistro := distro.GetInstance()
	if debug.Debug() {
		fmt.Println("currentDistro := distro.GetInstance()")
	}
	if strings.Contains(reflect.TypeOf(currentDistro).String(), "MacBased") {
		if debug.Debug() {
			fmt.Println("strings.Contains(reflect.TypeOf(currentDistro).String(), MacBased)")
		}
		distro.MacGetAllInfoInMap()
	}
	m.ComputerBaseboard = *(bg.GetComputerBaseboard())
	if debug.Debug() {
		fmt.Println("m.ComputerBaseboard = *(bg.GetComputerBaseboard())")
	}
	m.ComputerBios = *(bg.GetComputerBios())
	if debug.Debug() {
		fmt.Println("m.ComputerBios = *(bg.GetComputerBios())")
	}
	m.ComputerCPU = *(bg.GetComputerCPU())
	if debug.Debug() {
		fmt.Println("m.ComputerCPU = *(bg.GetComputerCPU())")
	}
	m.ComputerEndpointProtection = *(bg.GetComputerEndpointProtectionSoftwares())
	if debug.Debug() {
		fmt.Println("m.ComputerEndpointProtection = *(bg.GetComputerEndpointProtectionSoftwares())")
	}
	m.ComputerFirewallRules = *(bg.GetComputerFirewallRules())
	if debug.Debug() {
		fmt.Println("m.ComputerFirewallRules = *(bg.GetComputerFirewallRules())")
	}
	m.ComputerNICS = *(bg.GetComputerNIC())
	if debug.Debug() {
		fmt.Println("m.ComputerNICS = *(bg.GetComputerNIC())")
	}
	m.ComputerOS = *(bg.GetComputerOS())
	if debug.Debug() {
		fmt.Println("m.ComputerOS = *(bg.GetComputerOS())")
	}
	m.ComputerServices = *(bg.GetComputerServices())
	if debug.Debug() {
		fmt.Println("m.ComputerServices = *(bg.GetComputerServices())")
	}
	m.ComputerSoftwaresInstalled = *(bg.GetComputerSoftwaresInstalled())
	if debug.Debug() {
		fmt.Println("m.ComputerSoftwaresInstalled = *(bg.GetComputerSoftwaresInstalled())")
	}
	m.ComputerSystem = *(bg.GetComputerSystem())
	if debug.Debug() {
		fmt.Println("m.ComputerSystem = *(bg.GetComputerSystem())")
	}
	m.ComputerPatches = *(bg.GetComputerPatches())
	if debug.Debug() {
		fmt.Println("m.ComputerPatches = *(bg.GetComputerPatches())")
	}
	return &m
}
