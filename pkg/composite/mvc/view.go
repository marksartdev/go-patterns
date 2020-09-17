package mvc

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"

	log "github.com/sirupsen/logrus"
)

const (
	progressLen   = 10
	progressSleep = 50
)

// DJView Интерфейс представления.
type DJView interface {
	Run()
}

const (
	width  = 200
	height = 0
)

// Представление.
type djView struct {
	model      beatModelInterface
	controller controllerInterface
	viewW      fyne.Window
	controlW   fyne.Window
	progress   *widget.Label
	bpm        *widget.Label
	input      *widget.Entry
}

// Создать UI.
func (d *djView) init() {
	a := app.New()

	d.createView(a)
	d.createControls(a)
}

// Создать UI визуального представления.
func (d *djView) createView(a fyne.App) {
	d.progress = widget.NewLabel("")
	d.bpm = widget.NewLabel("offline")

	d.viewW = a.NewWindow("View")
	d.viewW.SetContent(widget.NewVBox(d.progress, d.bpm))

	d.viewW.Resize(fyne.Size{Width: width, Height: height})
}

// Создать UI элементов управления.
func (d *djView) createControls(a fyne.App) {
	startItem := fyne.NewMenuItem("Start", d.controller.start)
	stopItem := fyne.NewMenuItem("Stop", d.controller.stop)
	quitItem := fyne.NewMenuItem("Quit", d.quit)
	menu := fyne.NewMenu("DJ Control", startItem, stopItem, quitItem)
	mainMenu := fyne.NewMainMenu(menu)

	label := widget.NewLabel("Enter BPM:")
	d.input = widget.NewEntry()
	input := widget.NewHBox(label, d.input)

	setButton := widget.NewButton("Set", d.setBPM)

	decButton := widget.NewButton("<<", d.controller.decreaseBPM)
	incButton := widget.NewButton(">>", d.controller.increaseBPM)

	stepButtons := widget.NewHBox(decButton, incButton)

	d.controlW = a.NewWindow("Control")
	d.controlW.SetMainMenu(mainMenu)
	d.controlW.SetContent(widget.NewVBox(input, setButton, stepButtons))

	d.controlW.Resize(fyne.Size{Width: width, Height: height})
}

// Закрыть приложение.
func (d *djView) quit() {
	d.viewW.Close()
	d.controlW.Close()
}

// Установить BPM.
func (d *djView) setBPM() {
	text := d.input.Text

	bpm, err := strconv.Atoi(text)
	if err != nil {
		log.Error(err)

		return
	}

	d.controller.setBPM(bpm)

	d.input.Text = ""
	d.input.Refresh()
}

// Обновить BPM.
func (d *djView) updateBPM() {
	bpm := d.model.getBPM()
	if bpm == 0 {
		d.bpm.SetText("offline")
	} else {
		d.bpm.SetText(fmt.Sprintf("Current BPM: %d", bpm))
	}
}

// Отобразить удар.
func (d *djView) updateBeat() {
	for i := 1; i <= progressLen; i++ {
		d.progress.SetText(strings.Repeat("=", i))
		time.Sleep(progressSleep * time.Millisecond)
	}

	for i := 10; i >= 0; i-- {
		d.progress.SetText(strings.Repeat("=", i))
		time.Sleep(progressSleep * time.Millisecond)
	}
}

// Run Запустить.
func (d *djView) Run() {
	d.init()

	d.viewW.Show()
	d.controlW.ShowAndRun()
}

// NewDJView Создать представление.
func NewDJView() DJView {
	view := &djView{}
	view.model = newBeatModel()
	view.controller = newBeatController()

	view.model.registerBeatObserver(view)
	view.model.registerBPMObserver(view)

	return view
}
