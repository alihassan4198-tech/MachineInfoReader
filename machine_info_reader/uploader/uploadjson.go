package uploader

import (
	"fmt"
	"machine_info_gatherer/model"
)

func JsonDataUploader(info *model.ComputerInfoType) {
	fmt.Println("Uploading JSON files...")
	UploadJsonFile(info, "baseboard")
	UploadJsonFile(info, "bios")
	UploadJsonFile(info, "cpu")
	UploadJsonFile(info, "endpoint-protection")
	UploadJsonFile(info, "fw-rules")
	UploadJsonFile(info, "nics")
	UploadJsonFile(info, "os")
	UploadJsonFile(info, "patches")
	UploadJsonFile(info, "services")
	UploadJsonFile(info, "software")
	UploadJsonFile(info, "system")
}
