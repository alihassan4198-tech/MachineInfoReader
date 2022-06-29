package distro

import (
	"machine_info_gatherer/model"
)

// Debian

type DebianBased struct {
	LinuxBase
}

func (lb *DebianBased) GetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalled, error) {
	comSoftInst := model.ComputerSoftwaresInstalled{}

	return &comSoftInst, nil
}
