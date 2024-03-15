package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/pothosware/go-soapy-sdr/pkg/device"
	"github.com/pothosware/go-soapy-sdr/pkg/sdrlogger"
)

func main() {
	// Test log levels
	sdrlogger.RegisterLogHandler(logSoapy)
	sdrlogger.SetLogLevel(sdrlogger.SSI)
	sdrlogger.Log(sdrlogger.Info, "Soapy SDR")
	sdrlogger.Logf(sdrlogger.Info, "%v\n", "Demonstration")

	sdrlogger.Log(sdrlogger.Fatal, "Testing Fatal logging level")
	sdrlogger.Log(sdrlogger.Critical, "Testing Critical logging level")
	sdrlogger.Log(sdrlogger.Error, "Testing Error logging level")
	sdrlogger.Log(sdrlogger.Warning, "Testing Warning logging level")
	sdrlogger.Log(sdrlogger.Notice, "Testing Notice logging level")
	sdrlogger.Log(sdrlogger.Info, "Testing Info logging level")
	sdrlogger.Log(sdrlogger.Debug, "Testing Debug logging level")
	sdrlogger.Log(sdrlogger.Trace, "Testing Trace logging level")
	sdrlogger.Log(sdrlogger.SSI, "Testing SSI logger level")

	// List all devices
	devices := device.Enumerate(nil)
	for i, dev := range devices {
		fmt.Printf("Found device #%v:\n", i)
		for k, v := range dev {
			fmt.Printf("%v=%v\n", k, v)
		}
		fmt.Printf("\n")
	}
	if len(devices) == 0 {
		fmt.Printf("No devices found!!\n")
		return
	}

	// Convert device info arguments for opening all detected devices
	deviceArgs := make([]map[string]string, len(devices))

	for i, dev := range devices {
		deviceArgs[i] = map[string]string{
			"driver": dev["driver"],
		}
	}

	// Open all devices
	devs, err := device.MakeList(deviceArgs)
	if err != nil {
		log.Panic(err)
	}
	defer func([]*device.SDRDevice) {
		// Close all devices
		err := device.UnmakeList(devs)
		if err != nil {
			log.Panic(err)
		}
		sdrlogger.Log(sdrlogger.Trace, "All devices closed")
	}(devs)

	for i, dev := range devs {
		fmt.Printf("***************\n")
		fmt.Printf("Device: %v\n", devices[i]["label"])
		fmt.Printf("***************\n")

		displayDetails(dev)
		receiveSomeData(dev)
	}

	// Close all devices
	/*	err = device.UnmakeList(devs)
		if err != nil {
			log.Panic(err)
		}*/

	//sdrlogger.Log(sdrlogger.Trace, "All devices closed")
}

// displayDetails displays the details and information for a device (for all its directions and channels)
func displayDetails(dev *device.SDRDevice) {
	fmt.Printf("***************\n")
	fmt.Printf("Device Information\n")
	fmt.Printf("***************\n")

	// Print hardware info for the device
	displayHardwareInfo(dev)

	// GPIO
	displayGPIOBanks(dev)

	// Settings
	displaySettingInfo(dev)

	// UARTs
	displayUARTs(dev)

	// Clocking
	displayMasterClockRate(dev)
	displayClockSources(dev)

	// Registers
	displayRegisters(dev)

	// Device Sensor
	displaySensors(dev)

	// Time Sources
	displayTimeSources(dev)

	// Direction details
	displayDirectionDetails(dev, device.DirectionTX)
	displayDirectionDetails(dev, device.DirectionRX)
}

// displayHardwareInfo prints hardware info for the specified device
func displayHardwareInfo(dev *device.SDRDevice) {
	fmt.Printf("DriverKey: %v\n", dev.GetDriverKey())
	fmt.Printf("HardwareKey: %v\n", dev.GetHardwareKey())
	hardwareInfo := dev.GetHardwareInfo()
	if len(hardwareInfo) > 0 {
		for k, v := range hardwareInfo {
			fmt.Printf("HardwareInfo: %v: %v\n", k, v)
		}
	} else {
		fmt.Println("HardwareInfo: [none]")
	}
}

// displayGPIOBanks prints GPIO bank info for the specified device
func displayGPIOBanks(dev *device.SDRDevice) {
	banks := dev.ListGPIOBanks()
	if len(banks) > 0 {
		for i, bank := range banks {
			fmt.Printf("GPIO Bank#%d: %v\n", i, bank)
		}
	} else {
		fmt.Println("GPIO Banks: [none]")
	}
}

