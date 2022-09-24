//go:build darwin
// +build darwin

package distro

import (
	"fmt"
	"machine_info_gatherer/common"
	"machine_info_gatherer/distro/systemprofiler"
	"machine_info_gatherer/model"
	"os"
	"strconv"
	"strings"
)

type MacBased struct {
	DarwinBase
	// LinuxBase
}

// Mac

// const (
// 	find0 = "Name"
// 	find1 = "Version"
// 	find2 = "Architecture"
// 	find3 = "Description"
// )

var infoMap systemprofiler.DarwinSystemProfilerInfo

// var infoMap model.ComputerInfoType

func MacGetAllInfoInMap() {
	infoMap = systemprofiler.DarwinSystemProfiler()
	// infoMap = model.ComputerInfoType
}

func (mb *MacBased) DistroGatherInfo() (*model.ComputerInfoType, error) {

	return &model.ComputerInfoType{}, nil
}

func (mb *MacBased) DistroGetComputerBaseboard() (*model.ComputerBaseboardType, error) {

	cb := model.ComputerBaseboardType{}

	hw := len(infoMap.SPHardware.SpHardwareDataType) > 0
	sw := len(infoMap.SPSoftware.SpSoftwareDataType) > 0

	var err error
	computerName, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}

	cb.Computer_name = computerName
	cb.Caption = baseboardCaption

	cb.Creationclassname = ""

	if hw {
		cb.Serialnumber = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].SerialNumber)
		cb.Manufacturer = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].MachineName)
		cb.Tag = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].PlatformUuid)
		cb.Model = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].MachineModel)
	}

	if sw {
		cb.Version = common.RootNeeded(infoMap.SPSoftware.SpSoftwareDataType[0].OSVersion)
		cb.Product = common.RootNeeded(infoMap.SPSoftware.SpSoftwareDataType[0].KernelVersion)
		cb.Name = common.RootNeeded(infoMap.SPSoftware.SpSoftwareDataType[0].LocalHostName)
		cb.Status = common.RootNeeded(infoMap.SPSoftware.SpSoftwareDataType[0].BootMode)
	}

	date, err := common.RunFullCommand("ls -l /var/db/.AppleSetupDone")
	if err != nil {
		date = ""
	} else {
		split := strings.Split(date, " ")
		if len(split) < 8 {
			date = " "
		} else {
			date = strings.Join(split[8:12], " ")
		}
	}

	cb.Installdate = date

	cb.Description = "Baseboard Info"

	cb.Poweredon = true

	return &cb, nil
}

func (mb *MacBased) DistroGetComputerBios() (*model.ComputerBiosType, error) {

	cbios := model.ComputerBiosType{}

	hw := len(infoMap.SPHardware.SpHardwareDataType) > 0
	sw := len(infoMap.SPSoftware.SpSoftwareDataType) > 0

	if hw {
		cbios.Manufacturer = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].MachineName)
		cbios.Name = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].Name)
		cbios.Serialnumber = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].SerialNumber)
		cbios.Smbiosbiosversion = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].SmcVersionSystem)
		cbios.Biosversion = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].BootRomVersion)
	}

	if sw {
		cbios.Version = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].SmcVersionSystem)
	}

	cbios.Currentlanguage = infoMap.SPInternational.SpInternationalDataType[1].SystemLanguages[0]

	output := infoMap.SPInternational.SpInternationalDataType[1].SystemInterfaceLanguages
	var a string
	for index, i := range output {
		concatenated := fmt.Sprint(index, ") ", i, "  ")
		a += concatenated
	}
	cbios.Listoflanguages = strings.ReplaceAll(a, "  ", "\n")

	cbios.Status = ""
	cbios.Releasedate = ""
	cbios.Softwareelementid = ""
	cbios.Softwareelementstate = 0
	cbios.Targetoperatingsystem = 0
	cbios.Systembiosminorversion = 0
	cbios.Smbiospresent = true
	cbios.Installablelanguages = len(cbios.Listoflanguages)
	cbios.Description = "Bios Info"
	cbios.Primarybios = true
	cbios.Caption = biosCaption

	maj, min, found := strings.Cut(cbios.Smbiosbiosversion, ".")

	if found {
		cbios.Smbiosmajorversion = maj
		cbios.Smbiosminorversion = min
	}

	if cbios.Biosversion != "" {
		cbios.Status = "Installed"
	}

	imaj, err := strconv.Atoi(maj)
	if err != nil {
		fmt.Println(err)
	} else {
		cbios.Systembiosmajorversion = imaj
	}

	return &cbios, nil
}

