//go:build darwin
// +build darwin

package distro

func GetInstance() IDistro {

	macBaseDistro := MacBased{}
	return &macBaseDistro

}
