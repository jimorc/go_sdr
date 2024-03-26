package gosdrgui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// NewMainWindow creates the main window for the go_sdr app.
func NewMainWindow(sdrApp fyne.App) fyne.Window {
	mainWin := sdrApp.NewWindow("go_sdr")
	startStop := NewStartStopToolbarAction()
	toolBar := widget.NewToolbar(startStop)

	mainWin.SetContent(toolBar)

	return mainWin
}
