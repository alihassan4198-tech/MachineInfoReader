//go:build darwin
// +build darwin

package distro

import "time"

type SpParallelATAType struct {
	SPParallelATADataType []interface{} `json:"SPParallelATADataType"`
}

type SpUniversalAccessType struct {
	SPUniversalAccessDataType []struct {
		Name         string `json:"_name"`
		Contrast     string `json:"contrast"`
		CursorMag    string `json:"cursor_mag"`
		Display      string `json:"display"`
		FlashScreen  string `json:"flash_screen"`
		KeyboardZoom string `json:"keyboardZoom"`
		MouseKeys    string `json:"mouse_keys"`
		ScrollZoom   string `json:"scrollZoom"`
		SlowKeys     string `json:"slow_keys"`
		StickyKeys   string `json:"sticky_keys"`
		Voiceover    string `json:"voiceover"`
		ZoomMode     string `json:"zoomMode"`
	} `json:"SPUniversalAccessDataType"`
}

type SpSecureElementType struct {
	SPSecureElementDataType []interface{} `json:"SPSecureElementDataType"`
}

type SpApplicationsType struct {
	SpApplicationsDataType []struct {
		Name         string    `json:"_name"`
		ArchKind     string    `json:"arch_kind"`
		Info         string    `json:"info,omitempty"`
		LastModified time.Time `json:"lastModified"`
		ObtainedFrom string    `json:"obtained_from"`
		Path         string    `json:"path"`
		SignedBy     []string  `json:"signed_by,omitempty"`
		Version      string    `json:"version,omitempty"`
	} `json:"SPApplicationsDataType"`
}

type SpAudioType struct {
	SpAudioDataType []struct {
		Items []struct {
			Name                              string `json:"_name"`
			Properties                        string `json:"_properties,omitempty"`
			CoreaudioDefaultAudioInputDevice  string `json:"coreaudio_default_audio_input_device,omitempty"`
			CoreaudioDefaultAudioOutputDevice string `json:"coreaudio_default_audio_output_device,omitempty"`
			CoreaudioDefaultAudioSystemDevice string `json:"coreaudio_default_audio_system_device,omitempty"`
			CoreaudioDeviceInput              int    `json:"coreaudio_device_input,omitempty"`
			CoreaudioDeviceManufacturer       string `json:"coreaudio_device_manufacturer"`
			CoreaudioDeviceOutput             int    `json:"coreaudio_device_output,omitempty"`
			CoreaudioDeviceSrate              int    `json:"coreaudio_device_srate"`
			CoreaudioDeviceTransport          string `json:"coreaudio_device_transport"`
			CoreaudioInputSource              string `json:"coreaudio_input_source,omitempty"`
			CoreaudioOutputSource             string `json:"coreaudio_output_source,omitempty"`
		} `json:"_items"`
		Name string `json:"_name"`
	} `json:"SPAudioDataType"`
}

type SpBluetoothType struct {
	SpBluetoothDataType []struct {
		AppleBluetoothVersion    string `json:"apple_bluetooth_version"`
		IncomingSerialPortsTitle []struct {
			BluetoothIncomingPort struct {
				DeviceAuthentication string `json:"device_authentication"`
				DeviceChannel        string `json:"device_channel"`
			} `json:"Bluetooth-Incoming-Port"`
		} `json:"incoming_serial_ports_title"`
		LocalDeviceTitle struct {
			GeneralAddress                string `json:"general_address"`
			GeneralAutoseekKeyboard       string `json:"general_autoseek_keyboard"`
			GeneralAutoseekPointing       string `json:"general_autoseek_pointing"`
			GeneralDeviceClassComposite   string `json:"general_device_class_composite"`
			GeneralDeviceClassMajor       string `json:"general_device_class_major"`
			GeneralDeviceClassMinor       string `json:"general_device_class_minor"`
			GeneralFwVersion              string `json:"general_fw_version"`
			GeneralHardwareTransport      string `json:"general_hardware_transport"`
			GeneralHciRevision            string `json:"general_hci_revision"`
			GeneralHciVersion             string `json:"general_hci_version"`
			GeneralLmpSubversion          string `json:"general_lmp_subversion"`
			GeneralLmpVersion             string `json:"general_lmp_version"`
			GeneralMfg                    string `json:"general_mfg"`
			GeneralName                   string `json:"general_name"`
			GeneralPower                  string `json:"general_power"`
			GeneralProductID              string `json:"general_productID"`
			GeneralRemoteWake             string `json:"general_remoteWake"`
			GeneralServiceClass           string `json:"general_service_class"`
			GeneralSupportsHandoff        string `json:"general_supports_handoff"`
			GeneralSupportsInstantHotspot string `json:"general_supports_instantHotspot"`
			GeneralSupportsLowEnergy      string `json:"general_supports_lowEnergy"`
			GeneralTypeCompleteString     string `json:"general_type_complete_string"`
			GeneralTypeMajorString        string `json:"general_type_major_string"`
			GeneralVendorID               string `json:"general_vendorID"`
		} `json:"local_device_title"`
		ServicesTitle []struct {
			FileBrowsingTitle *struct {
				ServiceFTPRootFolder string `json:"service_FTPRootFolder"`
				ServiceOBEXReceive   string `json:"service_OBEXReceive"`
				ServiceState         string `json:"service_state"`
			} `json:"file_browsing_title,omitempty"`
			InternetSharingTitle *struct {
				ServiceState string `json:"service_state"`
			} `json:"internet_sharing_title,omitempty"`
			ObjectPushTitle *struct {
				ServiceOBEXFolder      string `json:"service_OBEXFolder"`
				ServiceOBEXHandleOther string `json:"service_OBEXHandle_Other"`
				ServiceOBEXReceive     string `json:"service_OBEXReceive"`
				ServiceState           string `json:"service_state"`
			} `json:"object_push_title,omitempty"`
		} `json:"services_title"`
	} `json:"SPBluetoothDataType"`
}

