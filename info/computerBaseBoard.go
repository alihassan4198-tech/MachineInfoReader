package info

import "os"

type BaseboardInterface interface {
	GetComputerName() (string, error)
}

type ComputerBaseboard struct {
	Computer_name     string
	Caption           string
	Configoptions     []string
	Creationclassname string
	Description       string
	Installdate       int64
	Manufacturer      string
	Model             string
	Name              string
	Partnumber        string
	Poweredon         bool
	Product           string
	Serialnumber      string
	Sku               string
	Status            string
	Tag               string
	Version           string
}

func (cb *ComputerBaseboard) GetComputerName() (string, error) {

	cname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	cb.Computer_name = cname
	return cname, nil
}
