package model

type ComputerBaseboard struct {
	Computer_name     string
	Caption           string
	Configoptions     []string
	Creationclassname string
	Description       string
	Installdate       int64
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

type ComputerBios struct {
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
	Releasedate            int64
	Smbiosmajorversion     int
	Smbiosminorversion     int
	Smbiospresent          bool
	Softwareelementid      string
	Softwareelementstate   int
	Status                 string
	Systembiosmajorversion int
	Systembiosminorversion int
	Targetoperatingsystem  int
}

type ComputerCPU struct {
	Caption                       string
	Device_id                     string
	Manufacturer                  string
	Max_clock_speed               int
	Name                          string
	Socket_designation            string
	Virtualizationfirmwareenabled bool
}

// EndpointProtection SUB TYPE
type EndpointProtectionSoftware struct {
	Type       string
	Name       string
	State      string
	Db_status  string
	Time_stamp int
	Is_default string
}

// SUB TYPE

type ComputerEndpointProtection struct {
	Softwares []EndpointProtectionSoftware
}

// FirewallRules SUB TYPE
type FirewallRule struct {
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

type ComputerFirewallRules struct {
	Total_rules  int
	Active_state string
	FW_rules     []FirewallRule
}

type ComputerNIC struct {
	Caption                 int
	Model                   string
	Default_ip_gateway      []string
	Dhcp_server             string
	Dns_server_search_order []string
	Ipaddress               []string
	Ip_subnet               []string
	Mac_address             string
	Ip_enabled              bool
}

type ComputerOS struct {
	Computer_name   string
	Caption         string
	Os_architecture string
	Os_version      string
	Lastbootuptime  int
	Lts             bool
}

// Services SUB TYPE
type Service struct {
	Name         string
	Display_name string
	Process_id   int
	Start_mode   string
	State        string
	Status       string
}

// SUB TYPE

type ComputerServices struct {
	TotalServciesRunning int
	Service
}

// Software Installed SUB TYPE
type SoftwareInstalled struct {
	Display_name string
	Product_code string
	Version      string
}

// SUB TYPE

type ComputerSoftwaresInstalled struct {
	Total_software int
	SoftwareInstalled
}

type ComputerSystem struct {
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

type ComputerPatches struct {
	Total_number_of_upates int
	SecurityUpdates        int
	Patches                []string
	Patch_name             string
	Patch_version          string
}

type ComputerInfo struct {
	ComputerBaseboard
	ComputerBios
	ComputerCPU
	ComputerEndpointProtection
	ComputerFirewallRules
	ComputerNICS []ComputerNIC
	ComputerOS
	ComputerServices
	ComputerSoftwaresInstalled
	ComputerSystem
	ComputerPatches
}
