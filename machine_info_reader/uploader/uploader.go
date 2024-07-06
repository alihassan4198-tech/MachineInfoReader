package uploader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"machine_info_gatherer/model"
	"strings"
)

func httpRequest(data interface{}, path string) error {

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	url := fmt.Sprintf("http://localhost:3020/%s", path)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error making POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	fmt.Println("Data uploaded successfully")
	return nil
}

func UploadJsonFile(info *model.ComputerInfoType, apiPath string) {

	prefix := "api/v1/"

	if strings.Contains(apiPath, "baseboard") {
		err := httpRequest(info.ComputerBaseboard, prefix+apiPath)
		if err != nil {
			fmt.Printf("Error uploading JSON data: %v\n", err)
		}
	}
}
