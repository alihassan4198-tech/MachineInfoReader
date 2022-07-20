package csvfilecreator

import (
	"machine_info_gatherer/builder"
	"machine_info_gatherer/model"
)

func CsvFilesCreator(info *model.ComputerInfoType) {
	// builder.CreateCSVFile("ComputerBaseboard")
	// builder.CreateCSVFile("ComputerBios")
	// builder.CreateCSVFile("ComputerCPU")
	// builder.CreateCSVFile("ComputerEndpointProtection")
	// builder.CreateCSVFile("ComputerFirewallRules")
	// builder.CreateCSVFile("ComputerNICS")
	// builder.CreateCSVFile("ComputerOS")
	// builder.CreateCSVFile("ComputerPatches")
	// builder.CreateCSVFile("ComputerServices")
	builder.CreateCSVFile("ComputerSoftwaresInstalled", info.ComputerSoftwaresInstalled.Total_software)
	// builder.CreateCSVFile("ComputerSystem")
}