type SpCameraType struct {
	SpCameraDataType []interface{} `json:"SPCameraDataType"`
}

type SpCardReaderType struct {
	SpCardReaderDataType []interface{} `json:"SPCardReaderDataType"`
}

type SpiBridgeType struct {
	SPiBridgeDataType []interface{} `json:"SPiBridgeDataType"`
}

type SpDeveloperToolsType struct {
	SpDeveloperToolsDataType []struct {
		Name           string `json:"_name"`
		SpdevtoolsApps struct {
			SpinstrumentsApp string `json:"spinstruments_app"`
			SpxcodeApp       string `json:"spxcode_app"`
		} `json:"spdevtools_apps"`
		SpdevtoolsPath string `json:"spdevtools_path"`
		SpdevtoolsSdks struct {
			IOS struct {
				_13_0 string `json:"13.0"`
			} `json:"iOS"`
			MacOS struct {
				_10_15 string `json:"10.15"`
				_19_0  string `json:"19.0"`
			} `json:"macOS"`
			TvOS struct {
				_13_0 string `json:"13.0"`
			} `json:"tvOS"`
			WatchOS struct {
				_6_0 string `json:"6.0"`
			} `json:"watchOS"`
			// "iOS Simulator" cannot be unmarshalled into a struct field by encoding/json.
			// "tvOS Simulator" cannot be unmarshalled into a struct field by encoding/json.
			// "watchOS Simulator" cannot be unmarshalled into a struct field by encoding/json.
		} `json:"spdevtools_sdks"`
		SpdevtoolsVersion string `json:"spdevtools_version"`
	} `json:"SPDeveloperToolsDataType"`
}

type SpDiagnosticsType struct {
	SpDiagnosticsDataType []struct {
		Name              string    `json:"_name"`
		SpdiagsLastRunKey time.Time `json:"spdiags_last_run_key"`
		SpdiagsResultKey  string    `json:"spdiags_result_key"`
	} `json:"SPDiagnosticsDataType"`
}

type SpDisabledSoftwareType struct {
	SpDisabledSoftwareDataType []struct {
		Name         string `json:"_name"`
		DisabledDate string `json:"disabledDate"`
		Reason       string `json:"reason"`
		Version      string `json:"version"`
	} `json:"SPDisabledSoftwareDataType"`
}

type SpDiscBurningType struct {
	SpDiscBurningDataType []struct {
		Name             string `json:"_name"`
		BurnSupport      string `json:"burn_support"`
		DeviceCache      string `json:"device_cache"`
		DeviceCdwrite    string `json:"device_cdwrite"`
		DeviceDvdwrite   string `json:"device_dvdwrite"`
		DeviceMedia      string `json:"device_media"`
		DeviceReaddvd    string `json:"device_readdvd"`
		DeviceStrategies string `json:"device_strategies"`
		Firmware         string `json:"firmware"`
		Interconnect     string `json:"interconnect"`
		ProfilePath      string `json:"profile_path"`
	} `json:"SPDiscBurningDataType"`
}

