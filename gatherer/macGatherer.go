package gatherer

import "machine_info_gatherer/info"

type MacGatherer struct {
	Baseboard info.ComputerBaseboard
}

func (mg *MacGatherer) FetchComputerName() {
	mg.Baseboard.GetComputerName()
}

func (mg *MacGatherer) FetchAllBaseBoard() {
	mg.FetchComputerName()
}
