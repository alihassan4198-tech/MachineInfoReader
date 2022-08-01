//go:build darwin
// +build darwin

package distro

import (
	"encoding/json"
	"fmt"
	"machine_info_gatherer/common"
)

func (dspi *DarwinSystemProfilerInfo) darwinSPParallelATADataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPParallelATADataType -json")
	if err == nil {
		var jsonObj SpParallelATAType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPParallelATA = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPUniversalAccessDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPUniversalAccessDataType -json")
	if err == nil {
		var jsonObj SpUniversalAccessType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPUniversalAccess = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPSecureElementDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPSecureElementDataType -json")
	if err == nil {
		var jsonObj SpSecureElementType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPSecureElement = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPApplicationsDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPApplicationsDataType -json")
	if err == nil {
		var jsonObj SpApplicationsType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPApplications = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPAudioDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPAudioDataType -json")
	if err == nil {
		var jsonObj SpAudioType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPAudio = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPBluetoothDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPBluetoothDataType -json")
	if err == nil {
		var jsonObj SpBluetoothType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPUniversalAccess = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPCameraDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPCameraDataType -json")
	if err == nil {
		var jsonObj SpCameraType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPUniversalAccess = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPCardReaderDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPCardReaderDataType -json")
	if err == nil {
		var jsonObj SpCardReaderType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPUniversalAccess = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPiBridgeDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPiBridgeDataType -json")
	if err == nil {
		var jsonObj SpiBridgeType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPDeveloperToolsDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPDeveloperToolsDataType -json")
	if err == nil {
		var jsonObj SpDeveloperToolsType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPUniversalAccess = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPDiagnosticsDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPDiagnosticsDataType -json")
	if err == nil {
		var jsonObj SpDiagnosticsType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPUniversalAccess = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPDisabledSoftwareDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPDisabledSoftwareDataType -json")
	if err == nil {
		var jsonObj SpDisabledSoftwareType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPUniversalAccess = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPDiscBurningDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPDiscBurningDataType -json")
	if err == nil {
		var jsonObj SpDiscBurningType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPUniversalAccess = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPEthernetDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPEthernetDataType -json")
	if err == nil {
		var jsonObj SpEthernetType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPUniversalAccess = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPExtensionsDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPExtensionsDataType -json")
	if err == nil {
		var jsonObj SpExtensionsType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPUniversalAccess = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPFibreChannelDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPFibreChannelDataType -json")
	if err == nil {
		var jsonObj SpFibreChannelType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPUniversalAccess = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPFireWireDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPFireWireDataType -json")
	if err == nil {
		var jsonObj SpFireWireType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPUniversalAccess = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPFirewallDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPFirewallDataType -json")
	if err == nil {
		var jsonObj SpFirewallType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPUniversalAccess = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPFontsDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPFontsDataType -json")
	if err == nil {
		var jsonObj SpFontsType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPUniversalAccess = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPFrameworksDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPFrameworksDataType -json")
	if err == nil {
		var jsonObj SpFrameworksType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		} else {
			dspi.SPUniversalAccess = jsonObj
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPDisplaysDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPDisplaysDataType -json")
	if err == nil {
		var jsonObj SpDisplaysType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPHardwareDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPHardwareDataType -json")
	if err == nil {
		var jsonObj SpHardwareType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPInstallHistoryDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPInstallHistoryDataType -json")
	if err == nil {
		var jsonObj SpInstallHistoryType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPInternationalDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPInternationalDataType -json")
	if err == nil {
		var jsonObj SpInternationalType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPLegacySoftwareDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPLegacySoftwareDataType -json")
	if err == nil {
		var jsonObj SpLegacySoftwareType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPNetworkLocationDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPNetworkLocationDataType -json")
	if err == nil {
		var jsonObj SpNetworkLocationType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPLogsDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPLogsDataType -json")
	if err == nil {
		var jsonObj SpLogsType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPManagedClientDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPManagedClientDataType -json")
	if err == nil {
		var jsonObj SpManagedClientType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPMemoryDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPMemoryDataType -json")
	if err == nil {
		var jsonObj SpMemoryType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPNVMeDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPNVMeDataType -json")
	if err == nil {
		var jsonObj SpnvMeType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPNetworkDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPNetworkDataType -json")
	if err == nil {
		var jsonObj SpNetworkType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPPCIDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPPCIDataType -json")
	if err == nil {
		var jsonObj SppciType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPParallelSCSIDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPParallelSCSIDataType -json")
	if err == nil {
		var jsonObj SpParallelScsiType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPPowerDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPPowerDataType -json")
	if err == nil {
		var jsonObj SpPowerType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPPrefPaneDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPPrefPaneDataType -json")
	if err == nil {
		var jsonObj SpPrefPaneType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPPrintersSoftwareDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPPrintersSoftwareDataType -json")
	if err == nil {
		var jsonObj SpPrintersSoftwareType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPPrintersDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPPrintersDataType -json")
	if err == nil {
		var jsonObj SpPrintersType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPConfigurationProfileDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPConfigurationProfileDataType -json")
	if err == nil {
		var jsonObj SpConfigurationProfileType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPRawCameraDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPRawCameraDataType -json")
	if err == nil {
		var jsonObj SpRawCameraType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPSASDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPSASDataType -json")
	if err == nil {
		var jsonObj SpsasType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPSerialATADataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPSerialATADataType -json")
	if err == nil {
		var jsonObj SpSerialAtaType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPSPIDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPSPIDataType -json")
	if err == nil {
		var jsonObj SpspiType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPSmartCardsDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPSmartCardsDataType -json")
	if err == nil {
		var jsonObj SpSmartCardsType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPSoftwareDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPSoftwareDataType -json")
	if err == nil {
		var jsonObj SPSoftwareType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPStartupItemDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPStartupItemDataType -json")
	if err == nil {
		var jsonObj SpStartupItemType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPStorageDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPStorageDataType -json")
	if err == nil {
		var jsonObj SpStorageType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPSyncServicesDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPSyncServicesDataType -json")
	if err == nil {
		var jsonObj SpSyncServicesType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPThunderboltDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPThunderboltDataType -json")
	if err == nil {
		var jsonObj SpThunderboltType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPUSBDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPUSBDataType -json")
	if err == nil {
		var jsonObj SpusbType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPNetworkVolumeDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPNetworkVolumeDataType -json")
	if err == nil {
		var jsonObj SpNetworkVolumeType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPWWANDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPWWANDataType -json")
	if err == nil {
		var jsonObj SpwwanType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func (dspi *DarwinSystemProfilerInfo) darwinSPAirPortDataType() {

	systemProfiler, err := common.RunFullCommand("system_profiler SPAirPortDataType -json")
	if err == nil {
		var jsonObj SpAirPortType
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func DarwinSystemProfiler() DarwinSystemProfilerInfo {

	var darwinSPInfo DarwinSystemProfilerInfo = DarwinSystemProfilerInfo{}

	darwinSPInfo.darwinSPParallelATADataType()
	darwinSPInfo.darwinSPUniversalAccessDataType()
	darwinSPInfo.darwinSPSecureElementDataType()
	darwinSPInfo.darwinSPApplicationsDataType()
	darwinSPInfo.darwinSPAudioDataType()
	darwinSPInfo.darwinSPBluetoothDataType()
	darwinSPInfo.darwinSPCameraDataType()
	darwinSPInfo.darwinSPCardReaderDataType()
	darwinSPInfo.darwinSPiBridgeDataType()
	darwinSPInfo.darwinSPDeveloperToolsDataType()
	darwinSPInfo.darwinSPDiagnosticsDataType()
	darwinSPInfo.darwinSPDisabledSoftwareDataType()
	darwinSPInfo.darwinSPDiscBurningDataType()
	darwinSPInfo.darwinSPEthernetDataType()
	darwinSPInfo.darwinSPExtensionsDataType()
	darwinSPInfo.darwinSPFibreChannelDataType()
	darwinSPInfo.darwinSPFireWireDataType()
	darwinSPInfo.darwinSPFirewallDataType()
	darwinSPInfo.darwinSPFontsDataType()
	darwinSPInfo.darwinSPFrameworksDataType()
	darwinSPInfo.darwinSPDisplaysDataType()
	darwinSPInfo.darwinSPHardwareDataType()
	darwinSPInfo.darwinSPInstallHistoryDataType()
	darwinSPInfo.darwinSPInternationalDataType()
	darwinSPInfo.darwinSPLegacySoftwareDataType()
	darwinSPInfo.darwinSPNetworkLocationDataType()
	darwinSPInfo.darwinSPLogsDataType()
	darwinSPInfo.darwinSPManagedClientDataType()
	darwinSPInfo.darwinSPMemoryDataType()
	darwinSPInfo.darwinSPNVMeDataType()
	darwinSPInfo.darwinSPNetworkDataType()
	darwinSPInfo.darwinSPPCIDataType()
	darwinSPInfo.darwinSPParallelSCSIDataType()
	darwinSPInfo.darwinSPPowerDataType()
	darwinSPInfo.darwinSPPrefPaneDataType()
	darwinSPInfo.darwinSPPrintersSoftwareDataType()
	darwinSPInfo.darwinSPPrintersDataType()
	darwinSPInfo.darwinSPConfigurationProfileDataType()
	darwinSPInfo.darwinSPRawCameraDataType()
	darwinSPInfo.darwinSPSASDataType()
	darwinSPInfo.darwinSPSerialATADataType()
	darwinSPInfo.darwinSPSPIDataType()
	darwinSPInfo.darwinSPSmartCardsDataType()
	darwinSPInfo.darwinSPSoftwareDataType()
	darwinSPInfo.darwinSPStartupItemDataType()
	darwinSPInfo.darwinSPStorageDataType()
	darwinSPInfo.darwinSPSyncServicesDataType()
	darwinSPInfo.darwinSPThunderboltDataType()
	darwinSPInfo.darwinSPUSBDataType()
	darwinSPInfo.darwinSPNetworkVolumeDataType()
	darwinSPInfo.darwinSPWWANDataType()
	darwinSPInfo.darwinSPAirPortDataType()

	return darwinSPInfo
}
