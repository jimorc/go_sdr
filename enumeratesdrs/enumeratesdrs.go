package main

import (
	"fmt"

	"github.com/pothosware/go-soapy-sdr/pkg/device"
)

func main() {
	// List all devices
	devices := device.Enumerate(nil)
	for i, dev := range devices {
		fmt.Printf("Found device #%v: ", i)
		for k, v := range dev {
			fmt.Printf("%v=%v, ", k, v)
		}
		fmt.Printf("\n")
	}
}
