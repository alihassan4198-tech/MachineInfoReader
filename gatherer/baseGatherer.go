package gatherer

import (
	"fmt"
	"machine_info_gatherer/model"
	"os"

	"github.com/jaypipes/ghw"
)

type BaseGatherer struct {
}

//--------------------------
// Baseboard Implementation
//--------------------------

const (
	baseboardCaption string = "Base Board"
)

func doesItNeedRootAccess(arg string) string {
	if arg == "None" || arg == "unknown" {
		return arg + " (need root access)"
	} else {
		return arg
	}
}

func (bg *BaseGatherer) GetComputerBaseboard() *model.ComputerBaseboard {
	cb := model.ComputerBaseboard{}

	var err error
	computerName, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	cb.Computer_name = computerName
	cb.Caption = baseboardCaption

	baseboard, err := ghw.Baseboard()
	if err != nil {
		fmt.Printf("baseboard error : %s", err)
	}

	cb.Serialnumber = doesItNeedRootAccess(baseboard.SerialNumber)
	cb.Version = doesItNeedRootAccess(baseboard.Version)
	cb.Product = doesItNeedRootAccess(baseboard.Product)
	cb.Manufacturer = doesItNeedRootAccess(baseboard.Vendor)
	cb.Tag = doesItNeedRootAccess(baseboard.AssetTag)

	return &cb
}

func (bg *BaseGatherer) GetRawInformation() {

}

//  All Info
func (bg *BaseGatherer) GatherInfo() model.ComputerInfo {

	m := model.ComputerInfo{}

	m.ComputerBaseboard = *(bg.GetComputerBaseboard())

	bg.GetRawInformation()

	return m
}
