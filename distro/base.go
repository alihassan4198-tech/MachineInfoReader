package distro

import (
	"machine_info_gatherer/model"
)

// Linux Base Distro

const (
	baseboardCaption string = "Base Board"
	cpuCaption       string = "CPU"
	osCaption        string = "Computer OS"
	biosCaption      string = "Computer Bios"
)

type LinuxBase struct {
	model.ComputerBaseboardType

	model.ComputerSoftwaresInstalledType
}
