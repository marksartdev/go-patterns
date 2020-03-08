package observer

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Интерфейс получения температуры
type temperatureGetter interface {
	GetTemperature() float64
}

// Интерфейс получения влажности
type humidityGetter interface {
	GetHumidity() float64
}

// Интерфейс получения давления
type pressureGetter interface {
	GetPressure() float64
}

// Интерфейс работы с измерениями
type measurementsHandler interface {
	SetMeasurements() error
	MeasurementsChanged()
}

// WeatherDater Интерфейс источника метеоданных, способного быть субъектом
type WeatherDater interface {
	reader
	writer
	temperatureGetter
	humidityGetter
	pressureGetter
	measurementsHandler
	subject
}

// Источник метеоданных
type weatherData struct {
	reader      io.Reader
	writer      io.Writer
	observers   map[observer]struct{}
	temperature float64
	humidity    float64
	pressure    float64
}

// SetReader Установить reader
func (w *weatherData) SetReader(reader io.Reader) {
	w.reader = reader
}

// SetWriter Установить writer
func (w *weatherData) SetWriter(writer io.Writer) {
	w.writer = writer
}

// GetTemperature Получить текущую температуру
func (w *weatherData) GetTemperature() float64 {
	return w.temperature
}

// GetHumidity Получить текущую влажность
func (w *weatherData) GetHumidity() float64 {
	return w.humidity
}

// GetPressure Получить текущее давление
func (w *weatherData) GetPressure() float64 {
	return w.pressure
}

// SetMeasurements Задать измерения
func (w *weatherData) SetMeasurements() error {
	newMeasurements, err := w.input()
	if err != nil {
		return err
	}

	w.temperature = newMeasurements.temperature
	w.humidity = newMeasurements.humidity
	w.pressure = newMeasurements.pressure

	w.MeasurementsChanged()

	return nil
}

// MeasurementsChanged Вызывается при каждом обновлении показаний датчиков
func (w *weatherData) MeasurementsChanged() {
	w.NotifyObservers()
}

// RegisterObserver Регистрация нового наблюдателя
func (w *weatherData) RegisterObserver(newObserver observer) {
	w.observers[newObserver] = struct{}{}
}

// RemoveObserver Удалить наблюдателя из списка
func (w *weatherData) RemoveObserver(removedObserver observer) {
	delete(w.observers, removedObserver)
}

// NotifyObservers Оповестить наблюдателей
func (w *weatherData) NotifyObservers() {
	newMeasurements := new(measurements)
	newMeasurements.temperature = w.temperature
	newMeasurements.humidity = w.humidity
	newMeasurements.pressure = w.pressure

	for currentObserver := range w.observers {
		currentObserver.Update(newMeasurements)
	}
}

// Получить данные
func (w *weatherData) input() (*measurements, error) {
	var input string
	var numbers [3]float64

	_, err := fmt.Fprintln(w.writer, "Введите новые данные через пробел (temperature humidity pressure):")
	if err != nil {
		return nil, err
	}

	_, err = fmt.Fscanln(w.reader, &input)
	if err != nil {
		return nil, err
	}

	inputs := strings.Split(input, " ")

	for i, measurement := range inputs {
		if i > 2 {
			break
		}

		numbers[i], err = strconv.ParseFloat(measurement, 64)
		if err != nil {
			return nil, err
		}
	}

	newMeasurements := new(measurements)
	newMeasurements.temperature = numbers[0]
	newMeasurements.humidity = numbers[1]
	newMeasurements.pressure = numbers[2]

	return newMeasurements, nil
}

// NewWeatherData Создать weatherData
func NewWeatherData() WeatherDater {
	wd := new(weatherData)
	wd.reader = os.Stdin
	wd.writer = os.Stdout
	wd.observers = make(map[observer]struct{})

	return wd
}
