package gatherer

import (
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
	return model.ComputerInfo{}
}
