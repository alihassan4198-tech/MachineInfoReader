package main

import (
	"fmt"
	"machine_info_gatherer/common"
	"machine_info_gatherer/csvfilecreator"
	"machine_info_gatherer/csvfileuploader"
	"machine_info_gatherer/gatherer"
	"machine_info_gatherer/jsoncreator"
	"os"
	"time"
)

func main() {
	fmt.Println("Launching Machine Info Gatherer...")

	// common.SetSudoPassword("123456")
	fmt.Println("new info : " + time.Now().Local().Format(time.UnixDate))

	common.PathSetter(os.Args[1])
	fmt.Println(common.PathGetter())
	i := gatherer.GetInstance()
	fmt.Println("Next Line of code : info := i.GatherInfo() ")
	info := i.GatherInfo()
	fmt.Println("Now Starts to create json files")
	jsoncreator.JsonFilesCreator(info)
	fmt.Println("Now Starts to create CSV files")
	csvfilecreator.CsvFilesCreator(info)
	fmt.Println("Now Starts to create CSVUploader files")
	csvfileuploader.CsvFilesUploader(info)
}
