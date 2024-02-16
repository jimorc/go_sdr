package soapysdr

// #cgo pkg-config: SoapySDR
// #include <SoapySDR/Device.h>
import "C"

// DeviceLastStatus returns the last status set by a SoapySDR Device API call. The status code is cleared on entry to each
// Device call. A return status of 0 indicates that the request succeeded, and a non-0 status indicates that the request failed.
//
// Call DeviceLastError() to retrieve the status as an error message strig.
func DeviceLastStatus() int {
	// TODO: Call SoapySDRDevice_lastStatus()
	return 0
}
