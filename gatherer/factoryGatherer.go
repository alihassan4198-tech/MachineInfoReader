package gatherer

import (
	"runtime"
)

func GetInstance() I_Gatherer {
	os := runtime.GOOS
	if os == "darwin" { // Mac
		macMachineInfo := MacGatherer{}
		return &macMachineInfo
	} else if os == "linux" {
		linuxMachineInfo := LinuxGatherer{}
		return &linuxMachineInfo
	}

	return nil
}
