// Package gosdrgui contains the detailed widgets used in the go_sdr app.
package gosdrgui

import (
	"fmt"
	"internal/gui"
	"log"

	"fyne.io/fyne/v2"
)

// StartStopToolbarAction is a gui.TwoStateToolbarAction that defines the actions performed by the SDR start/stop toolbar button.
type StartStopToolbarAction struct {
	action *gui.TwoStateToolbarAction
}

// NewStartStopToolbarAction creates a StartStopToolbarAction object.
func NewStartStopToolbarAction() *StartStopToolbarAction {
	startIcon, err := fyne.LoadResourceFromPath("images/start.svg")
	if err != nil {
		log.Fatal(err)
	}

	stopIcon, err := fyne.LoadResourceFromPath("images/stop.svg")
	if err != nil {
		log.Fatal(err)
	}

	startStop := StartStopToolbarAction{}
	startStop.action = gui.NewTwoStateToolbarAction(startIcon, stopIcon, startStop.startActivated, startStop.stopActivated)
	return &startStop
}

// ToolbarObject returns a pointer to the underlying TwoStateToolbarObject
func (t *StartStopToolbarAction) ToolbarObject() fyne.CanvasObject {
	return t.action.ToolbarObject()
}

func (t *StartStopToolbarAction) startActivated() {
	fmt.Println("In startActivated")
}

func (t *StartStopToolbarAction) stopActivated() {
	fmt.Println("In stopActivated")
}
