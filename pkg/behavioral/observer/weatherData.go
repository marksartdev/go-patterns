package observer

import (
	"fmt"
	"go-patterns/internal/log"
	"io"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

// WeatherDater Интерфейс метеостанции
type WeatherDater interface {
	getTemperature() float64
	SetTemperature()
	getHumidity() float64
	SetHumidity()
	getPressure() float64
	SetPressure()
	SetReader(io.Reader)
	MeasurementsChanged()
}

// Метеостанция
type weatherData struct {
	logger      *logrus.Logger
	reader      io.Reader
	temperature float64
	humidity    float64
	pressure    float64
}

// GetTemperature Вернуть текущую температуру
func (w *weatherData) getTemperature() float64 {
	return w.temperature
}

// SetTemperature Установить температуру
func (w *weatherData) SetTemperature() {
	w.temperature = w.input("Введите температуру")
}

// GetHumidity Вернуть текущую влажность
func (w *weatherData) getHumidity() float64 {
	return w.humidity
}

// SetHumidity Установить влажность
func (w *weatherData) SetHumidity() {
	w.humidity = w.input("Введите влажность")
}

// GetPressure Вернуть текущее давление
func (w *weatherData) getPressure() float64 {
	return w.pressure
}

// SetPressure Установить давление
func (w *weatherData) SetPressure() {
	w.pressure = w.input("Введите давление")
}

// MeasurementsChanged Вызывается при каждом обновлении показаний датчиков
func (w *weatherData) MeasurementsChanged() {
	// Здесь будет реализация
}

// Получить данные
func (w *weatherData) input(label string) float64 {
	var numberStr string
	var number float64

input:
	if w.reader == os.Stdin {
		fmt.Println(label)
	}

	_, err := fmt.Fscanln(w.reader, &numberStr)
	if err != nil {
		w.logger.Errorln(err)

		return number
	}

	number, err = strconv.ParseFloat(numberStr, 64)
	if err != nil {
		w.logger.Warnln("Введено некорректное значение")

		goto input
	}

	return number
}

// SetReader Установить reader
func (w *weatherData) SetReader(reader io.Reader) {
	w.reader = reader
}

// NewWeatherData Создать weatherData
func NewWeatherData() WeatherDater {
	logger := log.NewLogger()
	wd := new(weatherData)
	wd.logger = logger
	wd.reader = os.Stdin

	return wd
}
