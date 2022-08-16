//go:build darwin
// +build darwin

package distro

import (
	"fmt"
	"machine_info_gatherer/common"
	"machine_info_gatherer/distro/systemprofiler"
	"machine_info_gatherer/model"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Mac

const (
	find0 = "Name"
	find1 = "Version"
	find2 = "Architecture"
	find3 = "Description"
)

type MacBased struct {
	LinuxBase
}

var infoMap systemprofiler.DarwinSystemProfilerInfo

func MacGetAllInfoInMap() {
	infoMap = systemprofiler.DarwinSystemProfiler()
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
	cbios.Listoflanguages = infoMap.SPInternational.SpInternationalDataType[1].SystemInterfaceLanguages
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
	// imin, err := strconv.Atoi(min)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	cbios.Systembiosminorversion = imin
	// }

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
		c.No_of_cores = infoMap.SPHardware.SpHardwareDataType[0].NumberProcessors
	}

	if hw && processor < infoMap.SPHardware.SpHardwareDataType[0].NumberProcessors {
		for p := 0; p < c.No_of_cores; p++ {
			compCpu.CPUCores = append(compCpu.CPUCores, c)
		}
	} else {
		compCpu.CPUCores = append(compCpu.CPUCores, c)
	}
	return &compCpu, nil
}

// THIS FUNCTION IS PENDING---------------------------------------------

func (mb *MacBased) DistroGetComputerEndpointProtectionSoftwares() (*model.ComputerEndpointProtectionType, error) {

	epsoft := model.ComputerEndpointProtectionType{}

	return &epsoft, nil
}

func (mb *MacBased) DistroGetComputerFirewallRules() (*model.ComputerFirewallRulesType, error) {

	cfwRules := model.ComputerFirewallRulesType{}
	cfwRules.Active_state = common.RootNeeded(infoMap.SPFirewall.SpFirewallDataType[0].SpfirewallGlobalstate)

	cfw := model.FirewallRule{}
	cfw.Enabled = common.RootNeeded(infoMap.SPFirewall.SpFirewallDataType[0].SpfirewallStealthenabled)

	cfwRules.FW_rules = append(cfwRules.FW_rules, cfw)
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
	comNic = append(comNic, ipv4Nic)

	var dhcpNic model.ComputerNICType
	dhcpNic.Caption = "Dhcp"
	dhcpNic.Dhcp_server = infoMap.SPNetwork.SpNetworkDataType[0].Dhcp.DhcpDomainNameServers
	comNic = append(comNic, dhcpNic)

	var dnsNic model.ComputerNICType
	dnsNic.Caption = "DNS"
	dnsNic.Ipaddress = infoMap.SPNetwork.SpNetworkDataType[0].Dns.ServerAddresses
	comNic = append(comNic, dnsNic)

	var etherNic model.ComputerNICType
	etherNic.Caption = "Ethernet"
	comNic = append(comNic, etherNic)

	return &comNic, nil
}

func (mb *MacBased) DistroGetComputerOS() (*model.ComputerOSType, error) {

	comOS := model.ComputerOSType{}

	comOS.Computer_name = common.RootNeeded(infoMap.SPSoftware.SpSoftwareDataType[0].LocalHostName)
	comOS.Os_name = common.RootNeeded(infoMap.SPSoftware.SpSoftwareDataType[0].Name)
	comOS.Caption = osCaption
	comOS.Os_version = common.RootNeeded(infoMap.SPSoftware.SpSoftwareDataType[0].OSVersion)
	comOS.Lastbootuptime = common.RootNeeded(infoMap.SPSoftware.SpSoftwareDataType[0].Uptime)

	return &comOS, nil
}

//--------THIS FUNCTION IS PENDING---------------------
func (mb *MacBased) DistroGetComputerServices() (*model.ComputerServicesType, error) {

	comServ := model.ComputerServicesType{}

	comServ.TotalServciesRunning = len(comServ.Services)

	return &comServ, nil
}

func (mb *MacBased) DistroGetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalledType, error) {
	sit := model.ComputerSoftwaresInstalledType{}
	si := model.SoftwareInstalledType{}
	for _, a := range infoMap.SPApplications.SpApplicationsDataType {
		// a = append(a, si.Display_name, si.Version)
		si.Display_name = a.Name
		si.Version = a.Version
		sit.SoftwaresInstalled = append(sit.SoftwaresInstalled, si)
	}

	return &sit, nil
}

func (mb *MacBased) DistroGetComputerSystem() (*model.ComputerSystemType, error) {

	comSys := model.ComputerSystemType{}

	domainName, err := exec.Command("domainname").Output()
	if err != nil {
		fmt.Println(err)
	}
	comSys.Domain = strings.TrimSpace(string(domainName))

	manufacturer, err := common.RunFullCommandWithSudo("dmidecode -s baseboard-manufacturer")
	if err != nil {
		fmt.Println(err)
		comSys.Manufacturer = common.NeedSudoPreviliges(err)
	} else {
		comSys.Manufacturer = manufacturer
	}

	model, err := common.RunFullCommandWithSudo("dmidecode -s baseboard-product-name")
	if err != nil {
		fmt.Println(err)
		comSys.Model = common.NeedSudoPreviliges(err)
	} else {
		comSys.Model = model
	}

	return &comSys, nil
}

func (mb *MacBased) DistroGetComputerPatches() (*model.ComputerPatchesType, error) {
	comPatch := model.ComputerPatchesType{}

	return &comPatch, nil
}

func (mb *MacBased) GetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalledType, error) {
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
