package main

import (
	"fmt"
	"machine_info_gatherer/csvfilecreator"
	"machine_info_gatherer/gatherer"
	"machine_info_gatherer/jsoncreator"
)

func main() {
	fmt.Println("Launching Machine Info Gatherer...")

	// common.SetSudoPassword("786")

	i := gatherer.GetInstance()
	info := i.GatherInfo()
	jsoncreator.JsonFilesCreator(info)
	csvfilecreator.CsvFilesCreator(info)

}
