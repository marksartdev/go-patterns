package observer

import (
	"fmt"
	"io"
	"math/rand"
	"os"

	"github.com/gonum/floats"
)

// Интерфейс визуального элемента
type displayer interface {
	Display() error
}

// DisplayElement Интерфейс визуального элемента, способного быть наблюдателем
type DisplayElement interface {
	writer
	displayer
	observer
}

// Экран текущего состояния
type currentConditionsDisplay struct {
	writer      io.Writer
	temperature float64
	humidity    float64
	pressure    float64
}

// SetOutput Установить writer
func (d *currentConditionsDisplay) SetOutput(writer io.Writer) {
	d.writer = writer
}

// Display Отобразить экран
func (d *currentConditionsDisplay) Display() error {
	text := "Current conditions:\n"
	text += fmt.Sprintf("\tTemperature: %.1f\n", d.temperature)
	text += fmt.Sprintf("\tHumidity: %.1f\n", d.humidity)
	text += fmt.Sprintf("\tPressure: %.1f\n", d.pressure)

	_, err := fmt.Fprintln(d.writer, text)
	return err
}

// Update Обновить данные
func (d *currentConditionsDisplay) Update(data *measurements) {
	d.temperature = data.temperature
	d.humidity = data.humidity
	d.pressure = data.pressure
}

// NewCurrentConditionsDisplay Создать экран текущего состояния
func NewCurrentConditionsDisplay() Displayer {
	display := new(currentConditionsDisplay)
	display.writer = os.Stdout

	return display
}

// Экран статистики
type statisticsDisplay struct {
	writer      io.Writer
	temperature []float64
	humidity    []float64
	pressure    []float64
}

// SetOutput Установить writer
func (d *statisticsDisplay) SetOutput(writer io.Writer) {
	d.writer = writer
}

// Display Отобразить экран
func (d *statisticsDisplay) Display() error {
	var temperatureMax, temperatureMin, temperatureAvg float64
	var humidityMax, humidityMin, humidityAvg float64
	var pressureMax, pressureMin, pressureAvg float64

	if len(d.temperature) > 0 {
		temperatureMax = floats.Max(d.temperature)
		temperatureMin = floats.Min(d.temperature)
		temperatureAvg = floats.Sum(d.temperature) / float64(len(d.temperature))
	}

	if len(d.humidity) > 0 {
		humidityMax = floats.Max(d.humidity)
		humidityMin = floats.Min(d.humidity)
		humidityAvg = floats.Sum(d.humidity) / float64(len(d.humidity))
	}

	if len(d.pressure) > 0 {
		pressureMax = floats.Max(d.pressure)
		pressureMin = floats.Min(d.pressure)
		pressureAvg = floats.Sum(d.pressure) / float64(len(d.pressure))
	}

	text := "Statistics:\n"
	text += fmt.Sprintf("\tTemperature (max/min/avg): %.1f/%.1f/%.1f\n", temperatureMax, temperatureMin, temperatureAvg)
	text += fmt.Sprintf("\tHumidity (max/min/avg): %.1f/%.1f/%.1f\n", humidityMax, humidityMin, humidityAvg)
	text += fmt.Sprintf("\tPressure (max/min/avg): %.1f/%.1f/%.1f\n", pressureMax, pressureMin, pressureAvg)

	_, err := fmt.Fprintln(d.writer, text)
	return err
}

// Update Обновить данные
func (d *statisticsDisplay) Update(data *measurements) {
	d.temperature = append(d.temperature, data.temperature)
	d.humidity = append(d.humidity, data.humidity)
	d.pressure = append(d.pressure, data.pressure)
}

// NewStatisticsDisplay Создать экран статистики
func NewStatisticsDisplay() Displayer {
	display := new(statisticsDisplay)
	display.writer = os.Stdout

	return display
}

// Экран прогноза
type forecastDisplay struct {
	writer      io.Writer
	temperature float64
	humidity    float64
	pressure    float64
}

// SetOutput Установить writer
func (d *forecastDisplay) SetOutput(writer io.Writer) {
	d.writer = writer
}

// Display Отобразить экран
func (d *forecastDisplay) Display() error {
	text := "Forecast:\n"
	text += fmt.Sprintf("\tTemperature: %.1f\n", d.temperature)
	text += fmt.Sprintf("\tHumidity: %.1f\n", d.humidity)
	text += fmt.Sprintf("\tPressure: %.1f\n", d.pressure)

	_, err := fmt.Fprintln(d.writer, text)
	return err
}

// Update Обновить данные
func (d *forecastDisplay) Update(data *measurements) {
	d.temperature = data.temperature
	d.humidity = data.humidity
	d.pressure = data.pressure

	d.makeForecast()
}

// Сделать прогноз
func (d *forecastDisplay) makeForecast() {
	rand.Seed(512)

	d.temperature = d.getCoefficient() * d.temperature
	d.humidity = d.getCoefficient() * d.humidity
	d.pressure = d.getCoefficient() * d.pressure
}

// Получить коэффициент для прогноза
func (d *forecastDisplay) getCoefficient() float64 {
	return 0.7 + rand.Float64()*(1.3-0.7)
}

// NewForecastDisplay Создать экран прогноза
func NewForecastDisplay() Displayer {
	display := new(forecastDisplay)
	display.writer = os.Stdout

	return display
}
