package observer

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

// WeatherDater Интерфейс метеостанции
type WeatherDater interface {
	SetInput(io.Reader)
	SetOutput(io.Writer)
	getTemperature() float64
	SetTemperature() error
	getHumidity() float64
	SetHumidity() error
	getPressure() float64
	SetPressure() error
	MeasurementsChanged()
}

// Метеостанция
type weatherData struct {
	reader      io.Reader
	writer      io.Writer
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

// Вернуть текущую температуру
func (w *weatherData) getTemperature() float64 {
	return w.temperature
}

// SetTemperature Установить температуру
func (w *weatherData) SetTemperature() error {
	temperature, err := w.input("Введите температуру")
	if err != nil {
		return err
	}

	w.temperature = temperature

	return nil
}

// Вернуть текущую влажность
func (w *weatherData) getHumidity() float64 {
	return w.humidity
}

// SetHumidity Установить влажность
func (w *weatherData) SetHumidity() error {
	humidity, err := w.input("Введите влажность")
	if err != nil {
		return err
	}

	w.humidity = humidity

	return nil
}

// Вернуть текущее давление
func (w *weatherData) getPressure() float64 {
	return w.pressure
}

// SetPressure Установить давление
func (w *weatherData) SetPressure() error {
	pressure, err := w.input("Введите давление")
	if err != nil {
		return err
	}

	w.pressure = pressure

	return nil
}

// MeasurementsChanged Вызывается при каждом обновлении показаний датчиков
func (w *weatherData) MeasurementsChanged() {
	// Здесь будет реализация
}

// Получить данные
func (w *weatherData) input(label string) (float64, error) {
	var numberStr string
	var number float64

	_, err := fmt.Fprintln(w.writer, label)
	if err != nil {
		return 0, err
	}

	_, err = fmt.Fscanln(w.reader, &numberStr)
	if err != nil {
		return 0, err
	}

	number, err = strconv.ParseFloat(numberStr, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}

// NewWeatherData Создать weatherData
func NewWeatherData() WeatherDater {
	wd := new(weatherData)
	wd.reader = os.Stdin
	wd.writer = os.Stdout

	return wd
}
