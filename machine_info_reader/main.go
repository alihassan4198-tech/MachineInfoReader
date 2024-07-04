package main

import (
	"fmt"
	"machine_info_gatherer/common"
	"machine_info_gatherer/debug"
	"machine_info_gatherer/gatherer"
	"machine_info_gatherer/jsoncreator"
	"os"
)

func main() {
	fmt.Println("Launching Machine Info Gatherer...")

	// common.SetSudoPassword("123456")
	common.PathSetter(os.Args[1])
	if debug.Debug() {
		fmt.Println("common.PathSetter(os.Args[1])")
	}
	i := gatherer.GetInstance()
	if debug.Debug() {
		fmt.Println("i := gatherer.GetInstance()")
	}
	info := i.GatherInfo()
	if debug.Debug() {
		fmt.Println("info := i.GatherInfo()")
	}
	jsoncreator.JsonFilesCreator(info)
	if debug.Debug() {
		fmt.Println("jsoncreator.JsonFilesCreator(info)")
	}
	// csvfilecreator.CsvFilesCreator(info)
	// if debug.Debug() {
	// 	fmt.Println("csvfilecreator.CsvFilesCreator(info)")
	// }
	// csvfileuploader.CsvFilesUploader(info)
	// if debug.Debug() {
	// 	fmt.Println("csvfileuploader.CsvFilesUploader(info)")
	// }
}
