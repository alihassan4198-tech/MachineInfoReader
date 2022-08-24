package csvfileuploader

import (
	"fmt"
	"machine_info_gatherer/builder"
	"machine_info_gatherer/model"
)

func CsvFilesUploader(info *model.ComputerInfoType) {
	fmt.Println("Uploading CSV files...")
	builder.UploadCSVFile("ComputerBaseboard", info)
	builder.UploadCSVFile("ComputerBios", info)
	builder.UploadCSVFile("ComputerCPU", info)
	builder.UploadCSVFile("ComputerEndpointProtection", info)
	builder.UploadCSVFile("ComputerFirewallRules", info)
	builder.UploadCSVFile("ComputerNICS", info)
	builder.UploadCSVFile("ComputerOS", info)
	builder.UploadCSVFile("ComputerPatches", info)
	builder.UploadCSVFile("ComputerServices", info)
	builder.UploadCSVFile("ComputerSoftwaresInstalled", info)
	builder.UploadCSVFile("ComputerSystem", info)
}
