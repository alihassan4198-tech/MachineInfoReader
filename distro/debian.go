//go:build linux
// +build linux

package distro

import (
	"fmt"
	"machine_info_gatherer/common"
	"machine_info_gatherer/model"
	"os"
	"os/exec"
	"strings"

	"github.com/alihassan4198-tech/ghw"
	"github.com/coreos/go-iptables/iptables"
	"github.com/zcalusic/sysinfo"
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

func MacGetAllInfoInMap() {
}

func (db *DebianBased) DistroGatherInfo() (*model.ComputerInfoType, error) {
	// var info model.ComputerInfoType

	// var err error

	// info.ComputerBaseboard , err =  db.DistroGetComputerBaseboard()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// info.ComputerBios, err = db.DistroGetComputerBios()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// info.ComputerCPU, err = db.DistroGetComputerCPU()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// info.ComputerEndpointProtection, err = db.DistroGetComputerEndpointProtectionSoftwares()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// info.ComputerFirewallRules, err = db.DistroGetComputerFirewallRules()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// info.ComputerNICS, err = db.DistroGetComputerNIC()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// info.ComputerOS, err = db.DistroGetComputerOS()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// info.ComputerServices, err = db.DistroGetComputerServices()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// info.ComputerSoftwaresInstalled, err = db.DistroGetComputerSoftwaresInstalled()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// info.ComputerPatches, err = db.DistroGetComputerPatches()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	return &model.ComputerInfoType{}, nil
}

func (db *DebianBased) DistroGetComputerBios() (*model.ComputerBiosType, error) {
	cbios := model.ComputerBiosType{}
	// fmt.Println("********************************")
	// godmi.Init()
	// macBios := godmi.GetBIOSInformation()
	// fmt.Println(macBios)
	// fmt.Println("********************************")
	bios, err := ghw.BIOS()
	if err != nil {
		fmt.Printf("Error getting BIOS info: %v", err)
	}

	cbios.Name = bios.Name
	cbios.Biosversion = common.RootNeeded(bios.Version)
	cbios.Version = cbios.Biosversion
	cbios.Manufacturer = common.RootNeeded(bios.Vendor)
	cbios.Installdate = common.RootNeeded(bios.Date)
	cbios.Serialnumber = common.RootNeeded(bios.Serialnumber)
	cbios.Installdate = common.RootNeeded(bios.Installdate)
	cbios.Primarybios = true
	cbios.Caption = biosCaption
	maj, min, found := strings.Cut(cbios.Biosversion, ".")

	if found {
		cbios.Smbiosmajorversion = maj
		cbios.Smbiosminorversion = min
	}

	if cbios.Biosversion != "" {
		cbios.Status = "Installed"
	}

	return &cbios, nil
}

func (db *DebianBased) DistroGetComputerCPU() (*model.ComputerCPUType, error) {
	compCpu := model.ComputerCPUType{}

	cpu, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}

	compCpu.Caption = cpuCaption

	for _, proc := range cpu.Processors {
		c := model.ComputerCPU{}

		c.Manufacturer = proc.Vendor
		c.Max_clock_speed = proc.CPUMhz
		c.Name = proc.Name
		c.Device_id = proc.ID

		compCpu.CPUCores = append(compCpu.CPUCores, c)
	}

	return &compCpu, nil
}

