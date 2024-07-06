package uploader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"machine_info_gatherer/model"
	"net/http"
	"strings"
	"sync"
)

var (
	client  = &http.Client{}
	baseURL = "http://localhost:3020/api/v1/"
)

func httpRequest(data interface{}, path string) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	url := baseURL + path
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error making POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	log.Println("Data uploaded successfully to", url)
	return nil
}

func UploadJsonFile(info *model.ComputerInfoType, apiPath string) {

	var wg sync.WaitGroup
	checkAndUpload := func(data interface{}, path string) {
		defer wg.Done()
		if err := httpRequest(data, path); err != nil {
			log.Printf("Error uploading JSON data to %s: %v\n", path, err)
		}
	}
	if strings.Contains(apiPath, "baseboard") {
		wg.Add(1)
		go checkAndUpload(info.ComputerBaseboard, apiPath)
	}
	if strings.Contains(apiPath, "bios") {
		wg.Add(1)
		go checkAndUpload(info.ComputerBios, apiPath)
	}
	if strings.Contains(apiPath, "cpu") {
		wg.Add(1)
		go checkAndUpload(info.ComputerCPU.CPUCores, apiPath)
	}
	if strings.Contains(apiPath, "endpoint-protection") {
		wg.Add(1)
		go checkAndUpload(info.ComputerEndpointProtection, apiPath)
	}
	if strings.Contains(apiPath, "fw-rules") {
		wg.Add(1)
		go checkAndUpload(info.ComputerFirewallRules.FW_rules, apiPath)
	}
	if strings.Contains(apiPath, "nics") {
		wg.Add(1)
		go checkAndUpload(info.ComputerNICS, apiPath)
	}
	if strings.Contains(apiPath, "os") {
		wg.Add(1)
		go checkAndUpload(info.ComputerOS, apiPath)
	}
	if strings.Contains(apiPath, "patches") {
		wg.Add(1)
		go checkAndUpload(info.ComputerPatches.Patches, apiPath)
	}
	if strings.Contains(apiPath, "services") {
		wg.Add(1)
		go checkAndUpload(info.ComputerServices, apiPath)
	}
	if strings.Contains(apiPath, "software") {
		wg.Add(1)
		go checkAndUpload(info.ComputerSoftwaresInstalled, apiPath)
	}
	if strings.Contains(apiPath, "system") {
		wg.Add(1)
		go checkAndUpload(info.ComputerSystem, apiPath)
	}
	wg.Wait()
}
