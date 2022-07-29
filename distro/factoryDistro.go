package distro

import (
	"fmt"
	"machine_info_gatherer/common"
	"strings"
)

func GetInstance() IDistro {
	m := *(common.ReadOSRelease())

	linuxBase := strings.ToLower(m["ID_LIKE"])

	if linuxBase == "debian" {
		debianBaseDistro := DebianBased{}
		return &debianBaseDistro
	} else if strings.Contains(linuxBase, "arch") {
		archBaseDistro := ArchLinuxBased{}
		return &archBaseDistro
	} else if strings.Contains(linuxBase, "rehl") {
		redHatBaseDistro := RedHatBased{}
		return &redHatBaseDistro
	} else if strings.Contains(linuxBase, "darwin") {
		macBaseDistro := MacBased{}
		return &macBaseDistro
	} else {
		fmt.Println("Warning: debian is using because this distro is not supported, contact admin")
		debianBaseDistro := DebianBased{}
		return &debianBaseDistro
	}

}
