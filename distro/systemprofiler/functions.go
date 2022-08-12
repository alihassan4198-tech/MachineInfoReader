//go:build darwin
// +build darwin

package systemprofiler

import (
	"encoding/json"
	"machine_info_gatherer/common"
	"sync"
)

func DarwinGetSystemProfilerInfo[T SpParallelATAType |
	SpUniversalAccessType | SpSecureElementType | SpApplicationsType | SpAudioType | SpBluetoothType | SpCameraType | SpCardReaderType | SpiBridgeType | SpDeveloperToolsType | SpDiagnosticsType | SpDisabledSoftwareType | SpDiscBurningType | SpEthernetType | SpExtensionsType | SpFibreChannelType | SpFireWireType | SpFirewallType | SpFontsType | SpDisplaysType | SpFrameworksType | SpHardwareType | SpInstallHistoryType | SpInternationalType | SpLegacySoftwareType | SpNetworkLocationType | SpLogsType | SpManagedClientType | SpMemoryType | SpnvMeType | SpNetworkType | SppciType | SpParallelScsiType | SpPowerType | SpPrefPaneType | SpPrintersSoftwareType | SpPrintersType | SpConfigurationProfileType | SpRawCameraType | SpsasType | SpSerialAtaType | SpspiType | SpSmartCardsType | SPSoftwareType | SpStartupItemType | SpStorageType | SpSyncServicesType | SpThunderboltType | SpusbType | SpNetworkVolumeType | SpwwanType | SpAirPortType](
	command string) T {
	var jsonObj T

	systemProfiler, err := common.RunFullCommand("system_profiler " + command + " -json")
	if err == nil {
		err = json.Unmarshal([]byte(systemProfiler), &jsonObj)
		if err != nil {
			// fmt.Println(err)
		} else {
			return jsonObj
		}
	}
	return jsonObj
}

