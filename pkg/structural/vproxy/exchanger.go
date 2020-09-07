// Package vproxy Паттерн "Виртуальный заместитель".
package vproxy

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Exchanger Интерфейс табло курса валют.
type Exchanger interface {
	Show()
	GetRates() map[string]float64
}

// Табло курс валют.
type exchanger struct {
	rates map[string]float64
}

// Show Отобразить текущий курс валют.
func (e *exchanger) Show() {
	fmt.Println()

	for currency, rate := range e.rates {
		fmt.Printf("%s -> RUB\t%f\n", currency, rate)
	}

	fmt.Println()
}

// GetRates Получить курс валют.
func (e *exchanger) GetRates() map[string]float64 {
	return e.rates
}

// Загрузить курс валют.
func (e *exchanger) loadRates() {
	e.leadRate("USD")
	e.leadRate("EUR")
}

// Загрузить курс рубля к заданной валюте.
func (e *exchanger) leadRate(currency string) {
	url := fmt.Sprintf("https://api.exchangeratesapi.io/latest?base=%s&symbols=RUB", currency)
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
func newExchanger() Exchanger {
	ex := &exchanger{make(map[string]float64)}
	ex.loadRates()

	return ex
}
