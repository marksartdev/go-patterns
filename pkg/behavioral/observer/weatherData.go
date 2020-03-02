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
	GetTemperature() float64
	SetTemperature(io.Reader)
	GetHumidity() float64
	SetHumidity(io.Reader)
	GetPressure() float64
	SetPressure(io.Reader)
	MeasurementsChanged()
}

// Субъект данных
type weatherData struct {
	logger      *logrus.Logger
	temperature float64
	humidity    float64
	pressure    float64
}

// GetTemperature Вернуть текущую температуру
func (w *weatherData) GetTemperature() float64 {
	return w.temperature
}

// SetTemperature Установить температуру
func (w *weatherData) SetTemperature(reader io.Reader) {
	w.temperature = read("Введите температуру", reader, w.logger)
}

// GetHumidity Вернуть текущую влажность
func (w *weatherData) GetHumidity() float64 {
	return w.humidity
}

// SetHumidity Установить влажность
func (w *weatherData) SetHumidity(reader io.Reader) {
	w.humidity = read("Введите влажность", reader, w.logger)
}

// GetPressure Вернуть текущее давление
func (w *weatherData) GetPressure() float64 {
	return w.pressure
}

// SetPressure Установить давление
func (w *weatherData) SetPressure(reader io.Reader) {
	w.pressure = read("Введите давление", reader, w.logger)
}

// MeasurementsChanged Вызывается при каждом обновлении показаний датчиков
func (w *weatherData) MeasurementsChanged() {
	// Здесь будет реализация
}

// NewWeatherData Создать weatherData
func NewWeatherData() WeatherDater {
	logger := log.NewLogger()
	wd := new(weatherData)
	wd.logger = logger

	return wd
}

// Получить данные
func read(label string, reader io.Reader, logger *logrus.Logger) float64 {
	var numberStr string
	var number float64

input:
	if reader == os.Stdin {
		fmt.Println(label)
	}

	_, err := fmt.Fscanln(reader, &numberStr)
	if err != nil {
		logger.Fatalln(err)
	}

	number, err = strconv.ParseFloat(numberStr, 64)
	if err != nil {
		logger.Warnln("Введено некорректное значение")
		goto input
	}

	return number
}
