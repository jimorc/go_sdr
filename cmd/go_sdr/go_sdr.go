package main

// go:generate fyne go_sdr -o bundled.go $GOPATH/go_sdr/cmd/go_sdr/images/start.svg

import (
	"internal/gosdrgui"
	"internal/soapy_logging"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/pothosware/go-soapy-sdr/pkg/sdrlogger"
)

func main() {
	soapy_logging.CreateSoapyLogfileName("go_sdr.log")
	sdrlogger.RegisterLogHandler(soapy_logging.LogSoapy)
	sdrlogger.SetLogLevel(sdrlogger.Info)
	sdrlogger.Log(sdrlogger.Info, "go_sdr Logging")

	sdrApp := app.New()
	mainWindow := sdrApp.NewWindow("go_sdr")

	startStop := gosdrgui.NewStartStopToolbarAction()
	toolBar := widget.NewToolbar(startStop)

	mainWindow.SetContent(toolBar)

	mainWindow.ShowAndRun()

}
