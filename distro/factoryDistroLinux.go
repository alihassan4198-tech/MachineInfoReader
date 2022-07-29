//go:build linux
// +build linux

package distro

import (
	"fmt"
	"machine_info_gatherer/common"
	"strings"
)

func GetInstance() IDistro {
	debianBaseDistro := DebianBased{}
	return &debianBaseDistro
}
