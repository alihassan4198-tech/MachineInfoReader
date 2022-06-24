package main

import (
	"fmt"
	"machine_info_gatherer/gatherer"
	"runtime"
)

func main() {
	fmt.Println("Launching Machine Info Gatherer...")
	os := runtime.GOOS
	if os == "darwin" { // Mac
		macMachineInfo := gatherer.MacGatherer{}
		macMachineInfo.FetchAllBaseBoard()
		fmt.Printf("%#v\n", macMachineInfo)
	} else if os == "linux" {
		linuxMachineInfo := gatherer.LinuxGatherer{}
		linuxMachineInfo.FetchAllBaseBoard()
		fmt.Printf("%s\n", linuxMachineInfo.Baseboard.Computer_name)
	}
}