type SpEthernetType struct {
	SpEthernetDataType []struct {
		Name                       string `json:"_name"`
		SpethernetBsdName          string `json:"spethernet_BSD_Name"`
		SpethernetBundleIDentifier string `json:"spethernet_BUNDLE_IDentifier"`
		SpethernetBus              string `json:"spethernet_bus"`
		Spethernet_deviceID        string `json:"spethernet_device-id"`
		SpethernetDeviceType       string `json:"spethernet_device_type"`
		SpethernetKextPath         string `json:"spethernet_kext_path"`
		Spethernet_linkWidth       string `json:"spethernet_link-width"`
		SpethernetName             string `json:"spethernet_name"`
		Spethernet_revisionID      string `json:"spethernet_revision-id"`
		Spethernet_vendorID        string `json:"spethernet_vendor-id"`
		SpethernetVersion          string `json:"spethernet_version"`
	} `json:"SPEthernetDataType"`
}

type SpExtensionsType struct {
	SpExtensionsDataType []struct {
		Name                    string    `json:"_name"`
		SpextArchitectures      []string  `json:"spext_architectures,omitempty"`
		SpextBundleid           string    `json:"spext_bundleid"`
		SpextHas64BitIntelCode  string    `json:"spext_has64BitIntelCode,omitempty"`
		SpextHasAllDependencies string    `json:"spext_hasAllDependencies"`
		SpextInfo               string    `json:"spext_info,omitempty"`
		SpextLastModified       time.Time `json:"spext_lastModified"`
		SpextLoadAddress        string    `json:"spext_load_address,omitempty"`
		SpextLoadable           string    `json:"spext_loadable"`
		SpextLoaded             string    `json:"spext_loaded"`
		SpextNotarized          string    `json:"spext_notarized"`
		SpextObtainedFrom       string    `json:"spext_obtained_from"`
		SpextPath               string    `json:"spext_path"`
		SpextRuntimeEnvironment string    `json:"spext_runtime_environment,omitempty"`
		SpextSignedBy           string    `json:"spext_signed_by"`
		SpextValidErrors        *struct {
			// "Validation Failures" cannot be unmarshalled into a struct field by encoding/json.
		} `json:"spext_valid_errors,omitempty"`
		SpextVersion string `json:"spext_version"`
		Version      string `json:"version"`
	} `json:"SPExtensionsDataType"`
}

type SpFibreChannelType struct {
	SpFibreChannelDataType []interface{} `json:"SPFibreChannelDataType"`
}

type SpFireWireType struct {
	SpFireWireDataType []interface{} `json:"SPFireWireDataType"`
}

type SpFirewallType struct {
	SpFirewallDataType []struct {
		Name                     string `json:"_name"`
		SpfirewallGlobalstate    string `json:"spfirewall_globalstate"`
		SpfirewallLoggingenabled string `json:"spfirewall_loggingenabled"`
		SpfirewallStealthenabled string `json:"spfirewall_stealthenabled"`
	} `json:"SPFirewallDataType"`
}

type SpFontsType struct {
	SpFontsDataType []struct {
		Name      string `json:"_name"`
		Enabled   string `json:"enabled"`
		Path      string `json:"path"`
		Type      string `json:"type"`
		Typefaces []struct {
			Name          string `json:"_name"`
			CopyProtected string `json:"copy_protected"`
			Copyright     string `json:"copyright,omitempty"`
			Description   string `json:"description,omitempty"`
			Designer      string `json:"designer,omitempty"`
			Duplicate     string `json:"duplicate"`
			Embeddable    string `json:"embeddable"`
			Enabled       string `json:"enabled"`
			Family        string `json:"family"`
			Fullname      string `json:"fullname"`
			Outline       string `json:"outline"`
			Style         string `json:"style"`
			Trademark     string `json:"trademark,omitempty"`
			Unique        string `json:"unique"`
			Valid         string `json:"valid"`
			Vendor        string `json:"vendor,omitempty"`
			Version       string `json:"version"`
		} `json:"typefaces"`
		Valid string `json:"valid"`
	} `json:"SPFontsDataType"`
}

