// Package vproxy "Виртуальный заместитель" - разновидность паттерна "Заместитель".
package vproxy

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// Exchanger Интерфейс табло курса валют.
type Exchanger interface {
	Show()
	GetRates() map[string]float64
	SetWriter(writer io.Writer)
	write(msg string)
}

// Табло курс валют.
type exchanger struct {
	rates  map[string]float64
	url    string
	writer io.Writer
}

// Show Отобразить текущий курс валют.
func (e *exchanger) Show() {
	e.write("")

	for currency, rate := range e.rates {
		e.write(fmt.Sprintf("%s -> RUB\t%f", currency, rate))
	}

	e.write("")
}

// GetRates Получить курс валют.
func (e *exchanger) GetRates() map[string]float64 {
	return e.rates
}

// SetWriter Установить writer.
func (e *exchanger) SetWriter(writer io.Writer) {
	e.writer = writer
}

func (e *exchanger) write(msg string) {
	_, err := fmt.Fprintln(e.writer, msg)
	if err != nil {
		log.Error(err)
	}
}

// Загрузить курс валют.
func (e *exchanger) loadRates() {
	e.leadRate("USD")
	e.leadRate("EUR")
}

// Загрузить курс рубля к заданной валюте.
func (e *exchanger) leadRate(currency string) {
	url := fmt.Sprintf(e.url, currency)
	client := http.Client{}

	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	// Имитация долгой загрузки.
	time.Sleep(time.Second)

	result := map[string]interface{}{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Error(err)

		err = resp.Body.Close()
		if err != nil {
			log.Error(err)
		}

		log.Exit(1)
	}

	if rates, ok := result["rates"].(map[string]interface{}); ok {
		if rate, ok := rates["RUB"].(float64); ok {
			e.rates[currency] = rate
		}
	}
}

// Создать табло курса валют.
func newExchanger(date string) Exchanger {
	ex := &exchanger{
		make(map[string]float64),
		fmt.Sprintf("https://api.exchangeratesapi.io/%s?base=%%s&symbols=RUB", date),
		os.Stdout,
	}
	ex.loadRates()

	return ex
}
