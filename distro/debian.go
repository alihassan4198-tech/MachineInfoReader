package distro

import (
	"fmt"
	"machine_info_gatherer/common"
	"machine_info_gatherer/model"
	"os"
	"os/exec"
	"strings"

	"github.com/alihassan4198-tech/ghw"
)

// Debian

type DebianBased struct {
	LinuxBase
}

const (
	find0                   = "Name"
	find1                   = "Version"
	find2                   = "Architecture"
	find3                   = "Description"
	baseboardCaption string = "Base Board"
)

func (lb *DebianBased) GetComputerBaseboard() (*model.ComputerBaseboardType, error) {
	cb := model.ComputerBaseboardType{}

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

	cb.Serialnumber = common.RootNeeded(baseboard.SerialNumber)
	cb.Version = common.RootNeeded(baseboard.Version)
	cb.Product = common.RootNeeded(baseboard.Product)
	cb.Manufacturer = common.RootNeeded(baseboard.Vendor)
	cb.Tag = common.RootNeeded(baseboard.AssetTag)

	cb.Creationclassname, err = common.RunFullCommand("uname -m")
	if err != nil {
		cb.Creationclassname = ""
	}
	cb.Creationclassname = strings.ReplaceAll(cb.Creationclassname, "\n", "")

	installDate, err := common.RunFullCommand("ls -ld --time-style=full-iso /var/log/installer")
	if err != nil {
		cb.Installdate = ""
	} else {
		cb.Installdate = func(str string) string {
			res := strings.Split(str, "/var")
			res = strings.Split(res[0], " ")
			result := strings.Join(res[5:8], " ")
			return result
		}(installDate)
	}

	cb.Poweredon = true

	return &cb, nil
}

func (lb *DebianBased) GetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalledType, error) {
	comSoftInst := model.ComputerSoftwaresInstalledType{}

	installSoft, err := exec.Command("dpkg", "-l").Output()
	if err != nil {
		fmt.Println(err)
		return &comSoftInst, err
	}
	splittedInstallSoft := strings.Split(string(installSoft), "\n")

	for i, soft := range splittedInstallSoft {
		if strings.Contains(soft, find0) &&
			strings.Contains(soft, find1) &&
			strings.Contains(soft, find2) &&
			strings.Contains(soft, find3) {
			// Remove first 2 useless record strings
			splittedInstallSoft = splittedInstallSoft[i+2:]
			break
		}
	}

	// Remove last \n record string
	splittedInstallSoft = splittedInstallSoft[:len(splittedInstallSoft)-1]

	// Convert to Strcut
	for _, soft := range splittedInstallSoft {
		soft = strings.Join(strings.Fields(soft), " ")
		splittedSoft := strings.Split(soft, " ")
		softInstall := model.SoftwareInstalledType{}
		softInstall.Display_name = splittedSoft[1]
		softInstall.Version = splittedSoft[2]

		comSoftInst.SoftwaresInstalled = append(comSoftInst.SoftwaresInstalled, softInstall)
	}

	comSoftInst.Total_software = len(comSoftInst.SoftwaresInstalled)

	return &comSoftInst, nil
}