// displaySettingInfo prints a device's setting information
func displaySettingInfo(dev *device.SDRDevice) {
	settings := dev.GetSettingInfo()
	if len(settings) > 0 {
		for i, setting := range settings {
			fmt.Printf("Setting#%d:\n", i)
			displaySettingValues(setting)
		}
	} else {
		fmt.Println("Settings: [none]")
	}
}

// displaySettingValues prints each setting value
func displaySettingValues(setting device.SDRArgInfo) {
	fmt.Printf("  key: %v\n", setting.Key)
	fmt.Printf("  value: %v\n", setting.Value)
	fmt.Printf("  name: %v\n", setting.Name)
	fmt.Printf("  description: %v\n", setting.Description)
	fmt.Printf("  unit: %v\n", setting.Unit)
	var argType string = "unknown type"
	switch setting.Type {
	case device.ArgInfoBool:
		argType = "bool"
	case device.ArgInfoInt:
		argType = "integer"
	case device.ArgInfoFloat:
		argType = "float"
	case device.ArgInfoString:
		argType = "string"
	}
	fmt.Printf("  type: %v\n", argType)
	fmt.Printf("  range: %v\n", setting.Range.ToString())
	numOptions := setting.NumOptions
	if numOptions > 0 {
		fmt.Printf("  options: %v\n", setting.Options)
		fmt.Printf("  option names: %v\n", setting.OptionNames)
	} else {
		fmt.Println("  options: [none]")
		fmt.Println("  option names: [none]")
	}
}

// displayUARTs prints a devices's UARTs
func displayUARTs(dev *device.SDRDevice) {
	uarts := dev.ListUARTs()
	if len(uarts) > 0 {
		for i, uart := range uarts {
			fmt.Printf("UARTs#%d:%v\n", i, uart)
		}
	} else {
		fmt.Println("UARTs: [none]")
	}
}

// displayMasterClockRate prints a device's master clock rate and clock ranges
func displayMasterClockRate(dev *device.SDRDevice) {
	fmt.Printf("Master Clock Rate: %v\n", dev.GetMasterClockRate())
	clockRanges := dev.GetMasterClockRates()
	if len(clockRanges) > 0 {
		fmt.Println("Master Clock Rate Ranges:")
		for i, clockRange := range clockRanges {
			fmt.Printf("  Range#%d: %v\n", i, clockRange)
		}
	} else {
		fmt.Println("Clock Rate Ranges: [none]")
	}
}

// displayClockSources prints a device's clock sources
func displayClockSources(dev *device.SDRDevice) {
	clockSources := dev.ListClockSources()
	if len(clockSources) > 0 {
		fmt.Println("Clock Sources:")
		for i, clockSource := range clockSources {
			fmt.Printf("  Clock Source#%d: %v\n", i, clockSource)
		}
	} else {
		fmt.Println("Clock Sources: [none]")
	}
}

// displayRegisters prints a device's registers
func displayRegisters(dev *device.SDRDevice) {
	registers := dev.ListRegisterInterfaces()
	if len(registers) > 0 {
		fmt.Println("Registers:")
		for i, register := range registers {
			fmt.Printf("  Register#%d: %v\n", i, register)
		}
	} else {
		fmt.Println("Registers: [none]")
	}
}

// displaySensors prints a device's sensors
func displaySensors(dev *device.SDRDevice) {
	sensors := dev.ListSensors()
	if len(sensors) > 0 {
		fmt.Println("Sensors:")
		for i, sensor := range sensors {
			fmt.Printf("  Sensor#%d: %v/n", i, sensor)
		}
	} else {
		fmt.Println("Sensors: [none]")
	}
}

// displayTimeSources lists all of a device's time sources and hardware time if any
func displayTimeSources(dev *device.SDRDevice) {
	timeSources := dev.ListTimeSources()
	if len(timeSources) > 0 {
		fmt.Println("Time Sources:")
		for i, timeSource := range timeSources {
			fmt.Printf("  Time Source#%d: %v\n", i, timeSource)
		}
	} else {
		fmt.Println("Time Sources: [none]")
	}

	hasHardwareTime := dev.HasHardwareTime("")
	fmt.Printf("Has Hardware Time: %v\n", hasHardwareTime)
	if hasHardwareTime {
		fmt.Printf("  Hardware Time: %v\n", dev.GetHardwareTime(""))
	}
}

