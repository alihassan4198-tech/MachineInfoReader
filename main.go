package main

import (
	"fmt"
	"machine_info_gatherer/csvbuilder"
	"machine_info_gatherer/gatherer"
)

func main() {
	fmt.Println("Launching Machine Info Gatherer...")

	// common.SetSudoPassword("786")

	i := gatherer.GetInstance()
	info := i.GatherInfo()
	csvbuilder.CreateCSVFile(info)
}
