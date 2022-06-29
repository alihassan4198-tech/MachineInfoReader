package distro

import (
	"machine_info_gatherer/model"
)

// Arch Linux

type ArchLinuxBased struct {
	LinuxBase
}

func (alb *ArchLinuxBased) GetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalled, error) {
	comSoftInst := model.ComputerSoftwaresInstalled{}

	return &comSoftInst, nil
}
