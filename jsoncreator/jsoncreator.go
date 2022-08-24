package jsoncreator

import (
	"fmt"
	"machine_info_gatherer/builder"
	"machine_info_gatherer/model"
)

func JsonFilesCreator(info *model.ComputerInfoType) {
	fmt.Println("Creating JSON files...")
	builder.CreateJsonFile(info.ComputerBaseboard, "ComputerBaseboard")
	builder.CreateJsonFile(info.ComputerBios, "ComputerBios")
	builder.CreateJsonFile(info.ComputerCPU.CPUCores, "ComputerCPU")
	builder.CreateJsonFile(info.ComputerEndpointProtection.Softwares, "ComputerEndpointProtection")
	builder.CreateJsonFile(info.ComputerFirewallRules.FW_rules, "ComputerFirewallRules")
	builder.CreateJsonFile(info.ComputerNICS, "ComputerNICS")
	builder.CreateJsonFile(info.ComputerOS, "ComputerOS")
	builder.CreateJsonFile(info.ComputerPatches, "ComputerPatches")
	builder.CreateJsonFile(info.ComputerServices.Services, "ComputerServices")
	builder.CreateJsonFile(info.ComputerSoftwaresInstalled.SoftwaresInstalled, "ComputerSoftwaresInstalled")
	builder.CreateJsonFile(info.ComputerSystem, "ComputerSystem")
}
