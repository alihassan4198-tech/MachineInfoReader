package uploader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"machine_info_gatherer/model"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	client  = &http.Client{}
	baseURL string
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	baseURL = os.Getenv("BASEURL")
}

func HttpRequest(data interface{}, path string) error {
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

func UploadJsonServer(info *model.ComputerInfoType, apiPath string) {
	if strings.Contains(apiPath, "baseboard") {
		if err := HttpRequest(info.ComputerBaseboard, apiPath); err != nil {
			log.Printf("Error uploading JSON data to %s: %v\n", apiPath, err)
		}
	}
	if strings.Contains(apiPath, "bios") {
		if err := HttpRequest(info.ComputerBios, apiPath); err != nil {
			log.Printf("Error uploading JSON data to %s: %v\n", apiPath, err)
		}
	}
	if strings.Contains(apiPath, "cpu") {
		if err := HttpRequest(info.ComputerCPU.CPUCores, apiPath); err != nil {
			log.Printf("Error uploading JSON data to %s: %v\n", apiPath, err)
		}
	}
	if strings.Contains(apiPath, "endpoint-protection") {
		if err := HttpRequest(info.ComputerEndpointProtection, apiPath); err != nil {
			log.Printf("Error uploading JSON data to %s: %v\n", apiPath, err)
		}
	}
	if strings.Contains(apiPath, "fw-rules") {
		if err := HttpRequest(info.ComputerFirewallRules.FW_rules, apiPath); err != nil {
			log.Printf("Error uploading JSON data to %s: %v\n", apiPath, err)
		}
	}
	if strings.Contains(apiPath, "nics") {
		if err := HttpRequest(info.ComputerNICS, apiPath); err != nil {
			log.Printf("Error uploading JSON data to %s: %v\n", apiPath, err)
		}
	}
	if strings.Contains(apiPath, "os") {
		if err := HttpRequest(info.ComputerOS, apiPath); err != nil {
			log.Printf("Error uploading JSON data to %s: %v\n", apiPath, err)
		}
	}
	if strings.Contains(apiPath, "patches") {
		if err := HttpRequest(info.ComputerPatches.Patches, apiPath); err != nil {
			log.Printf("Error uploading JSON data to %s: %v\n", apiPath, err)
		}
	}
	if strings.Contains(apiPath, "services") {
		if err := HttpRequest(info.ComputerServices, apiPath); err != nil {
			log.Printf("Error uploading JSON data to %s: %v\n", apiPath, err)
		}
	}
	if strings.Contains(apiPath, "software") {
		if err := HttpRequest(info.ComputerSoftwaresInstalled, apiPath); err != nil {
			log.Printf("Error uploading JSON data to %s: %v\n", apiPath, err)
		}
	}
	if strings.Contains(apiPath, "system") {
		if err := HttpRequest(info.ComputerSystem, apiPath); err != nil {
			log.Printf("Error uploading JSON data to %s: %v\n", apiPath, err)
		}
	}
}
