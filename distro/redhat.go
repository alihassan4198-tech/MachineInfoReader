package distro

import (
	"machine_info_gatherer/model"
)

// Red Hat

type RedHatBased struct {
	LinuxBase
}

func (rhb *RedHatBased) GetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalled, error) {
	comSoftInst := model.ComputerSoftwaresInstalled{}

	return &comSoftInst, nil
}
