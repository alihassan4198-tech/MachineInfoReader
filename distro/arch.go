package distro

import (
	"machine_info_gatherer/model"
)

// Arch Linux

type ArchLinuxBased struct {
	LinuxBase
}

func (alb *ArchLinuxBased) GetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalledType, error) {
	comSoftInst := model.ComputerSoftwaresInstalledType{}

	return &comSoftInst, nil
}