func (mb *MacBased) DistroGetComputerCPU() (*model.ComputerCPUType, error) {

	compCpu := model.ComputerCPUType{}

	compCpu.Caption = cpuCaption

	hw := len(infoMap.SPHardware.SpHardwareDataType) > 0
	processor := len(infoMap.SPHardware.SpHardwareDataType)

	c := model.ComputerCPU{}
	if hw {
		c.Manufacturer = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].MachineName)
		c.Name = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].Name)
		c.Max_clock_speed = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].CurrentProcessorSpeed)
		compCpu.No_of_cores = infoMap.SPHardware.SpHardwareDataType[0].NumberProcessors
	}

	if hw && processor < infoMap.SPHardware.SpHardwareDataType[0].NumberProcessors {
		for p := 0; p < compCpu.No_of_cores; p++ {
			c.Device_id = p
			compCpu.CPUCores = append(compCpu.CPUCores, c)
		}
	} else {
		compCpu.CPUCores = append(compCpu.CPUCores, c)
	}
	return &compCpu, nil
}

func (mb *MacBased) DistroGetComputerEndpointProtectionSoftwares() (*model.ComputerEndpointProtectionType, error) {

	epsoft := model.ComputerEndpointProtectionType{}
	e := model.EndpointProtectionSoftwareType{}
	e.Type = ""
	e.Name = ""
	e.State = ""
	e.Db_status = ""
	e.Time_stamp = 0
	e.Is_default = ""
	epsoft.Softwares = append(epsoft.Softwares, e)
	return &epsoft, nil
}

func (mb *MacBased) DistroGetComputerFirewallRules() (*model.ComputerFirewallRulesType, error) {

	cfwRules := model.ComputerFirewallRulesType{}
	cfwRules.Active_state = common.RootNeeded(infoMap.SPFirewall.SpFirewallDataType[0].SpfirewallGlobalstate)
	return &cfwRules, nil
}

func (mb *MacBased) DistroGetComputerNIC() (*[]model.ComputerNICType, error) {

	comNic := []model.ComputerNICType{}

	var networkNic model.ComputerNICType
	networkNic.Caption = "NetworkNic"
	networkNic.Model = infoMap.SPNetwork.SpNetworkDataType[0].Name
	comNic = append(comNic, networkNic)

	var ipv4Nic model.ComputerNICType
	ipv4Nic.Caption = "IPv4"
	ipv4Nic.Default_ip_gateway = []string{infoMap.SPNetwork.SpNetworkDataType[0].IPv4.ArpResolvedIpAddress}
	ipv4Nic.Mac_address = infoMap.SPNetwork.SpNetworkDataType[0].IPv4.ArpResolvedHardwareAddress
	ipv4Nic.Ipaddress = append(ipv4Nic.Ipaddress, infoMap.SPNetwork.SpNetworkDataType[0].IPv4.AdditionalRoutes[0].DestinationAddress, infoMap.SPNetwork.SpNetworkDataType[0].IPv4.AdditionalRoutes[1].DestinationAddress)
	ipv4Nic.Ip_subnet = infoMap.SPNetwork.SpNetworkDataType[0].IPv4.SubnetMasks
	if len(ipv4Nic.Ipaddress) > 0 {
		ipv4Nic.Ip_enabled = true
	}
	comNic = append(comNic, ipv4Nic)

	var dhcpNic model.ComputerNICType
	dhcpNic.Caption = "Dhcp"
	dhcpNic.Dhcp_server = infoMap.SPNetwork.SpNetworkDataType[0].Dhcp.DhcpDomainNameServers
	dhcpNic.Ipaddress = []string{infoMap.SPNetwork.SpNetworkDataType[0].Dhcp.DhcpRouters}
	dhcpNic.Ip_subnet = []string{infoMap.SPNetwork.SpNetworkDataType[0].Dhcp.DhcpSubnetMask}
	if len(dhcpNic.Ipaddress) > 0 {
		dhcpNic.Ip_enabled = true
	}
	comNic = append(comNic, dhcpNic)

	var dnsNic model.ComputerNICType
	dnsNic.Caption = "DNS"
	dnsNic.Ipaddress = infoMap.SPNetwork.SpNetworkDataType[0].Dns.ServerAddresses
	if len(dhcpNic.Ipaddress) > 0 {
		dnsNic.Ip_enabled = true
	}
	comNic = append(comNic, dnsNic)

	return &comNic, nil
}

