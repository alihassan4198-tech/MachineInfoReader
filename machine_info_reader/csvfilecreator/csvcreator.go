package csvfilecreator

import (
	"fmt"
	"machine_info_gatherer/builder"
	"machine_info_gatherer/model"
)

func CsvFilesCreator(info *model.ComputerInfoType) {
	fmt.Println("Creating CSV files...")

	builder.CreateCSVFile("ComputerBaseboard", info)
	builder.CreateCSVFile("ComputerBios", info)
	builder.CreateCSVFile("ComputerCPU", info)
	builder.CreateCSVFile("ComputerEndpointProtection", info)
	builder.CreateCSVFile("ComputerFirewallRules", info)
	builder.CreateCSVFile("ComputerNICS", info)
	builder.CreateCSVFile("ComputerOS", info)
	builder.CreateCSVFile("ComputerPatches", info)
	builder.CreateCSVFile("ComputerServices", info)
	builder.CreateCSVFile("ComputerSoftwaresInstalled", info)
	builder.CreateCSVFile("ComputerSystem", info)
}
