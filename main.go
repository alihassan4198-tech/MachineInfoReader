package main

import (
	"fmt"
	"machine_info_gatherer/common"
	"machine_info_gatherer/csvfilecreator"
	"machine_info_gatherer/csvfileuploader"
	"machine_info_gatherer/gatherer"
	"machine_info_gatherer/jsoncreator"
	"os"
)

func main() {
	fmt.Println("Launching Machine Info Gatherer...")

	// common.SetSudoPassword("123456")

	common.PathSetter(os.Args[1])
	i := gatherer.GetInstance()
	info := i.GatherInfo()
	jsoncreator.JsonFilesCreator(info)
	csvfilecreator.CsvFilesCreator(info)
	csvfileuploader.CsvFilesUploader(info)
}
