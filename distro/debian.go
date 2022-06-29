package distro

import (
	"fmt"
	"machine_info_gatherer/model"
	"os/exec"
	"strings"
)

// Debian

type DebianBased struct {
	LinuxBase
}

const (
	find0 = "Name"
	find1 = "Version"
	find2 = "Architecture"
	find3 = "Description"
)

func (lb *DebianBased) GetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalled, error) {
	comSoftInst := model.ComputerSoftwaresInstalled{}

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
		softInstall := model.SoftwareInstalled{}
		softInstall.Display_name = splittedSoft[1]
		softInstall.Version = splittedSoft[2]

		comSoftInst.SoftwaresInstalled = append(comSoftInst.SoftwaresInstalled, softInstall)
	}

	comSoftInst.Total_software = len(comSoftInst.SoftwaresInstalled)

	return &comSoftInst, nil
}
