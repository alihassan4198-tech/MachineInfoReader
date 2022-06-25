package gatherer

import "machine_info_gatherer/model"

type I_Gatherer interface {
	GetComputerName() (string, error)
	GatherInfo() model.ComputerInfo
}
