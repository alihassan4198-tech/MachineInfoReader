package gatherer

import (
	"fmt"
	"machine_info_gatherer/model"
	"os"
)

type BaseGatherer struct {
}


//--------------------------
// Baseboard Implementation
//--------------------------

const (
	baseboardCaption string = "Base Board"
)

func (bg *BaseGatherer) GetComputerBaseboard() *model.ComputerBaseboard {
	cb := model.ComputerBaseboard{}

	var err error
	computerName, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	cb.Computer_name = computerName

	return &cb
}




func (bg *BaseGatherer) GetSystem() string {

}



//  All Info
func (bg *BaseGatherer) GatherInfo() model.ComputerInfo {

	m := model.ComputerInfo{}

	m.ComputerBaseboard = *(bg.GetComputerBaseboard())
	

	return m
}
