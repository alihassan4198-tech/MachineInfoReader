package distro

import "machine_info_gatherer/model"

type IDistro interface {
	GetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalled, error)
}
