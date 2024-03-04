package main

import (
	"fmt"
	"log"

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
			"label": dev["label"],
		}
	}

	// Open all devices
	devs, err := device.MakeList(deviceArgs)
	if err != nil {
		log.Panic(err)
	}

	for i, dev := range devs {
		fmt.Printf("***************\n")
		fmt.Printf("Device: %v\n", devices[i]["label"])
		fmt.Printf("***************\n")

		displayDetails(dev)
	}

	// Close all devices
	err = device.UnmakeList(devs)
	if err != nil {
		log.Panic(err)
	}
	sdrlogger.Log(sdrlogger.Trace, "All devices closed")
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

// displayTimeSources lists all of a device's time sources.
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
