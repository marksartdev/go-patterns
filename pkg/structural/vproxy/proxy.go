package vproxy

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

// Заместитель табло курса валют.
type proxyExchanger struct {
	exchanger Exchanger
	date      string
	state     state
	initState state
	mainState state
	writer    io.Writer
}

// Show Отобразить текущий курс валют.
func (e *proxyExchanger) Show() {
	e.state.show()
}

// GetRates Получить курс валют.
func (e *proxyExchanger) GetRates() map[string]float64 {
	return e.state.getRates()
}

// SetWriter Установить writer.
func (e *proxyExchanger) SetWriter(writer io.Writer) {
	e.writer = writer
	if e.exchanger != nil {
		e.exchanger.SetWriter(writer)
	}
}

func (e *proxyExchanger) write(msg string) {
	_, err := fmt.Fprintln(e.writer, msg)
	if err != nil {
		log.Error(err)
	}
}

// Загрузить настоящее табло курса валют.
func (e *proxyExchanger) initExchanger() {
	e.exchanger = newExchanger(e.date)
	e.exchanger.SetWriter(e.writer)
	e.state = e.mainState
	e.Show()
}

// NewExchanger Создать табло курса валют из заместителя.
func NewExchanger(date string) Exchanger {
	proxy := &proxyExchanger{}
	proxy.date = date
	proxy.initState = initStage{proxy}
	proxy.mainState = mainState{proxy}
	proxy.state = proxy.initState
	proxy.writer = os.Stdout

	return proxy
}
