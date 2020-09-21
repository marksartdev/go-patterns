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
	progressLen   = 21
	progressSleep = 5
	width         = 200
	height        = 0
)

// Интерфейс представления.
type djViewInterface interface {
	init()
	enableStopButton()
	disableStopButton()
	enableStartButton()
	disableStartButton()
	Run()
}

// Представление.
type djView struct {
	model      BeatModelInterface
	controller ControllerInterface
	viewW      fyne.Window
	controlW   fyne.Window
	progress   *widget.Label
	bpm        *widget.Label
	input      *widget.Entry
	startBtm   *widget.Button
	stopBtm    *widget.Button
	quiteBtm   *widget.Button
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
	d.startBtm = widget.NewButton("Start", d.controller.start)
	d.stopBtm = widget.NewButton("Stop", d.controller.stop)
	d.quiteBtm = widget.NewButton("Quit", d.quit)
	menu := widget.NewHBox(d.startBtm, d.stopBtm, d.quiteBtm)

	label := widget.NewLabel("Enter BPM:")
	d.input = widget.NewEntry()
	input := widget.NewHBox(label, d.input)

	setButton := widget.NewButton("Set", d.setBPM)

	decButton := widget.NewButton("<<", d.controller.decreaseBPM)
	incButton := widget.NewButton(">>", d.controller.increaseBPM)

	stepButtons := widget.NewHBox(decButton, incButton)

	d.controlW = a.NewWindow("Control")
	d.controlW.SetContent(widget.NewVBox(menu, input, setButton, stepButtons))

	d.controlW.Resize(fyne.Size{Width: width, Height: height})
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

// Разблокировать кнопку остановки.
func (d *djView) enableStopButton() {
	d.stopBtm.Enable()
}

// Заблокировать кнопку остановки.
func (d *djView) disableStopButton() {
	d.stopBtm.Disable()
}

// Разблокировать кнопку старта.
func (d *djView) enableStartButton() {
	d.startBtm.Enable()
}

// Заблокировать кнопку старта.
func (d *djView) disableStartButton() {
	d.startBtm.Disable()
}

// Закрыть приложение.
func (d *djView) quit() {
	d.controller.stop()

	d.viewW.Close()
	d.controlW.Close()
}

// Установить BPM.
func (d *djView) setBPM() {
	bpm, err := strconv.Atoi(d.input.Text)
	if err != nil {
		log.Error(err)

		return
	}

	d.controller.setBPM(bpm)

	d.input.Text = ""
	d.input.Refresh()
}

// Run Запустить.
func (d *djView) Run() {
	d.viewW.Show()
	d.controlW.ShowAndRun()
}

// Создать представление.
func newDJView(controller ControllerInterface, model BeatModelInterface) djViewInterface {
	view := &djView{}
	view.controller = controller
	view.model = model

	view.model.registerBeatObserver(view)
	view.model.registerBPMObserver(view)

	return view
}
