package distro

import (
	"errors"
	"machine_info_gatherer/errorslist"
	"machine_info_gatherer/model"
)

// Linux Base Distro

type LinuxBase struct {
	model.ComputerSoftwaresInstalledType
}

func (lb *LinuxBase) GetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalledType, error) {
	comSoftInst := model.ComputerSoftwaresInstalledType{}

	return &comSoftInst, errors.New(errorslist.ErrNotImplemented)

}
