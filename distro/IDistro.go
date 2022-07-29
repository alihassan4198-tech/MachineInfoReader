package distro

import "machine_info_gatherer/model"

type IDistro interface {
	DistroGatherInfo() (*model.ComputerInfoType, error)
	DistroGetComputerBaseboard() (*model.ComputerBaseboardType, error)
	DistroGetComputerBios() (*model.ComputerBiosType,error)
	DistroGetComputerCPU() (*model.ComputerCPUType, error)
	DistroGetComputerEndpointProtectionSoftwares() (*model.ComputerEndpointProtectionType, error)
	DistroGetComputerFirewallRules() (*model.ComputerFirewallRulesType, error)
	DistroGetComputerNIC() (*[]model.ComputerNICType, error)
	DistroGetComputerOS() (*model.ComputerOSType, error)
	DistroGetComputerServices() (*model.ComputerServicesType, error)
	DistroGetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalledType, error)
	DistroGetComputerSystem() (*model.ComputerSystemType, error)
	DistroGetComputerPatches() (*model.ComputerPatchesType, error)
}
