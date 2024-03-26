package main

// go:generate fyne go_sdr -o bundled.go $GOPATH/go_sdr/cmd/go_sdr/images/start.svg

import (
	"fmt"
	"internal/gui"
	"internal/soapy_logging"
	"log"

	"fyne.io/fyne/v2"
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

	start, err := fyne.LoadResourceFromPath("images/start.svg")
	if err != nil {
		log.Fatal(err)
	}

	stop, err := fyne.LoadResourceFromPath("images/stop.svg")
	if err != nil {
		log.Fatal(err)
	}

	startAction := gui.NewTwoStateToolbarAction(start, stop, startButtonTapped, stopButtonTapped)
	toolBar := widget.NewToolbar(startAction)

	mainWindow.SetContent(toolBar)

	mainWindow.ShowAndRun()

}

func startButtonTapped() {
	fmt.Println("In startButtonTapped")
}

func stopButtonTapped() {
	fmt.Println("In stopButtonTapped")
}