type SpDisplaysType struct {
	SpDisplaysDataType []struct {
		Name                string `json:"_name"`
		Spdisplays_deviceID string `json:"spdisplays_device-id"`
		SpdisplaysNdrvs     []struct {
			Name                             string `json:"_name"`
			_spdisplays_displayProductID     string `json:"_spdisplays_display-product-id"`
			_spdisplays_displaySerialNumber2 string `json:"_spdisplays_display-serial-number2"`
			_spdisplays_displayVendorID      string `json:"_spdisplays_display-vendor-id"`
			SpdisplaysDisplayID              string `json:"_spdisplays_displayID"`
			SpdisplaysDisplayPath            string `json:"_spdisplays_displayPath"`
			SpdisplaysDisplayRegID           string `json:"_spdisplays_displayRegID"`
			SpdisplaysPixels                 string `json:"_spdisplays_pixels"`
			SpdisplaysResolution_            string `json:"_spdisplays_resolution"`
			SpdisplaysAmbientBrightness      string `json:"spdisplays_ambient_brightness"`
			SpdisplaysDepth                  string `json:"spdisplays_depth"`
			SpdisplaysMain                   string `json:"spdisplays_main"`
			SpdisplaysMirror                 string `json:"spdisplays_mirror"`
			SpdisplaysOnline                 string `json:"spdisplays_online"`
			SpdisplaysPixelresolution        string `json:"spdisplays_pixelresolution"`
			SpdisplaysResolution             string `json:"spdisplays_resolution"`
		} `json:"spdisplays_ndrvs"`
		Spdisplays_revisionID string `json:"spdisplays_revision-id"`
		Spdisplays_vendorID   string `json:"spdisplays_vendor-id"`
		SpdisplaysVram        string `json:"spdisplays_vram"`
		SppciDeviceType       string `json:"sppci_device_type"`
	} `json:"SPDisplaysDataType"`
}

type SpFrameworksType struct {
	SpFrameworksDataType []struct {
		Name             string    `json:"_name"`
		ArchKind         string    `json:"arch_kind,omitempty"`
		Info             string    `json:"info,omitempty"`
		LastModified     time.Time `json:"lastModified"`
		ObtainedFrom     string    `json:"obtained_from"`
		Path             string    `json:"path"`
		PrivateFramework string    `json:"private_framework"`
		SignedBy         []string  `json:"signed_by,omitempty"`
		Version          string    `json:"version,omitempty"`
	} `json:"SPFrameworksDataType"`
}

type SpHardwareType struct {
	SpHardwareDataType []struct {
		SmcVersionSystem      string `json:"SMC_version_system"`
		Name                  string `json:"_name"`
		AppleRomInfo          string `json:"apple_rom_info"`
		BootRomVersion        string `json:"boot_rom_version"`
		CurrentProcessorSpeed string `json:"current_processor_speed"`
		L2Cache               string `json:"l2_cache"`
		MachineModel          string `json:"machine_model"`
		MachineName           string `json:"machine_name"`
		NumberProcessors      int    `json:"number_processors"`
		Packages              int    `json:"packages"`
		PhysicalMemory        string `json:"physical_memory"`
		PlatformUuid          string `json:"platform_UUID"`
		SerialNumber          string `json:"serial_number"`
	} `json:"SPHardwareDataType"`
}

type SpInstallHistoryType struct {
	SpInstallHistoryDataType []struct {
		Name           string    `json:"_name"`
		InstallDate    time.Time `json:"install_date"`
		InstallVersion string    `json:"install_version,omitempty"`
		PackageSource  string    `json:"package_source"`
	} `json:"SPInstallHistoryDataType"`
}

type SpInternationalType struct {
	SpInternationalDataType []struct {
		Name                            string   `json:"_name"`
		BootKbd                         string   `json:"boot_kbd,omitempty"`
		BootLocale                      string   `json:"boot_locale,omitempty"`
		LinguisticDataAssetsRequested   []string `json:"linguistic_data_assets_requested,omitempty"`
		SystemCountry                   string   `json:"system_country,omitempty"`
		SystemInterfaceLanguages        []string `json:"system_interface_languages,omitempty"`
		SystemLanguages                 []string `json:"system_languages,omitempty"`
		SystemLocale                    string   `json:"system_locale,omitempty"`
		SystemTextDirection             string   `json:"system_text_direction,omitempty"`
		SystemUsesMetricSystem          string   `json:"system_uses_metric_system,omitempty"`
		UserAssistantLanguage           string   `json:"user_assistant_language,omitempty"`
		UserAssistantVoiceLanguage      string   `json:"user_assistant_voice_language,omitempty"`
		UserCalendar                    string   `json:"user_calendar,omitempty"`
		UserCountryCode                 string   `json:"user_country_code,omitempty"`
		UserCurrentInputSource          string   `json:"user_current_input_source,omitempty"`
		UserLanguageCode                string   `json:"user_language_code,omitempty"`
		UserLocale                      string   `json:"user_locale,omitempty"`
		UserPreferredInterfaceLanguages []string `json:"user_preferred_interface_languages,omitempty"`
		UserTemperatureUnit             string   `json:"user_temperature_unit,omitempty"`
		UserUsesMetricSystem            string   `json:"user_uses_metric_system,omitempty"`
	} `json:"SPInternationalDataType"`
}

