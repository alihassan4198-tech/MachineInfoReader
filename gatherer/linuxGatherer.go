package gatherer

import (
	"machine_info_gatherer/info"
)

type LinuxGatherer struct {
	Baseboard info.ComputerBaseboard `csv:"baseboard"`
}

func (lg *LinuxGatherer) FetchComputerName() {
	lg.Baseboard.GetComputerName()
}

func (lg *LinuxGatherer) FetchAllBaseBoard() {
	lg.FetchComputerName()
}