// displayDirectionDetails prints info about TX and RX channels
func displayDirectionDetails(dev *device.SDRDevice, direction device.Direction) {
	if direction == device.DirectionTX {
		fmt.Println("Direction TX")
	} else {
		fmt.Println("Direction RX")
	}

	frontEndMapping := dev.GetFrontendMapping(direction)
	if len(frontEndMapping) > 0 {
		fmt.Printf("  FrontendMapping: %v\n", frontEndMapping)
	} else {
		fmt.Println("  FrontendMapping: [none]")
	}

	numChannels := dev.GetNumChannels(direction)
	fmt.Printf("  Number of Channels: %v\n", numChannels)

	for channel := uint(0); channel < numChannels; channel++ {
		displayDirectionChannelDetails(dev, direction, channel)
	}
}

// displayDirectionChannelDetails prints out details and info of a device / direction / channel
func displayDirectionChannelDetails(dev *device.SDRDevice, direction device.Direction, channel uint) {
	// Settings
	settings := dev.GetChannelSettingInfo(direction, channel)
	if len(settings) > 0 {
		for i, setting := range settings {
			fmt.Printf("Channel#%d Setting#%d Banks: %v\n", channel, i, setting)
		}
	} else {
		fmt.Printf("Channel#%d Settings: [none]\n", channel)
	}

	// Channel
	channelInfo := dev.GetChannelInfo(direction, channel)
	if len(channelInfo) > 0 {
		for k, v := range channelInfo {
			fmt.Printf("Channel#%d ChannelInfo: {%v: %v}\n", channel, k, v)
		}
	} else {
		fmt.Printf("Channel#%d ChannelInfo: [none]\n", channel)
	}

	fmt.Printf("Channel#%d Fullduplex: %v\n", channel, dev.GetFullDuplex(direction, channel))

	// Antenna
	antennas := dev.ListAntennas(direction, channel)
	fmt.Printf("Channel#%d NumAntennas: %v\n", channel, len(antennas))

	for i, antenna := range antennas {
		fmt.Printf("Channel#%d Antenna#%d: %v\n", channel, i, antenna)
	}

	// Bandwidth
	fmt.Printf("Channel#%d Baseband filter width: %v Hz\n", channel, dev.GetBandwidth(direction, channel))

	bandwidthRanges := dev.GetBandwidthRanges(direction, channel)
	for i, bandwidthRange := range bandwidthRanges {
		fmt.Printf("Channel#%d Baseband filter#%d: %v\n", channel, i, bandwidthRange)
	}

	// Gain
	fmt.Printf("Channel#%d HasGainMode (Automatic gain possible): %v\n", channel, dev.HasGainMode(direction, channel))
	fmt.Printf("Channel#%d GainMode (Automatic gain enabled): %v\n", channel, dev.GetGainMode(direction, channel))
	fmt.Printf("Channel#%d Gain: %v\n", channel, dev.GetGain(direction, channel))
	fmt.Printf("Channel#%d GainRange: %v\n", channel, dev.GetGainRange(direction, channel))
	gainElements := dev.ListGains(direction, channel)
	fmt.Printf("Channel#%d NumGainElements: %v\n", channel, len(gainElements))

	for i, gainElement := range gainElements {
		fmt.Printf("Channel#%d Gain Element#%d Name: %v\n", channel, i, gainElement)
		fmt.Printf("Channel#%d Gain Element#%d Value: %v\n", channel, i, dev.GetGainElement(direction, channel, gainElement))
		fmt.Printf("Channel#%d Gain Element#%d Range: %v\n", channel, i, dev.GetGainElementRange(direction, channel, gainElement).ToString())
	}

	// SampleRate
	fmt.Printf("Channel#%d Sample Rate: %v\n", channel, dev.GetSampleRate(direction, channel))
	sampleRateRanges := dev.GetSampleRateRange(direction, channel)
	for i, sampleRateRange := range sampleRateRanges {
		fmt.Printf("Channel#%d Sample Rate Range#%d: %v\n", channel, i, sampleRateRange.ToString())
	}

	// Frequencies
	fmt.Printf("Channel#%d Frequency: %v\n", channel, dev.GetFrequency(direction, channel))
	frequencyRanges := dev.GetFrequencyRange(direction, channel)
	for i, frequencyRange := range frequencyRanges {
		fmt.Printf("Channel#%d Frequency Range#%d: %v\n", channel, i, frequencyRange.ToString())
	}

	frequencyArgsInfos := dev.GetFrequencyArgsInfo(direction, channel)

	if len(frequencyArgsInfos) > 0 {
		for i, argInfo := range frequencyArgsInfos {
			fmt.Printf("Channel#%d Frequency ArgInfo#%d: %v\n", channel, i, argInfo.ToString())
		}
	} else {
		fmt.Printf("Channel#%d Frequency ArgInfo: [none]\n", channel)
	}

	frequencyComponents := dev.ListFrequencies(direction, channel)
	fmt.Printf("Channel#%d NumFrequencyComponents: %v\n", channel, len(frequencyComponents))

	for i, frequencyComponent := range frequencyComponents {
		fmt.Printf("Channel#%d Frequency Component#%d Name: %v\n", channel, i, frequencyComponent)
		fmt.Printf("Channel#%d Frequency Component#%d Frequency: %v\n", channel, i, dev.GetFrequencyComponent(direction, channel, frequencyComponent))
	}

	// Stream
	streamFormats := dev.GetStreamFormats(direction, channel)
	fmt.Printf("Channel#%d Stream Formats: %v\n", channel, streamFormats)
	nativeStreamFormat, nativeStreamFullScale := dev.GetNativeStreamFormat(direction, channel)
	fmt.Printf("Channel#%d Stream Native Format: %v\n", channel, nativeStreamFormat)
	fmt.Printf("Channel#%d Stream Native FullScale: %v\n", channel, nativeStreamFullScale)

	streamArgsInfos := dev.GetStreamArgsInfo(direction, channel)
	if len(streamArgsInfos) > 0 {
		for i, argInfo := range streamArgsInfos {
			fmt.Printf("Channel#%d Stream ArgInfo#%d: %v\n", channel, i, argInfo.ToString())
		}
	} else {
		fmt.Printf("Channel#%d Stream ArgInfo: [none]\n", channel)
	}

	// Frontend correctiion
	available := dev.HasDCOffsetMode(direction, channel)
	fmt.Printf("Channel#%d Stream Correction Auto DC correction available: %v\n", channel, available)
	if available {
		offsetMode := dev.GetDCOffsetMode(direction, channel)
		fmt.Printf("Channel#%d Stream Correction Auto DEC correction: %v\n", channel, offsetMode)
	}

	available = dev.HasDCOffset(direction, channel)
	fmt.Printf("Channel#%d Stream Correction DC Correction available: %v\n", channel, available)
	if available {
		I, Q, err := dev.GetDCOffset(direction, channel)
		fmt.Printf("Channel#%d Stream Correction DC correction I: %v, Q: %v, err: %v\n", channel, I, Q, err)
	}

	available = dev.HasIQBalance(direction, channel)
	fmt.Printf("Channel#%d Stream Correction IQ Balance available: %v\n", channel, available)
	if available {
		I, Q, err := dev.GetIQBalance(direction, channel)
		fmt.Printf("Channel#%d Stream Correction IQ Balnance I: %v, Q: %v, err: %v\n", channel, I, Q, err)
	}

	available = dev.HasFrequencyCorrection(direction, channel)
	fmt.Printf("Channel#%d Stream Correction Frequency correction available: %v\n", channel, available)
	if available {
		frequencyCorrection := dev.GetFrequencyCorrection(direction, channel)
		fmt.Printf("Channel#%d Stream Correction Frequency correction: %v PPM\n", channel, frequencyCorrection)
	}

	// Channel Sensor
	sensors := dev.ListChannelSensors(direction, channel)
	if len(sensors) > 0 {
		for i, sensor := range sensors {
			fmt.Printf("Channel#%d Sensor#%d: %v\n", channel, i, sensor)
		}
	} else {
		fmt.Printf("Channel#%d Sensors: [none]\n", channel)
	}
}