type SpLegacySoftwareType struct {
	SpLegacySoftwareDataType []interface{} `json:"SPLegacySoftwareDataType"`
}

type SpNetworkLocationType struct {
	SpNetworkLocationDataType []struct {
		Name                      string `json:"_name"`
		SpnetworklocationIsActive string `json:"spnetworklocation_isActive"`
		SpnetworklocationServices []struct {
			IPv4 struct {
				ConfigMethod string `json:"ConfigMethod"`
			} `json:"IPv4"`
			IPv6 struct {
				ConfigMethod string `json:"ConfigMethod"`
			} `json:"IPv6"`
			Proxies struct {
				ExceptionsList []string `json:"ExceptionsList"`
				FtpPassive     string   `json:"FTPPassive"`
			} `json:"Proxies"`
			Name            string `json:"_name"`
			BsdDeviceName   string `json:"bsd_device_name"`
			HardwareAddress string `json:"hardware_address,omitempty"`
			Type            string `json:"type"`
		} `json:"spnetworklocation_services"`
	} `json:"SPNetworkLocationDataType"`
}

type SpLogsType struct {
	SpLogsDataType []struct {
		Name         string    `json:"_name"`
		ByteSize     int       `json:"byteSize"`
		Contents     string    `json:"contents"`
		LastModified time.Time `json:"lastModified,omitempty"`
		Source       string    `json:"source"`
	} `json:"SPLogsDataType"`
}

type SpManagedClientType struct {
	SpManagedClientDataType []interface{} `json:"SPManagedClientDataType"`
}

type SpMemoryType struct {
	SpMemoryDataType []struct {
		Items []struct {
			Name             string `json:"_name"`
			DimmManufacturer string `json:"dimm_manufacturer"`
			DimmPartNumber   string `json:"dimm_part_number"`
			DimmSerialNumber string `json:"dimm_serial_number"`
			DimmSize         string `json:"dimm_size"`
			DimmSpeed        string `json:"dimm_speed"`
			DimmStatus       string `json:"dimm_status"`
			DimmType         string `json:"dimm_type"`
		} `json:"_items"`
		Name                string `json:"_name"`
		GlobalEccState      string `json:"global_ecc_state"`
		IsMemoryUpgradeable string `json:"is_memory_upgradeable"`
	} `json:"SPMemoryDataType"`
}

type SpnvMeType struct {
	SpnvMeDataType []interface{} `json:"SPNVMeDataType"`
}

type SpNetworkType struct {
	SpNetworkDataType []struct {
		Dns *struct {
			ServerAddresses []string `json:"ServerAddresses"`
		} `json:"DNS,omitempty"`
		Ethernet *struct {
			MediaOptions []string `json:"MediaOptions"`
			MediaSubType string   `json:"MediaSubType"`
			// "MAC Address" cannot be unmarshalled into a struct field by encoding/json.
		} `json:"Ethernet,omitempty"`
		IPv4 struct {
			ArpResolvedHardwareAddress string `json:"ARPResolvedHardwareAddress,omitempty"`
			ArpResolvedIpAddress       string `json:"ARPResolvedIPAddress,omitempty"`
			AdditionalRoutes           []struct {
				DestinationAddress string `json:"DestinationAddress"`
				SubnetMask         string `json:"SubnetMask"`
			} `json:"AdditionalRoutes,omitempty"`
			Addresses              []string `json:"Addresses,omitempty"`
			ConfigMethod           string   `json:"ConfigMethod"`
			ConfirmedInterfaceName string   `json:"ConfirmedInterfaceName,omitempty"`
			InterfaceName          string   `json:"InterfaceName,omitempty"`
			NetworkSignature       string   `json:"NetworkSignature,omitempty"`
			Router                 string   `json:"Router,omitempty"`
			SubnetMasks            []string `json:"SubnetMasks,omitempty"`
		} `json:"IPv4"`
		IPv6 struct {
			ConfigMethod string `json:"ConfigMethod"`
		} `json:"IPv6"`
		Proxies struct {
			ExceptionsList []string `json:"ExceptionsList"`
			FtpPassive     string   `json:"FTPPassive"`
		} `json:"Proxies"`
		Name string `json:"_name"`
		Dhcp *struct {
			DhcpDomainNameServers string `json:"dhcp_domain_name_servers"`
			DhcpLeaseDuration     int    `json:"dhcp_lease_duration"`
			DhcpMessageType       string `json:"dhcp_message_type"`
			DhcpRouters           string `json:"dhcp_routers"`
			DhcpServerIdentifier  string `json:"dhcp_server_identifier"`
			DhcpSubnetMask        string `json:"dhcp_subnet_mask"`
		} `json:"dhcp,omitempty"`
		Hardware              string   `json:"hardware"`
		Interface             string   `json:"interface"`
		IpAddress             []string `json:"ip_address,omitempty"`
		SpnetworkServiceOrder int      `json:"spnetwork_service_order"`
		Type                  string   `json:"type"`
	} `json:"SPNetworkDataType"`
}

