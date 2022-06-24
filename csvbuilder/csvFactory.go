package csvbuilder

import (
	"fmt"
	"machine_info_gatherer/gatherer"

	"github.com/jszwec/csvutil"
)

func CreateLinuxCSVFile(linuxInfo gatherer.LinuxGatherer) {
	b, err := csvutil.Marshal(linuxInfo.Baseboard)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}

func CreateMacCSVFile(macInfo gatherer.MacGatherer) {
	b, err := csvutil.Marshal(macInfo)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
