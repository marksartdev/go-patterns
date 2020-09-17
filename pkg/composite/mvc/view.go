package mvc

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"strings"
	"time"
)

const progressLen = 10

// DJView Интерфейс представления.
type DJView interface {
	Run()
}

// Представление.
type djView struct {
	model      beatModelInterface
	controller controllerInterface
	viewW      fyne.Window
	controlW   fyne.Window
	viewElms   viewElms
}

// Создание UI.
func (d *djView) createView() {
	d.viewElms = newViewElms()

	a := app.New()

	d.viewW = a.NewWindow("View")
	d.viewW.SetContent(widget.NewVBox(d.viewElms.progress, d.viewElms.label))
	d.viewW.Resize(fyne.Size{Width: 200, Height: 75})
}

// Обновить BPM.
func (d *djView) updateBPM() {
	bpm := d.model.getBPM()
	if bpm == 0 {
		d.viewElms.label.SetText("offline")
	} else {
		d.viewElms.label.SetText(fmt.Sprintf("Current BPM: %d", bpm))
	}
}

// Отобразить удар.
func (d *djView) updateBeat() {
	for i := 1; i <= progressLen; i++ {
		d.viewElms.progress.SetText(strings.Repeat("=", i))
		time.Sleep(50 * time.Millisecond)
	}

	for i := 10; i >= 0; i-- {
		d.viewElms.progress.SetText(strings.Repeat("=", i))
		time.Sleep(50 * time.Millisecond)
	}
}

// Run Запустить.
func (d *djView) Run() {
	d.createView()

	d.viewW.ShowAndRun()
	//d.controlW.ShowAndRun()
}

// NewDJView Создать представление.
func NewDJView() DJView {
	view := &djView{}
	view.model = newBeatModel()
	// todo set controller

	view.model.registerBeatObserver(view)
	view.model.registerBPMObserver(view)

	return view
}
