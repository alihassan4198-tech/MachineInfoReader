package main

import (
	"fmt"
	"machine_info_gatherer/common"
	"machine_info_gatherer/gatherer"
)

func main() {
	fmt.Println("Launching Machine Info Gatherer...")
	common.SetSudoPassword("123456")
	i := gatherer.GetInstance()
	i.GatherInfo()
}
