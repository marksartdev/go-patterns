package observer

import (
	"fmt"
	"io"
	"math/rand"

	"github.com/gonum/floats"
)

// Displayer Интерфейс дисплея
type Displayer interface {
	Display(io.Writer) error
}

// Дисплей "Текущее состояние"
type currentConditionsDisplay struct {
	temperature float64
	humidity    float64
	pressure    float64
}

// Display Отобразить дисплей
func (d *currentConditionsDisplay) Display(writer io.Writer) error {
	text := "Current conditions:\n"
	text += fmt.Sprintf("\tTemperature: %.1f\n", d.temperature)
	text += fmt.Sprintf("\tHumidity: %.1f\n", d.humidity)
	text += fmt.Sprintf("\tPressure: %.1f\n", d.pressure)

	_, err := fmt.Fprintln(writer, text)
	return err
}

// NewCurrentConditionsDisplay Создать дисплей "Текущее состояние"
func NewCurrentConditionsDisplay() Displayer {
	return &currentConditionsDisplay{}
}

// Дисплей "Статистика"
type statisticsDisplay struct {
	temperature []float64
	humidity    []float64
	pressure    []float64
}

// Display Отобразить дисплей
func (d *statisticsDisplay) Display(writer io.Writer) error {
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

	_, err := fmt.Fprintln(writer, text)
	return err
}

// NewStatisticsDisplay Создать дисплей "Статистика"
func NewStatisticsDisplay() Displayer {
	return &statisticsDisplay{}
}

// Дисплей "Прогноз"
type forecastDisplay struct {
	temperature float64
	humidity    float64
	pressure    float64
}

// Display Отобразить дисплей
func (d *forecastDisplay) Display(writer io.Writer) error {
	text := "Forecast:\n"
	text += fmt.Sprintf("\tTemperature: %.1f\n", d.temperature)
	text += fmt.Sprintf("\tHumidity: %.1f\n", d.humidity)
	text += fmt.Sprintf("\tPressure: %.1f\n", d.pressure)

	_, err := fmt.Fprintln(writer, text)
	return err
}

// Сделать прогноз
func (d *forecastDisplay) makeForecast() {
	var k float64

	k = 0.7 + rand.Float64()*(1.3-0.7)
	d.temperature = k * d.temperature
	k = 0.7 + rand.Float64()*(1.3-0.7)
	d.humidity = k * d.humidity
	k = 0.7 + rand.Float64()*(1.3-0.7)
	d.pressure = k * d.pressure
}

// NewForecastDisplay Создать дисплей "Прогноз"
func NewForecastDisplay() Displayer {
	return &forecastDisplay{}
}