type SppciType struct {
	SppciDataType []interface{} `json:"SPPCIDataType"`
}

type SpParallelScsiType struct {
	SpParallelScsiDataType []interface{} `json:"SPParallelSCSIDataType"`
}

type SpPowerType struct {
	SpPowerDataType []struct {
		Name                string `json:"_name"`
		SppowerUpsInstalled string `json:"sppower_ups_installed,omitempty"`
		// "AC Power" cannot be unmarshalled into a struct field by encoding/json.
	} `json:"SPPowerDataType"`
}

type SpPrefPaneType struct {
	SpPrefPaneDataType []struct {
		Name                 string `json:"_name"`
		SpprefpaneBundlePath string `json:"spprefpane_bundlePath"`
		SpprefpaneIdentifier string `json:"spprefpane_identifier"`
		SpprefpaneIsVisible  string `json:"spprefpane_isVisible"`
		SpprefpaneKind       string `json:"spprefpane_kind"`
		SpprefpaneSupport    string `json:"spprefpane_support"`
		SpprefpaneVersion    string `json:"spprefpane_version"`
	} `json:"SPPrefPaneDataType"`
}

type SpPrintersSoftwareType struct {
	SpPrintersSoftwareDataType []struct {
		Name       string `json:"_name"`
		Extensions []struct {
			// "info path" cannot be unmarshalled into a struct field by encoding/json.
			// "info version" cannot be unmarshalled into a struct field by encoding/json.
		} `json:"extensions,omitempty"`
		LibExtensions []interface{} `json:"libExtensions"`
		Ppds          []struct {
			// "info path" cannot be unmarshalled into a struct field by encoding/json.
			// "info version" cannot be unmarshalled into a struct field by encoding/json.
		} `json:"ppds,omitempty"`
		Printers []struct {
			// "info path" cannot be unmarshalled into a struct field by encoding/json.
			// "info version" cannot be unmarshalled into a struct field by encoding/json.
		} `json:"printers,omitempty"`
		// "image capture devices" cannot be unmarshalled into a struct field by encoding/json.
		// "image capture support" cannot be unmarshalled into a struct field by encoding/json.
	} `json:"SPPrintersSoftwareDataType"`
}

type SpPrintersType struct {
	SpPrintersDataType []struct {
		Cupsversion string `json:"cupsversion"`
		Status      string `json:"status"`
	} `json:"SPPrintersDataType"`
}

type SpConfigurationProfileType struct {
	SpConfigurationProfileDataType []interface{} `json:"SPConfigurationProfileDataType"`
}

type SpRawCameraType struct {
	SpRawCameraDataType []struct {
		Name string `json:"_name"`
	} `json:"SPRawCameraDataType"`
}

type SpsasType struct {
	SpsasDataType []interface{} `json:"SPSASDataType"`
}

type SpSerialAtaType struct {
	SpSerialAtaDataType []struct {
		Items []struct {
			Name              string `json:"_name"`
			BsdName           string `json:"bsd_name,omitempty"`
			DetachableDrive   string `json:"detachable_drive"`
			DeviceModel       string `json:"device_model"`
			DeviceRevision    string `json:"device_revision"`
			DeviceSerial      string `json:"device_serial"`
			PartitionMapType  string `json:"partition_map_type,omitempty"`
			RemovableMedia    string `json:"removable_media,omitempty"`
			Size              string `json:"size,omitempty"`
			SizeInBytes       int    `json:"size_in_bytes,omitempty"`
			SpsataAsyncNotify string `json:"spsata_async_notify,omitempty"`
			SpsataMediumType  string `json:"spsata_medium_type,omitempty"`
			SpsataNcq         string `json:"spsata_ncq"`
			SpsataNcqDepth    string `json:"spsata_ncq_depth,omitempty"`
			SpsataPowerOff    string `json:"spsata_power_off,omitempty"`
			SpsataTrimSupport string `json:"spsata_trim_support,omitempty"`
			Volumes           []struct {
				Name        string `json:"_name"`
				BsdName     string `json:"bsd_name"`
				FileSystem  string `json:"file_system,omitempty"`
				Iocontent   string `json:"iocontent"`
				Size        string `json:"size"`
				SizeInBytes int    `json:"size_in_bytes"`
				VolumeUuid  string `json:"volume_uuid,omitempty"`
			} `json:"volumes,omitempty"`
		} `json:"_items"`
		Name                       string `json:"_name"`
		SpsataPhysicalInterconnect string `json:"spsata_physical_interconnect,omitempty"`
		SpsataPortdescription      string `json:"spsata_portdescription"`
		SpsataProduct              string `json:"spsata_product"`
		SpsataVendor               string `json:"spsata_vendor"`
	} `json:"SPSerialATADataType"`
}

