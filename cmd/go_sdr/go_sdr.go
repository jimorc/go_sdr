package main

import (
	"internal/soapy_logging"

	"github.com/pothosware/go-soapy-sdr/pkg/sdrlogger"
)

func main() {
	soapy_logging.CreateSoapyLogfileName("go_sdr.log")
	sdrlogger.RegisterLogHandler(soapy_logging.LogSoapy)
	sdrlogger.SetLogLevel(sdrlogger.Info)
	sdrlogger.Log(sdrlogger.Info, "go_sdr Logging")
}
