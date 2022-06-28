package gatherer

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"machine_info_gatherer/model"
	"os"
	"os/exec"
	"strings"

	"github.com/alihassan4198-tech/ghw"
	"github.com/quay/claircore/osrelease"
)

type BaseGatherer struct {
}

//--------------------------
// Baseboard Implementation
//--------------------------

const (
	baseboardCaption string = "Base Board"
	cpuCaption       string = "CPU"
	osCaption        string = "Computer OS"
)

func rootNeeded(arg string) string {
	if arg == "None" || arg == "unknown" {
		return arg + " (need root access)"
	} else {
		return arg
	}
}

func isService(svc string) bool {
	if strings.Contains(svc, "loaded act") || strings.Contains(svc, "loaded fail") {
		return true
	} else {
		return false
	}
}

func isServiceRunning(svc string) bool {
	if strings.Contains(svc, "running") {
		return true
	} else {
		return false
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

	cb.Serialnumber = rootNeeded(baseboard.SerialNumber)
	cb.Version = rootNeeded(baseboard.Version)
	cb.Product = rootNeeded(baseboard.Product)
	cb.Manufacturer = rootNeeded(baseboard.Vendor)
	cb.Tag = rootNeeded(baseboard.AssetTag)

	return &cb
}

func (bg *BaseGatherer) GetComputerBios() *model.ComputerBios {
	cbios := model.ComputerBios{}

	bios, err := ghw.BIOS()
	if err != nil {
		fmt.Printf("Error getting BIOS info: %v", err)
	}

	cbios.Biosversion = rootNeeded(bios.Version)
	cbios.Manufacturer = rootNeeded(bios.Vendor)
	cbios.Installdate = rootNeeded(bios.Date)

	return &cbios
}

func (bg *BaseGatherer) GetComputerCPU() *model.ComputerCPU {
	compCpu := model.ComputerCPU{}

	_, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}

	compCpu.Caption = cpuCaption

	return &compCpu
}

func (bg *BaseGatherer) GetComputerEndpointProtectionSoftwares() *model.ComputerEndpointProtection {

	epsoft := model.ComputerEndpointProtection{}

	return &epsoft
}

func (bg *BaseGatherer) GetComputerFirewallRules() *model.ComputerFirewallRules {

	cfwRules := model.ComputerFirewallRules{}

	return &cfwRules
}

func (bg *BaseGatherer) GetComputerNIC() *[]model.ComputerNIC {

	comNic := []model.ComputerNIC{}
	net, err := ghw.Network()
	if err != nil {
		fmt.Printf("Error getting network info: %v", err)
	}

	for _, nic := range net.NICs {

		comNic = append(comNic, model.ComputerNIC{
			Caption:     nic.Name,
			Mac_address: nic.MacAddress,
		})

	}

	return &comNic
}

func ReadOSRelease() *map[string]string {
	ctx := context.Background()
	var b []byte

	sys := os.DirFS("/")

	// Look for an os-release file.
	b, err := fs.ReadFile(sys, osrelease.Path)
	if err != nil {
		fmt.Printf("error:%#v", err)
	}
	m, err := osrelease.Parse(ctx, bytes.NewReader(b))
	if err != nil {
		fmt.Printf("error:%#v", err)
	}
	return &m
}

func (bg *BaseGatherer) GetComputerOS() *model.ComputerOS {

	comOS := model.ComputerOS{}

	cname, err := os.Hostname()
	if err != nil {
		fmt.Printf("error:%#v", err)
	}

	comOS.Computer_name = cname
	comOS.Caption = osCaption

	m := *(ReadOSRelease())

	comOS.Os_version = m["VERSION_CODENAME"] + " " + m["VERSION_ID"]
	comOS.Lts = false

	return &comOS
}

func ParseService(svc string) *model.Service {
	service := model.Service{}

	svc = strings.ReplaceAll(svc, "●", " ")

	svc = strings.Join(strings.Fields(svc), " ")

	splitedSvc := strings.Split(svc, " ")

	service.Display_name = strings.TrimSpace(splitedSvc[0])
	service.State = strings.TrimSpace(splitedSvc[2])
	service.Status = strings.TrimSpace(splitedSvc[3])

	return &service
}

func (bg *BaseGatherer) GetComputerServices() *model.ComputerServices {

	comServ := model.ComputerServices{}

	cmd, err := exec.Command("systemctl", "--type=service").Output()
	if err != nil {
		fmt.Println(err)
	}
	cmdOutput := strings.Split(string(cmd), "\n")

	for _, svc := range cmdOutput {
		if isService(svc) && isServiceRunning(svc) {
			comServ.Services = append(comServ.Services, *ParseService(svc))
		}
	}

	comServ.TotalServciesRunning = len(comServ.Services)

	return &comServ
}

func (bg *BaseGatherer) GetComputerSoftwaresInstalled() *model.ComputerSoftwaresInstalled {
	comInsSoft := model.ComputerSoftwaresInstalled{}

	return &comInsSoft
}

//  All Info
func (bg *BaseGatherer) GatherInfo() *model.ComputerInfo {

	m := model.ComputerInfo{}

	m.ComputerBaseboard = *(bg.GetComputerBaseboard())
	m.ComputerBios = *(bg.GetComputerBios())
	m.ComputerCPU = *(bg.GetComputerCPU())
	m.ComputerEndpointProtection = *(bg.GetComputerEndpointProtectionSoftwares())
	m.ComputerFirewallRules = *(bg.GetComputerFirewallRules())
	m.ComputerNICS = *(bg.GetComputerNIC())
	m.ComputerOS = *(bg.GetComputerOS())
	m.ComputerServices = *(bg.GetComputerServices())
	m.ComputerSoftwaresInstalled = *(bg.GetComputerSoftwaresInstalled())

	return &m
}
