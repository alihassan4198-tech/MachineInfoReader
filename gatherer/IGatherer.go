package gatherer

import "machine_info_gatherer/model"

type I_Gatherer interface {
	GatherInfo() *model.ComputerInfo
	GetComputerBaseboard() *model.ComputerBaseboard
	GetComputerBios() *model.ComputerBios
	GetComputerCPU() *model.ComputerCPU
}
