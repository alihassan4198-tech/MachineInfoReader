package main

import (
	"fmt"
	"machine_info_gatherer/common"
	"machine_info_gatherer/debug"
	"machine_info_gatherer/gatherer"
	"machine_info_gatherer/uploader"
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
	uploader.JsonDataUploader(info)
	if debug.Debug() {
		fmt.Println("uploader.JsonDataUploader(info)")
	}
}
