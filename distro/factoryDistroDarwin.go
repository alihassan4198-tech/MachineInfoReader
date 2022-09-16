//go:build darwin
// +build darwin

package distro

import "fmt"

func GetInstance() IDistro {
	fmt.Println("in GetInstance Start")
	macBaseDistro := MacBased{}
	fmt.Println("in GetInstace end")
	return &macBaseDistro
}
