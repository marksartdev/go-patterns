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
	SetTemperature(io.Reader)
	getHumidity() float64
	SetHumidity(io.Reader)
	getPressure() float64
	SetPressure(io.Reader)
	MeasurementsChanged()
}

// Метеостанция
type weatherData struct {
	logger      *logrus.Logger
	temperature float64
	humidity    float64
	pressure    float64
}

// GetTemperature Вернуть текущую температуру
func (w *weatherData) getTemperature() float64 {
	return w.temperature
}

// SetTemperature Установить температуру
func (w *weatherData) SetTemperature(reader io.Reader) {
	w.temperature = w.input("Введите температуру", reader)
}

// GetHumidity Вернуть текущую влажность
func (w *weatherData) getHumidity() float64 {
	return w.humidity
}

// SetHumidity Установить влажность
func (w *weatherData) SetHumidity(reader io.Reader) {
	w.humidity = w.input("Введите влажность", reader)
}

// GetPressure Вернуть текущее давление
func (w *weatherData) getPressure() float64 {
	return w.pressure
}

// SetPressure Установить давление
func (w *weatherData) SetPressure(reader io.Reader) {
	w.pressure = w.input("Введите давление", reader)
}

// MeasurementsChanged Вызывается при каждом обновлении показаний датчиков
func (w *weatherData) MeasurementsChanged() {
	// Здесь будет реализация
}

// Получить данные
func (w *weatherData) input(label string, reader io.Reader) float64 {
	var numberStr string
	var number float64

input:
	if reader == os.Stdin {
		fmt.Println(label)
	}

	_, err := fmt.Fscanln(reader, &numberStr)
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
	wd.logger = logger

	return wd
}