type SpspiType struct {
	SpspiDataType []interface{} `json:"SPSPIDataType"`
}

type SpSmartCardsType struct {
	SpSmartCardsDataType []struct {
		_01   string `json:"#01,omitempty"`
		Items []struct {
			Name string `json:"_name"`
		} `json:"_items,omitempty"`
		Name string `json:"_name"`
	} `json:"SPSmartCardsDataType"`
}

type SPSoftwareType struct {
	SpSoftwareDataType []struct {
		Name            string `json:"_name"`
		BootMode        string `json:"boot_mode"`
		BootVolume      string `json:"boot_volume"`
		KernelVersion   string `json:"kernel_version"`
		LocalHostName   string `json:"local_host_name"`
		OSVersion       string `json:"os_version"`
		SecureVm        string `json:"secure_vm"`
		SystemIntegrity string `json:"system_integrity"`
		Uptime          string `json:"uptime"`
		UserName        string `json:"user_name"`
	} `json:"SPSoftwareDataType"`
}

type SpStartupItemType struct {
	SpStartupItemDataType []interface{} `json:"SPStartupItemDataType"`
}

type SpStorageType struct {
	SpStorageDataType []struct {
		Name             string `json:"_name"`
		BsdName          string `json:"bsd_name"`
		FileSystem       string `json:"file_system"`
		FreeSpaceInBytes int    `json:"free_space_in_bytes"`
		IgnoreOwnership  string `json:"ignore_ownership"`
		MountPoint       string `json:"mount_point"`
		PhysicalDrive    struct {
			DeviceName       string `json:"device_name"`
			IsInternalDisk   string `json:"is_internal_disk"`
			MediaName        string `json:"media_name"`
			MediumType       string `json:"medium_type"`
			PartitionMapType string `json:"partition_map_type"`
			Protocol         string `json:"protocol"`
		} `json:"physical_drive"`
		SizeInBytes int    `json:"size_in_bytes"`
		VolumeUuid  string `json:"volume_uuid"`
		Writable    string `json:"writable"`
	} `json:"SPStorageDataType"`
}

type SpSyncServicesType struct {
	SpSyncServicesDataType []struct {
		Items []struct {
			Name             string    `json:"_name"`
			Contents         string    `json:"contents,omitempty"`
			Description      string    `json:"description,omitempty"`
			LastModified     time.Time `json:"lastModified,omitempty"`
			Size             string    `json:"size,omitempty"`
			SummaryOfSyncLog string    `json:"summary_of_sync_log"`
		} `json:"_items"`
		Name             string `json:"_name"`
		SummaryOSVersion string `json:"summary_os_version,omitempty"`
	} `json:"SPSyncServicesDataType"`
}

type SpThunderboltType struct {
	SpThunderboltDataType []struct {
		Thunderbolt string `json:"Thunderbolt"`
	} `json:"SPThunderboltDataType"`
}

type SpusbType struct {
	SpusbDataType []struct {
		Items []struct {
			Items []struct {
				Name             string `json:"_name"`
				BcdDevice        string `json:"bcd_device"`
				BusPower         string `json:"bus_power"`
				BusPowerUsed     string `json:"bus_power_used"`
				DeviceSpeed      string `json:"device_speed"`
				ExtraCurrentUsed string `json:"extra_current_used"`
				LocationID       string `json:"location_id"`
				Manufacturer     string `json:"manufacturer"`
				ProductID        string `json:"product_id"`
				VendorID         string `json:"vendor_id"`
			} `json:"_items,omitempty"`
			Name             string `json:"_name"`
			BcdDevice        string `json:"bcd_device"`
			BusPower         string `json:"bus_power"`
			BusPowerUsed     string `json:"bus_power_used"`
			DeviceSpeed      string `json:"device_speed"`
			ExtraCurrentUsed string `json:"extra_current_used"`
			LocationID       string `json:"location_id"`
			Manufacturer     string `json:"manufacturer"`
			ProductID        string `json:"product_id"`
			SerialNum        string `json:"serial_num,omitempty"`
			VendorID         string `json:"vendor_id"`
		} `json:"_items,omitempty"`
		Name           string `json:"_name"`
		HostController string `json:"host_controller"`
		PciDevice      string `json:"pci_device"`
		PciRevision    string `json:"pci_revision"`
		PciVendor      string `json:"pci_vendor"`
	} `json:"SPUSBDataType"`
}