func (mb *MacBased) DistroGetComputerOS() (*model.ComputerOSType, error) {

	comOS := model.ComputerOSType{}

	comOS.Computer_name = common.RootNeeded(infoMap.SPSoftware.SpSoftwareDataType[0].LocalHostName)
	comOS.Os_name = common.RootNeeded(infoMap.SPSoftware.SpSoftwareDataType[0].OSVersion)
	comOS.Vendor = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].MachineName)
	comOS.Caption = osCaption
	comOS.Os_version = common.RootNeeded(infoMap.SPSoftware.SpSoftwareDataType[0].OSVersion)
	comOS.Lastbootuptime = common.RootNeeded(infoMap.SPSoftware.SpSoftwareDataType[0].Uptime)

	return &comOS, nil
}

func (mb *MacBased) DistroGetComputerServices() (*model.ComputerServicesType, error) {

	comServ := model.ComputerServicesType{}
	comS := model.Service{}
	cmd, err := common.RunFullCommand("launchctl list")
	if err != nil {
		fmt.Println(err)
	} else {
		cmdOutput := strings.Split(string(cmd), "\n")
		for i, svc := range cmdOutput {
			if i == 0 {
				continue
			}
			if svc == "" {
				break
			}
			svc = strings.ReplaceAll(svc, "\t", "  ")
			splitedSvc := strings.Split(svc, "  ")
			comS.Name = splitedSvc[2]
			comS.Process_id, _ = strconv.Atoi(splitedSvc[0])
			comS.Status = splitedSvc[1]
			if comS.Process_id != 0 {
				comServ.TotalServciesRunning++
			}
			comServ.Services = append(comServ.Services, comS)
		}
	}
	return &comServ, nil
}

func (mb *MacBased) DistroGetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalledType, error) {

	sit := model.ComputerSoftwaresInstalledType{}
	si := model.SoftwareInstalledType{}
	for _, a := range infoMap.SPApplications.SpApplicationsDataType {
		si.Display_name = a.Name
		si.Version = a.Version
		sit.SoftwaresInstalled = append(sit.SoftwaresInstalled, si)
	}
	totalsit := len(sit.SoftwaresInstalled)
	sit.Total_software = totalsit
	sit.SoftwaresInstalled = append(sit.SoftwaresInstalled, si)

	return &sit, nil
}

func (mb *MacBased) DistroGetComputerSystem() (*model.ComputerSystemType, error) {

	comSys := model.ComputerSystemType{}

	comSys.Manufacturer = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].MachineName)
	comSys.Model = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].MachineModel)
	comSys.Total_phsical_memory = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].PhysicalMemory)
	comSys.Primary_owner_name = common.RootNeeded(infoMap.SPHardware.SpHardwareDataType[0].MachineName)

	cmd, err := common.RunFullCommand("scutil --dns | grep domain")
	if err != nil {
		fmt.Println(err)
	} else {
		s := strings.ReplaceAll(cmd, "  domain   :", ",")
		s1 := strings.ReplaceAll(s, "\n", "")
		s2 := strings.Replace(s1, ",", "", 1)
		s2 = strings.TrimSpace(s2)
		comSys.Domain = s2
	}
	return &comSys, nil
}

func (mb *MacBased) DistroGetComputerPatches() (*model.ComputerPatchesType, error) {
	comPatch := model.ComputerPatchesType{}
	comPatch.Total_number_of_updates = 0
	comPatch.SecurityUpdates = 0
	comPatch.Patches = []string{}
	comPatch.Patch_name = ""
	comPatch.Patch_version = ""
	return &comPatch, nil
}
