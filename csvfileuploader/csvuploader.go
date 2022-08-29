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

func UploadCSVFile(fieldName, fileName string, url string) error {
	b, w, err := CSVFileToBytes(fieldName, fileName)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	return nil
}

func CSVFileToBytes(fieldName, fileName string) (bytes.Buffer, *multipart.Writer, error) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)
	var fw io.Writer
	file, err := mustOpen(fileName)
	if err != nil {
		return b,nil,err
	}


	if fw, err = w.CreateFormFile(fieldName, file.Name()); err != nil {
		fmt.Println(err)
	}
	if _, err = io.Copy(fw, file); err != nil {
		fmt.Println(err)
	}
	w.Close()
	return b, w, err
}

func mustOpen(f string) (*os.File, error) {
	pwd, _ := os.Getwd()
	path := pwd + "/" + f
	r, err := os.Open(path)
	if err != nil {
		fmt.Println("PWD: ", pwd)
		return nil, err
	}
	return r, nil
}

func CsvFilesUploader(info *model.ComputerInfoType) error {
	fmt.Println("Uploading CSV files...")
	UploadCSVFile("ComputerBaseboard", info.ComputerBaseboard.Computer_name+"_"+"ComputerBaseboard.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerBios", info.ComputerBaseboard.Computer_name+"_"+"ComputerBios.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerCPU", info.ComputerBaseboard.Computer_name+"_"+"ComputerCPU.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerEndpointProtection", info.ComputerBaseboard.Computer_name+"_"+"ComputerEndpointProtection.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerFirewallRules", info.ComputerBaseboard.Computer_name+"_"+"ComputerFirewallRules.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerNICS", info.ComputerBaseboard.Computer_name+"_"+"ComputerNICS.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerOS", info.ComputerBaseboard.Computer_name+"_"+"ComputerOS.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerPatches", info.ComputerBaseboard.Computer_name+"_"+"ComputerPatches.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerServices", info.ComputerBaseboard.Computer_name+"_"+"ComputerServices.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerSoftwaresInstalled", info.ComputerBaseboard.Computer_name+"_"+"ComputerSoftwaresInstalled.csv", "http://localhost:3010/uploadfile")
	UploadCSVFile("ComputerSystem", info.ComputerBaseboard.Computer_name+"_"+"ComputerSystem.csv", "http://localhost:3010/uploadfile")
	return nil
}	