func DarwinSystemProfiler() DarwinSystemProfilerInfo {

	var wg sync.WaitGroup

	var darwinSPInfo DarwinSystemProfilerInfo = DarwinSystemProfilerInfo{}

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPNVMe = DarwinGetSystemProfilerInfo[SpnvMeType]("SPNVMeDataType")
	// 	dspi.SPNetwork = DarwinGetSystemProfilerInfo[SpNetworkType]("SPNetworkDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPUSB = DarwinGetSystemProfilerInfo[SpusbType]("SPUSBDataType")
	// 	dspi.SPNetworkVolume = DarwinGetSystemProfilerInfo[SpNetworkVolumeType]("SPNetworkVolumeDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPSyncServicesType = DarwinGetSystemProfilerInfo[SpSyncServicesType]("SPSyncServicesDataType")
	// 	dspi.SPThunderbolt = DarwinGetSystemProfilerInfo[SpThunderboltType]("SPThunderboltDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPStartupItemType = DarwinGetSystemProfilerInfo[SpStartupItemType]("SPStartupItemDataType")
	// 	dspi.SPStorage = DarwinGetSystemProfilerInfo[SpStorageType]("SPStorageDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPPrintersSoftware = DarwinGetSystemProfilerInfo[SpPrintersSoftwareType]("SPPrintersSoftwareDataType")
	// 	dspi.SPSoftware = DarwinGetSystemProfilerInfo[SPSoftwareType]("SPSoftwareDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPPower = DarwinGetSystemProfilerInfo[SpPowerType]("SPPowerDataType")
	// 	dspi.SPPrefPane = DarwinGetSystemProfilerInfo[SpPrefPaneType]("SPPrefPaneDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPManagedClient = DarwinGetSystemProfilerInfo[SpManagedClientType]("SPManagedClientDataType")
	// 	dspi.SPMemory = DarwinGetSystemProfilerInfo[SpMemoryType]("SPMemoryDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPInstallHistory = DarwinGetSystemProfilerInfo[SpInstallHistoryType]("SPInstallHistoryDataType")
	// 	dspi.SPInternational = DarwinGetSystemProfilerInfo[SpInternationalType]("SPInternationalDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPHardware = DarwinGetSystemProfilerInfo[SpHardwareType]("SPHardwareDataType")
	// 	dspi.SPSerialATA = DarwinGetSystemProfilerInfo[SpSerialAtaType]("SPSerialATADataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPDiagnostics = DarwinGetSystemProfilerInfo[SpDiagnosticsType]("SPDiagnosticsDataType")
	// 	dspi.SPDisabledSoftware = DarwinGetSystemProfilerInfo[SpDisabledSoftwareType]("SPDisabledSoftwareDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPSAS = DarwinGetSystemProfilerInfo[SpsasType]("SPSASDataType")
	// 	dspi.SPCardReader = DarwinGetSystemProfilerInfo[SpCardReaderType]("SPCardReaderDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPCamera = DarwinGetSystemProfilerInfo[SpCameraType]("SPCameraDataType")
	// 	dspi.SPUniversalAccess = DarwinGetSystemProfilerInfo[SpUniversalAccessType]("SPUniversalAccessDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPAudio = DarwinGetSystemProfilerInfo[SpAudioType]("SPAudioDataType")
	// 	dspi.SPParallelATA = DarwinGetSystemProfilerInfo[SpParallelATAType]("SPParallelATADataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPSecureElement = DarwinGetSystemProfilerInfo[SpSecureElementType]("SPSecureElementDataType")
	// 	dspi.SPBluetooth = DarwinGetSystemProfilerInfo[SpBluetoothType]("SPBluetoothDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPRawCameraDataType = DarwinGetSystemProfilerInfo[SpRawCameraType]("SPRawCameraDataType")
	// 	dspi.SPApplications = DarwinGetSystemProfilerInfo[SpApplicationsType]("SPApplicationsDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPiBridge = DarwinGetSystemProfilerInfo[SpiBridgeType]("SPiBridgeDataType")
	// 	dspi.SPDeveloperTools = DarwinGetSystemProfilerInfo[SpDeveloperToolsType]("SPDeveloperToolsDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPDiscBurning = DarwinGetSystemProfilerInfo[SpDiscBurningType]("SPDiscBurningDataType")
	// 	dspi.SPEthernet = DarwinGetSystemProfilerInfo[SpEthernetType]("SPEthernetDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPExtensions = DarwinGetSystemProfilerInfo[SpExtensionsType]("SPExtensionsDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPFibreChannel = DarwinGetSystemProfilerInfo[SpFibreChannelType]("SPFibreChannelDataType")
	// 	dspi.SPFireWire = DarwinGetSystemProfilerInfo[SpFireWireType]("SPFireWireDataType")
	// 	dspi.SPFirewall = DarwinGetSystemProfilerInfo[SpFirewallType]("SPFirewallDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPFonts = DarwinGetSystemProfilerInfo[SpFontsType]("SPFontsDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPFrameworks = DarwinGetSystemProfilerInfo[SpFrameworksType]("SPFrameworksDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPDisplays = DarwinGetSystemProfilerInfo[SpDisplaysType]("SPDisplaysDataType")
	// 	dspi.SPLegacySoftware = DarwinGetSystemProfilerInfo[SpLegacySoftwareType]("SPLegacySoftwareDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPNetworkLocation = DarwinGetSystemProfilerInfo[SpNetworkLocationType]("SPNetworkLocationDataType")
	// 	dspi.SPLogs = DarwinGetSystemProfilerInfo[SpLogsType]("SPLogsDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPPCI = DarwinGetSystemProfilerInfo[SppciType]("SPPCIDataType")
	// 	dspi.SPParallelSCSI = DarwinGetSystemProfilerInfo[SpParallelScsiType]("SPParallelSCSIDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPPrinters = DarwinGetSystemProfilerInfo[SpPrintersType]("SPPrintersDataType")
	// 	dspi.SPConfigurationProfile = DarwinGetSystemProfilerInfo[SpConfigurationProfileType]("SPConfigurationProfileDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPSPI = DarwinGetSystemProfilerInfo[SpspiType]("SPSPIDataType")
	// 	dspi.SPSmartCards = DarwinGetSystemProfilerInfo[SpSmartCardsType]("SPSmartCardsDataType")

	// }(&darwinSPInfo)

	// wg.Add(1)
	// go func(dspi *DarwinSystemProfilerInfo) {
	// 	defer wg.Done()
	// 	dspi.SPWWAN = DarwinGetSystemProfilerInfo[SpwwanType]("SPWWANDataType")
	// 	dspi.SPAirPort = DarwinGetSystemProfilerInfo[SpAirPortType]("SPAirPortDataType")
	// }(&darwinSPInfo)

	wg.Wait()
	return darwinSPInfo
}
