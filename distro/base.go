package distro

import (
	"errors"
	"machine_info_gatherer/errorslist"
	"machine_info_gatherer/model"
)

// Linux Base Distro

type LinuxBase struct {
	model.ComputerSoftwaresInstalled
}

func (lb *LinuxBase) GetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalled, error) {
	comSoftInst := model.ComputerSoftwaresInstalled{}

	return &comSoftInst, errors.New(errorslist.ErrNotImplemented)

}