// receiveSomeData receives CS8 data from stream 0
func receiveSomeData(dev *device.SDRDevice) {
	fmt.Println("\n---------------")
	fmt.Println("Data Reception")
	fmt.Println("---------------")

	// Apply settings
	if err := dev.SetSampleRate(device.DirectionRX, 0, 1e6); err != nil {
		log.Fatal(fmt.Printf("SetSampleRate fail: error: %v\n", err))
	}
	if err := dev.SetFrequency(device.DirectionRX, 0, 99.9e6, nil); err != nil {
		log.Fatal(fmt.Printf("SetFrequency fail: error: %v\n", err))
	}

	stream, err := dev.SetupSDRStreamCS8(device.DirectionRX, []uint{0}, nil)

	if err != nil {
		log.Fatal(fmt.Printf("SetupStream fail: error: %v\n", err))
	}

	if err := stream.Activate(0, 0, 0); err != nil {
		log.Fatal(fmt.Printf("Stream Activate fail: error: %v\n", err))
	}

	mtu := stream.GetMTU()
	fmt.Printf("Stream MTU: %v\n", mtu)
	numBuffers := stream.GetNumDirectAccessBuffers()
	fmt.Printf("NumDirectAccessBuffers: %v\n", numBuffers)

	buffers := make([][]int8, 1)
	buffers[0] = make([]int8, 1024)
	flags := make([]int, 1)

	for i := 0; i < 10; i++ {
		timeNs, numElementsRead, err := stream.Read(buffers, 511, flags, 1000000)
		fmt.Printf("flags=")
		var flag int = flags[0]
		streamFlags := buildStreamFlagsString(flag)
		fmt.Printf("flags=%v\n", streamFlags)
		fmt.Printf("numElemsRead=%v, timeNS=%v, err=%v\n", numElementsRead, timeNs, err)
		if err == nil {
			for j := uint(0); j < numElementsRead; j++ {
				fmt.Printf("{%v, %v} ", buffers[0][int(2*j)], buffers[0][2*j+1])
				if (j+1)%8 == 0 {
					fmt.Println("")
				}
			}
			fmt.Println("")
		}
	}

	if err := stream.Deactivate(0, 1000000); err != nil {
		log.Fatal(fmt.Printf("Stream Deactivate fail: error: %v\n", err))
	}

	if err := stream.Close(); err != nil {
		log.Fatal(fmt.Printf("Stream close fail: error: %v\n", err))
	}
}

