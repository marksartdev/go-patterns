package observer

import (
	"fmt"
	"math/rand"

	"github.com/gonum/floats"
)

// Интерфейс визуального элемента
type displayer interface {
	Display() string
}

// DisplayElement Интерфейс визуального элемента, способного быть наблюдателем
type DisplayElement interface {
	displayer
	observer
}

// Визуальный элемент текущего состояния
type currentConditionsDisplay struct {
	temperature float64
	humidity    float64
	weatherData subject
}

// Update Обновить данные
func (c *currentConditionsDisplay) Update(data *measurements) string {
	c.temperature = data.temperature
	c.humidity = data.humidity

	return c.Display()
}

// Display Вывести информацию
func (c *currentConditionsDisplay) Display() string {
	text := "Current conditions:\n"
	text += fmt.Sprintf("\tTemperature: %.1f\n", c.temperature)
	text += fmt.Sprintf("\tHumidity: %.1f\n", c.humidity)

	return text
}

// NewCurrentConditionsDisplay Создать визуальный элемент текущего состояния
func NewCurrentConditionsDisplay(weatherData subject) DisplayElement {
	display := new(currentConditionsDisplay)
	display.weatherData = weatherData

	display.weatherData.RegisterObserver(display)

	return display
}

// Визуальный элемент статистики
type statisticsDisplay struct {
	temperature []float64
	humidity    []float64
	pressure    []float64
	weatherData subject
}

// Update Обновить данные
func (s *statisticsDisplay) Update(data *measurements) string {
	s.temperature = append(s.temperature, data.temperature)
	s.humidity = append(s.humidity, data.humidity)
	s.pressure = append(s.pressure, data.pressure)

	return s.Display()
}

// Display Вывести информацию
func (s *statisticsDisplay) Display() string {
	var temperatureMax, temperatureMin, temperatureAvg float64
	var humidityMax, humidityMin, humidityAvg float64
	var pressureMax, pressureMin, pressureAvg float64

	if len(s.temperature) > 0 {
		temperatureMax = floats.Max(s.temperature)
		temperatureMin = floats.Min(s.temperature)
		temperatureAvg = floats.Sum(s.temperature) / float64(len(s.temperature))
	}

	if len(s.humidity) > 0 {
		humidityMax = floats.Max(s.humidity)
		humidityMin = floats.Min(s.humidity)
		humidityAvg = floats.Sum(s.humidity) / float64(len(s.humidity))
	}

	if len(s.pressure) > 0 {
		pressureMax = floats.Max(s.pressure)
		pressureMin = floats.Min(s.pressure)
		pressureAvg = floats.Sum(s.pressure) / float64(len(s.pressure))
	}

	text := "Statistics:\n"
	text += fmt.Sprintf("\tTemperature (max/min/avg): %.1f/%.1f/%.1f\n", temperatureMax, temperatureMin, temperatureAvg)
	text += fmt.Sprintf("\tHumidity (max/min/avg): %.1f/%.1f/%.1f\n", humidityMax, humidityMin, humidityAvg)
	text += fmt.Sprintf("\tPressure (max/min/avg): %.1f/%.1f/%.1f\n", pressureMax, pressureMin, pressureAvg)

	return text
}

// NewStatisticsDisplay Создать визуальный элемент статистики
func NewStatisticsDisplay(weatherData subject) DisplayElement {
	display := new(statisticsDisplay)
	display.weatherData = weatherData

	display.weatherData.RegisterObserver(display)

	return display
}

// Визуальный элемент прогноза
type forecastDisplay struct {
	temperature float64
	humidity    float64
	pressure    float64
	weatherData subject
}

// Update Обновить данные
func (f *forecastDisplay) Update(data *measurements) string {
	f.temperature = data.temperature
	f.humidity = data.humidity
	f.pressure = data.pressure

	f.makeForecast()

	return f.Display()
}

// Display Вывести информацию
func (f *forecastDisplay) Display() string {
	text := "Forecast:\n"
	text += fmt.Sprintf("\tTemperature: %.1f\n", f.temperature)
	text += fmt.Sprintf("\tHumidity: %.1f\n", f.humidity)
	text += fmt.Sprintf("\tPressure: %.1f\n", f.pressure)

	return text
}

// Сделать прогноз
func (f *forecastDisplay) makeForecast() {
	rand.Seed(512)

	f.temperature = f.getCoefficient() * f.temperature
	f.humidity = f.getCoefficient() * f.humidity
	f.pressure = f.getCoefficient() * f.pressure
}

// Получить коэффициент для прогноза
func (f *forecastDisplay) getCoefficient() float64 {
	return 0.7 + rand.Float64()*(1.3-0.7)
}

// NewForecastDisplay Создать визуальный элемент прогноза
func NewForecastDisplay(weatherData subject) DisplayElement {
	display := new(forecastDisplay)
	display.weatherData = weatherData

	display.weatherData.RegisterObserver(display)

	return display
}
