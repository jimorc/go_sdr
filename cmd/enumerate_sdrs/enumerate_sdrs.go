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
	fmt.Printf("%v", *dev)
}

// logSoapy is a function that is used to receive Soapy messages to be logged
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
