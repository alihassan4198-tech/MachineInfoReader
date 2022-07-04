package distro

import (
	"errors"
	"machine_info_gatherer/errorslist"
	"machine_info_gatherer/model"
)

// Linux Base Distro

type LinuxBase struct {
	model.ComputerSoftwaresInstalledType
	model.ComputerBaseboardType
}

func (lb *LinuxBase) GetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalledType, error) {
	comSoftInst := model.ComputerSoftwaresInstalledType{}

	return &comSoftInst, errors.New(errorslist.ErrNotImplemented)

}

func (lb *LinuxBase) GetComputerBaseboard() (*model.ComputerBaseboardType, error) {
	comBaseBoard := model.ComputerBaseboardType{}

	return &comBaseBoard, errors.New(errorslist.ErrNotImplemented)

}
