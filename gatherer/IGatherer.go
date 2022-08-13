package gatherer

import "machine_info_gatherer/model"

type I_Gatherer interface {
	GatherInfo() *model.ComputerInfoType
	GetComputerBaseboard() *model.ComputerBaseboardType
	GetComputerBios() *model.ComputerBiosType
	GetComputerCPU() *model.ComputerCPUType
	GetComputerEndpointProtectionSoftwares() *model.ComputerEndpointProtectionType
	GetComputerFirewallRules() *model.ComputerFirewallRulesType
	GetComputerNIC() *[]model.ComputerNICType
	GetComputerOS() *model.ComputerOSType
	GetComputerServices() *model.ComputerServicesType
	GetComputerSoftwaresInstalled() *model.ComputerSoftwaresInstalledType
	GetComputerSystem() *model.ComputerSystemType
	GetComputerPatches() *model.ComputerPatchesType
}
