package gatherer

import "machine_info_gatherer/model"

type I_Gatherer interface {
	GetComputerBaseboard() *model.ComputerBaseboard
	GatherInfo() *model.ComputerInfo
	GetComputerBios() *model.ComputerBios
}
