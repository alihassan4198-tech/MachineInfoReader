package distro

import (
	"fmt"
	"machine_info_gatherer/gatherer"
	"strings"
)

func GetInstance() IDistro {
	m := *(gatherer.ReadOSRelease())

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
	} else {
		fmt.Println("Warning: debian is using because this distro is not supported, contact admin")
		debianBaseDistro := DebianBased{}
		return &debianBaseDistro
	}

}