// buildStreamFlagsString builds a string of stream flag names.
func buildStreamFlagsString(flag int) string {
	var haveFlag = false
	var flagStringBuilder strings.Builder
	addFlagStringToStringBuilder(flag, "EndBurst", device.StreamFlagEndBurst, &haveFlag, &flagStringBuilder)
	addFlagStringToStringBuilder(flag, "HasTime", device.StreamFlagHasTime, &haveFlag, &flagStringBuilder)
	addFlagStringToStringBuilder(flag, "EndAbrupt", device.StreamFlagEndAbrupt, &haveFlag, &flagStringBuilder)
	addFlagStringToStringBuilder(flag, "OnePacket", device.StreamFlagOnePacket, &haveFlag, &flagStringBuilder)
	addFlagStringToStringBuilder(flag, "MoreFragments", device.StreamFlagMoreFragments, &haveFlag, &flagStringBuilder)
	addFlagStringToStringBuilder(flag, "WaitTrigger", device.StreamFlagWaitTrigger, &haveFlag, &flagStringBuilder)
	if flagStringBuilder.Len() > 0 {
		return flagStringBuilder.String()
	} else {
		return "[none]"
	}
}

// addFlagStringToStringBuilder adds stream flag name to string.Builder if flag is set.
func addFlagStringToStringBuilder(flag int, flagName string, testFlag device.StreamFlag, haveFlags *bool, flagString *strings.Builder) {
	if flag&int(testFlag) == int(testFlag) {
		if *haveFlags {
			flagString.WriteString(", ")
		}
		*haveFlags = true
		flagString.WriteString(flagName)
	}
}

// logSoapy receives and prints Soapy messages to be logged
func logSoapy(level sdrlogger.SDRLogLevel, message string) {
	levelStr := "Unknown"
	switch level {
	case sdrlogger.Fatal:
		levelStr = "Fatal"
	case sdrlogger.Critical:
		levelStr = "Critical"
	case sdrlogger.Error:
		levelStr = "Error"
	case sdrlogger.Warning:
		levelStr = "Warning"
	case sdrlogger.Notice:
		levelStr = "Notice"
	case sdrlogger.Info:
		levelStr = "Info"
	case sdrlogger.Debug:
		levelStr = "Debug"
	case sdrlogger.Trace:
		levelStr = "Trace"
	case sdrlogger.SSI:
		levelStr = "SSI"
	}

	fmt.Printf("Soapy Logged: [%v] %v\n", levelStr, message)
}
