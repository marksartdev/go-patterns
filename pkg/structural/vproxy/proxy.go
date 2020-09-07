package vproxy

import "fmt"

// Заместитель табло курса валют.
type exchangerProxy struct {
	exchanger Exchanger
}

// Show Отобразить текущий курс валют.
func (e *exchangerProxy) Show() {
	if e.exchanger != nil {
		e.exchanger.Show()
	} else {
		fmt.Println("Loading...")

		go e.initExchanger()
	}
}

// Загрузить настоящее табло курса валют.
func (e *exchangerProxy) initExchanger() {
	e.exchanger = newExchanger()
	e.Show()
}

// NewExchanger Создать табло курса валют из заместителя.
func NewExchanger() Exchanger {
	return &exchangerProxy{}
}
