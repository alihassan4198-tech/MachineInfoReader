package gatherer

import "machine_info_gatherer/model"

type I_Gatherer interface {
	GatherInfo() *model.ComputerInfo
	GetComputerBaseboard() *model.ComputerBaseboard
	GetComputerBios() *model.ComputerBios
	GetComputerCPU() *model.ComputerCPU
	GetComputerEndpointProtectionSoftwares() *model.ComputerEndpointProtection
	GetComputerFirewallRules() *model.ComputerFirewallRules
	GetComputerNIC() *[]model.ComputerNIC
	GetComputerOS() *model.ComputerOS
	GetComputerServices() *model.ComputerServices
	GetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalled, error)
}
