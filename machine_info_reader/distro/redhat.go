package distro

import (
	"machine_info_gatherer/model"
)

// Red Hat

type RedHatBased struct {
	LinuxBase
}

func (rhb *RedHatBased) GetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalledType, error) {
	comSoftInst := model.ComputerSoftwaresInstalledType{}

	return &comSoftInst, nil
}
