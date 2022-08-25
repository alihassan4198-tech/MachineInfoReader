package csvfileuploader

import (
	"bytes"
	"fmt"
	"io"
	"machine_info_gatherer/model"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadCSVFile(fieldName, fileName string, url string) {
	b, w := CSVFileToBytes(fieldName, fileName)
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
}

func CSVFileToBytes(fieldName, fileName string) (bytes.Buffer, *multipart.Writer) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)
	var fw io.Writer
	file := mustOpen(fileName)

	if fw, err = w.CreateFormFile(fieldName, file.Name()); err != nil {
		fmt.Println(err)
	}
	if _, err = io.Copy(fw, file); err != nil {
		fmt.Println(err)
	}
	w.Close()
	return b, w
}

func mustOpen(f string) *os.File {
	pwd, _ := os.Getwd()
	path := pwd + "/" + f
	r, err := os.Open(path)
	if err != nil {
		fmt.Println("PWD: ", pwd)
		panic(err)
	}
	return r
}

func CsvFilesUploader(info *model.ComputerInfoType) {
	fmt.Println("Uploading CSV files...")
	UploadCSVFile("ComputerBaseboard", info.ComputerBaseboard.Computer_name+"ComputerBaseboard.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerBios", info.ComputerBios.Name+"ComputerBios.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerCPU", info.ComputerBios.Name+"ComputerCPU.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerEndpointProtection", info.ComputerBios.Name+"ComputerEndpointProtection.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerFirewallRules", info.ComputerBios.Name+"ComputerFirewallRules.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerNICS", info.ComputerBios.Name+"ComputerNICS.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerOS", info.ComputerBios.Name+"ComputerOS.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerPatches", info.ComputerBios.Name+"ComputerPatches.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerServices", info.ComputerBios.Name+"ComputerServices.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerSoftwaresInstalled", info.ComputerBios.Name+"ComputerSoftwaresInstalled.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerSystem", info.ComputerBios.Name+"ComputerSystem.csv", "http://localhost:3010/uploadfile")
}
