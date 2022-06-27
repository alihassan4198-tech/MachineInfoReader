package main

import (
	"fmt"
	"machine_info_gatherer/gatherer"
)

func main() {
	fmt.Println("Launching Machine Info Gatherer...")
	i := gatherer.GetInstance()
	i.GatherInfo()
}
