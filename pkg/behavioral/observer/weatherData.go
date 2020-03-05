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
	SetInput(io.Reader)
	SetOutput(io.Writer)
	SetLogOutput(io.Writer)
	getTemperature() float64
	SetTemperature()
	getHumidity() float64
	SetHumidity()
	getPressure() float64
	SetPressure()
	MeasurementsChanged()
}

// Метеостанция
type weatherData struct {
	reader      io.Reader
	writer      io.Writer
	logger      *logrus.Logger
	temperature float64
	humidity    float64
	pressure    float64
}

// SetInput Установить reader
func (w *weatherData) SetInput(reader io.Reader) {
	w.reader = reader
}

// SetOutput Установить writer
func (w *weatherData) SetOutput(writer io.Writer) {
	w.writer = writer
}

// SetLogOutput Установить writer для логгера
func (w *weatherData) SetLogOutput(writer io.Writer) {
	w.logger.SetOutput(writer)
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
	_, err := fmt.Fprintln(w.writer, label)
	if err != nil {
		w.logger.Errorln(err)

		return number
	}

	_, err = fmt.Fscanln(w.reader, &numberStr)
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

// NewWeatherData Создать weatherData
func NewWeatherData() WeatherDater {
	logger := log.NewLogger()
	wd := new(weatherData)
	wd.reader = os.Stdin
	wd.writer = os.Stdout
	wd.logger = logger

	return wd
}
