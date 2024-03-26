package main

// go:generate fyne go_sdr -o bundled.go $GOPATH/go_sdr/cmd/go_sdr/images/start.svg

import (
	"internal/gosdrgui"
	"internal/soapy_logging"

	"fyne.io/fyne/v2/app"
	"github.com/pothosware/go-soapy-sdr/pkg/sdrlogger"
)

func main() {
	soapy_logging.CreateSoapyLogfileName("go_sdr.log")
	sdrlogger.RegisterLogHandler(soapy_logging.LogSoapy)
	sdrlogger.SetLogLevel(sdrlogger.Info)
	sdrlogger.Log(sdrlogger.Info, "go_sdr Logging")

	sdrApp := app.New()
	mainWindow := gosdrgui.NewMainWindow(sdrApp)
	mainWindow.ShowAndRun()
}
