package model

type ComputerBaseboardType struct {
	Computer_name     string
	Caption           string
	Configoptions     []string
	Creationclassname string
	Description       string
	Installdate       string
	Manufacturer      string
	Model             string
	Name              string
	Partnumber        string
	Poweredon         bool
	Product           string
	Serialnumber      string
	Sku               string
	Status            string
	Tag               string
	Version           string
}

type ComputerBiosType struct {
	Manufacturer           string
	Name                   string
	Serialnumber           string
	Smbiosbiosversion      string
	Version                string
	Biosversion            string
	Caption                string
	Currentlanguage        string
	Description            string
	Installablelanguages   int
	Installdate            string
	Listoflanguages        []string
	Primarybios            bool
	Releasedate            string
	Smbiosmajorversion     string
	Smbiosminorversion     string
	Smbiospresent          bool
	Softwareelementid      string
	Softwareelementstate   int
	Status                 string
	Systembiosmajorversion int
	Systembiosminorversion int
	Targetoperatingsystem  int
}

// ComputerCPU SUB TYPE
type ComputerCPU struct {
	Device_id                     int
	Manufacturer                  string
	Max_clock_speed               string
	Name                          string
	Socket_designation            string
	Virtualizationfirmwareenabled bool
}

type ComputerCPUType struct {
	Caption  string
	CPUCores []ComputerCPU
}

// EndpointProtection SUB TYPE
type EndpointProtectionSoftwareType struct {
	Type       string
	Name       string
	State      string
	Db_status  string
	Time_stamp int
	Is_default string
}

// SUB TYPE

type ComputerEndpointProtectionType struct {
	Softwares []EndpointProtectionSoftwareType
}

// FirewallRules SUB TYPE
type FirewallRuleType struct {
	Name           string
	Enabled        string
	Direction      string
	Profiles       string
	Grouping       string
	Local_ip       string
	Remote_ip      string
	Protocol       string
	Local_port     string
	Remote_port    string
	Edge_traversal string
	Action         string
}

// SUB TYPE

type ComputerFirewallRulesType struct {
	Total_rules  int
	Active_state string
	FW_rules     []FirewallRuleType
}

type ComputerNICType struct {
	Caption                 string
	Model                   string
	Default_ip_gateway      []string
	Dhcp_server             string
	Dns_server_search_order []string
	Ipaddress               []string
	Ip_subnet               []string
	Mac_address             string
	Ip_enabled              bool
}

type ComputerOSType struct {
	Computer_name   string
	Caption         string
	Os_architecture string
	Os_version      string
	Lastbootuptime  int
	Lts             bool
}

// Services SUB TYPE
type ServiceType struct {
	Name         string
	Display_name string
	Process_id   int
	Start_mode   string
	State        string
	Status       string
}

// SUB TYPE

type ComputerServicesType struct {
	TotalServciesRunning int
	Services             []ServiceType
}

// Software Installed SUB TYPE
type SoftwareInstalledType struct {
	Display_name string
	Product_code string
	Version      string
}

// SUB TYPE

type ComputerSoftwaresInstalledType struct {
	Total_software     int
	SoftwaresInstalled []SoftwareInstalledType
}

type ComputerSystemType struct {
	Domain               string
	Manufacturer         string
	Model                string
	Chassis_sku_number   string
	Pc_system_type       int64
	System_family        string
	System_sku_number    string
	System_type          string
	Primary_owner_name   string
	Total_phsical_memory int64
}

type ComputerPatchesType struct {
	Total_number_of_upates int
	SecurityUpdates        int
	Patches                []string
	Patch_name             string
	Patch_version          string
}

type ComputerInfoType struct {
	ComputerBaseboard          ComputerBaseboardType
	ComputerBios               ComputerBiosType
	ComputerCPU                ComputerCPUType
	ComputerEndpointProtection ComputerEndpointProtectionType
	ComputerFirewallRules      ComputerFirewallRulesType
	ComputerNICS               []ComputerNICType
	ComputerOS                 ComputerOSType
	ComputerServices           ComputerServicesType
	ComputerSoftwaresInstalled ComputerSoftwaresInstalledType
	ComputerSystem             ComputerSystemType
	ComputerPatches            ComputerPatchesType
}
