package soapy_logging

import (
	"fmt"
	"log"
	"os"

	"github.com/pothosware/go-soapy-sdr/pkg/sdrlogger"
)

var soapyLogfileName string

func CreateSoapyLogfileName(name string) {
	soapyLogfileName = name
	logFile, err := os.Create(soapyLogfileName)
	if err != nil {
		log.Fatal(err)
	}
	err = logFile.Close()
	if err != nil {
		log.Fatal(err)
	}

}

// logSoapy receives and prints Soapy messages to be logged to the log file
func LogSoapy(level sdrlogger.SDRLogLevel, message string) {
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
	logFile, err := os.OpenFile(soapyLogfileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	_, err = logFile.WriteString(fmt.Sprintf("Soapy Logged: [%v] %v\n", levelStr, message))
	if err != nil {
		log.Panic(err)
	}
}
