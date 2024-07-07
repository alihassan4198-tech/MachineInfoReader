package uploader

import (
	"fmt"
	"machine_info_gatherer/model"
)

func JsonDataUploader(info *model.ComputerInfoType) {
	fmt.Println("Uploading JSON files...")
	Init()
	UploadJsonServer(info, "baseboard")
	UploadJsonServer(info, "bios")
	UploadJsonServer(info, "cpu")
	UploadJsonServer(info, "endpoint-protection")
	UploadJsonServer(info, "fw-rules")
	UploadJsonServer(info, "nics")
	UploadJsonServer(info, "os")
	UploadJsonServer(info, "patches")
	UploadJsonServer(info, "services")
	UploadJsonServer(info, "software")
	UploadJsonServer(info, "system")
}