type SpNetworkVolumeType struct {
	SpNetworkVolumeDataType []struct {
		Name                       string `json:"_name"`
		SpnetworkvolumeAutomounted string `json:"spnetworkvolume_automounted"`
		SpnetworkvolumeFsmtnonname string `json:"spnetworkvolume_fsmtnonname"`
		SpnetworkvolumeFstypename  string `json:"spnetworkvolume_fstypename"`
		SpnetworkvolumeMntfromname string `json:"spnetworkvolume_mntfromname"`
	} `json:"SPNetworkVolumeDataType"`
}

type SpwwanType struct {
	SpwwanDataType []interface{} `json:"SPWWANDataType"`
}

type SpAirPortType struct {
	SpAirPortDataType []struct {
		SpairportSoftwareInformation struct {
			SpairportCorewlanVersion    string `json:"spairport_corewlan_version"`
			SpairportCorewlankitVersion string `json:"spairport_corewlankit_version"`
			SpairportDiagnosticsVersion string `json:"spairport_diagnostics_version"`
			SpairportExtraVersion       string `json:"spairport_extra_version"`
			SpairportFamilyVersion      string `json:"spairport_family_version"`
			SpairportProfilerVersion    string `json:"spairport_profiler_version"`
			SpairportUtilityVersion     string `json:"spairport_utility_version"`
		} `json:"spairport_software_information"`
	} `json:"SPAirPortDataType"`
}

type DarwinSystemProfilerInfo struct {
	SPParallelATA          SpParallelATAType
	SPUniversalAccess      SpUniversalAccessType
	SPSecureElement        SpSecureElementType
	SPApplications         SpApplicationsType
	SPAudio                SpAudioType
	SPBluetooth            SpBluetoothType
	SPCamera               SpCameraType
	SPCardReader           SpCardReaderType
	SPiBridge              SpiBridgeType
	SPDeveloperTools       SpDeveloperToolsType
	SPDiagnostics          SpDiagnosticsType
	SPDisabledSoftware     SpDisabledSoftwareType
	SPDiscBurning          SpDiscBurningType
	SPEthernet             SpEthernetType
	SPExtensions           SpExtensionsType
	SPFibreChannel         SpFibreChannelType
	SPFireWire             SpFireWireType
	SPFirewall             SpFirewallType
	SPFonts                SpFontsType
	SPFrameworks           SpFrameworksType
	SPDisplays             SpDisplaysType
	SPHardware             SpHardwareType
	SPInstallHistory       SpInstallHistoryType
	SPInternational        SpInternationalType
	SPLegacySoftware       SpLegacySoftwareType
	SPNetworkLocation      SpNetworkLocationType
	SPLogs                 SpLogsType
	SPManagedClient        SpManagedClientType
	SPMemory               SpMemoryType
	SPNVMe                 SpnvMeType
	SPNetwork              SpNetworkType
	SPPCI                  SppciType
	SPParallelSCSI         SpParallelScsiType
	SPPower                SpPowerType
	SPPrefPane             SpPrefPaneType
	SPPrintersSoftware     SpPrintersSoftwareType
	SPPrinters             SpPrintersType
	SPConfigurationProfile SpConfigurationProfileType
	SPRawCameraDataType    SpRawCameraType
	SPSAS                  SpsasType
	SPSerialATA            SpSerialAtaType
	SPSPI                  SpspiType
	SPSmartCards           SpSmartCardsType
	SPSoftware             SPSoftwareType
	SPStartupItemType      SpStartupItemType
	SPStorage              SpStorageType
	SPSyncServicesType     SpSyncServicesType
	SPThunderbolt          SpThunderboltType
	SPUSB                  SpusbType
	SPNetworkVolume        SpNetworkVolumeType
	SPWWAN                 SpwwanType
	SPAirPort              SpAirPortType
}
