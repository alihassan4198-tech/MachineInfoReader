package gatherer

import (
	"fmt"
	"machine_info_gatherer/model"
	"os"
)

type BaseGatherer struct {
}

func (bg *BaseGatherer) GetComputerName() (string, error) {

	cname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return cname, nil
}

func (bg *BaseGatherer) GatherInfo() model.ComputerInfo {

	m := model.ComputerInfo{}

	cn, err := bg.GetComputerName()
	if err != nil {
		fmt.Printf("%v", err)
	}
	m.Computer_name = cn

	return m
}
