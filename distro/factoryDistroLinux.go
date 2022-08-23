//go:build linux
// +build linux

package distro

func GetInstance() IDistro {
	debianBaseDistro := DebianBased{}
	return &debianBaseDistro
}
