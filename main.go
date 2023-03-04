package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Go Health Checker")
	w.Resize(fyne.NewSize(600, 600))
	w.SetContent(widget.NewLabel("Health Check aplication!ƒ©"))
	w.ShowAndRun()
}
