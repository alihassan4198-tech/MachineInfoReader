package main

import (
	"fmt"
	"machine_info_gatherer/debug"
	"machine_info_gatherer/gatherer"
	"machine_info_gatherer/uploader"
)

func main() {
	fmt.Println("Launching Machine Info Gatherer...")
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