func (db *DebianBased) DistroGetComputerEndpointProtectionSoftwares() (*model.ComputerEndpointProtectionType, error) {
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

func (db *DebianBased) DistroGetComputerFirewallRules() (*model.ComputerFirewallRulesType, error) {

	cfwRules := model.ComputerFirewallRulesType{}

	tables := []string{"filter", "mangle", "nat", "raw"}

	ipt, err := iptables.New()
	if err != nil {
		fmt.Println(err)
	}

	for _, table := range tables {
		chains, err := ipt.ListChains(table)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Chains Length", len(chains))
		for _, chain := range chains {
			structStat, err := ipt.StructuredStats(table, chain)
			if err != nil {
				fmt.Println("parsing error: ", err)
				continue
			}
			for _, s := range structStat {
				cfwRule := model.FirewallRule{}

				cfwRule.TableName = table
				cfwRule.ChainName = chain
				cfwRule.Action = s.Target
				cfwRule.Enabled = "enable"
				cfwRule.Local_ip = s.Source.IP.String()
				cfwRule.Local_port = s.Input
				cfwRule.Remote_ip = s.Destination.IP.String()
				cfwRule.Remote_port = s.Output
				cfwRule.Protocol = s.Protocol
				cfwRule.Direction = chain
				// fmt.Printf("%#v\n", cfwRule)
				cfwRules.FW_rules = append(cfwRules.FW_rules, cfwRule)
			}
		}
	}

	return &cfwRules, nil
}

func (db *DebianBased) DistroGetComputerNIC() (*[]model.ComputerNICType, error) {

	comNic := []model.ComputerNICType{}
	net, err := ghw.Network()
	if err != nil {
		fmt.Printf("Error getting network info: %v", err)
	}

	for _, nic := range net.NICs {

		comNic = append(comNic, model.ComputerNICType{
			Caption:     nic.Name,
			Mac_address: nic.MacAddress,
		})

	}

	return &comNic, nil
}

func (db *DebianBased) DistroGetComputerOS() (*model.ComputerOSType, error) {

	comOS := model.ComputerOSType{}
	var si sysinfo.SysInfo
	si.GetSysInfo()
	Os := si.OS
	cname, err := os.Hostname()
	if err != nil {
		fmt.Printf("error:%#v", err)
	}

	comOS.Computer_name = cname
	m := *(common.ReadOSRelease())

	comOS.Os_version = m["VERSION_CODENAME"] + " " + m["VERSION_ID"]
	comOS.Os_name = Os.Name
	comOS.Vendor = Os.Vendor
	comOS.Caption = osCaption
	comOS.Os_architecture = Os.Architecture
	comOS.Os_version = Os.Version
	comOS.Release = Os.Release
	comOS.Lts = false

	return &comOS, nil
}

func (db *DebianBased) DistroGetComputerServices() (*model.ComputerServicesType, error) {

	comServ := model.ComputerServicesType{}

	cmd, err := exec.Command("systemctl", "--type=service").Output()
	if err != nil {
		fmt.Println(err)
	}
	cmdOutput := strings.Split(string(cmd), "\n")

	for _, svc := range cmdOutput {
		if common.IsService(svc) && common.IsServiceRunning(svc) {
			comServ.Services = append(comServ.Services, *common.ParseService(svc))
		}
	}

	comServ.TotalServciesRunning = len(comServ.Services)

	return &comServ, nil
}

func (db *DebianBased) DistroGetComputerSystem() (*model.ComputerSystemType, error) {

	comSys := model.ComputerSystemType{}

	domainName, err := exec.Command("domainname").Output()
	if err != nil {
		fmt.Println(err)
	}
	comSys.Domain = strings.TrimSpace(string(domainName))

	comSys.Total_phsical_memory, err = common.RunFullCommand("grep MemTotal /proc/meminfo | awk '{print $2}'")
	if err != nil {
		fmt.Println(err)
	}

	// a, _ := common.RunFullCommand("uname")
	// comSys.Pc_system_type, err = strconv.ParseInt(a, 10, 64)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	baseboard, err := ghw.Baseboard()
	if err != nil {
		fmt.Println(err)
	}
	comSys.Manufacturer = common.RootNeeded(baseboard.Vendor)
	comSys.Model = common.RootNeeded(baseboard.Product)
	// manufacturer, err := common.RunFullCommandWithSudo("dmidecode -s baseboard-manufacturer")
	// if err != nil {
	// 	fmt.Println(err)
	// 	comSys.Manufacturer = common.NeedSudoPreviliges(err)
	// } else {
	// 	comSys.Manufacturer = manufacturer
	// }

	// model, err := common.RunFullCommandWithSudo("dmidecode -s baseboard-product-name")
	// if err != nil {
	// 	fmt.Println(err)
	// 	comSys.Model = common.NeedSudoPreviliges(err)
	// } else {
	// 	comSys.Model = model
	// }

	return &comSys, nil
}

func (db *DebianBased) DistroGetComputerPatches() (*model.ComputerPatchesType, error) {
	comPatch := model.ComputerPatchesType{}
	comPatch.Total_number_of_updates = 0
	comPatch.SecurityUpdates = 0
	comPatch.Patches = []string{}
	comPatch.Patch_name = ""
	comPatch.Patch_version = ""

	return &comPatch, nil
}

func (db *DebianBased) DistroGetComputerBaseboard() (*model.ComputerBaseboardType, error) {
	cb := model.ComputerBaseboardType{}

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

	cb.Serialnumber = common.RootNeeded(baseboard.SerialNumber)
	cb.Version = common.RootNeeded(baseboard.Version)
	cb.Product = common.RootNeeded(baseboard.Product)
	cb.Manufacturer = common.RootNeeded(baseboard.Vendor)
	cb.Tag = common.RootNeeded(baseboard.AssetTag)

	cb.Creationclassname, err = common.RunFullCommand("uname -m")
	if err != nil {
		cb.Creationclassname = ""
	}
	cb.Creationclassname = strings.ReplaceAll(cb.Creationclassname, "\n", "")

	installDate, err := common.RunFullCommand("ls -ld --time-style=full-iso /var/log/installer")
	if err != nil {
		cb.Installdate = ""
	} else {
		cb.Installdate = func(str string) string {
			res := strings.Split(str, "/var")
			res = strings.Split(res[0], " ")
			result := strings.Join(res[5:8], " ")
			return result
		}(installDate)
	}

	cb.Poweredon = true

	return &cb, nil
}

func (db *DebianBased) DistroGetComputerSoftwaresInstalled() (*model.ComputerSoftwaresInstalledType, error) {
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
