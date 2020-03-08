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
	weatherData subject
}

// SetWriter Установить writer
func (c *currentConditionsDisplay) SetWriter(writer io.Writer) {
	c.writer = writer
}

// Update Обновить данные
func (c *currentConditionsDisplay) Update(data *measurements) error {
	c.temperature = data.temperature
	c.humidity = data.humidity

	err := c.Display()

	return err
}

// Display Вывести информацию
func (c *currentConditionsDisplay) Display() error {
	text := "Current conditions:\n"
	text += fmt.Sprintf("\tTemperature: %.1f\n", c.temperature)
	text += fmt.Sprintf("\tHumidity: %.1f\n", c.humidity)

	_, err := fmt.Fprintln(c.writer, text)

	return err
}

// NewCurrentConditionsDisplay Создать экран текущего состояния
func NewCurrentConditionsDisplay(weatherData subject) DisplayElement {
	display := new(currentConditionsDisplay)
	display.writer = os.Stdout
	display.weatherData = weatherData

	display.weatherData.RegisterObserver(display)

	return display
}

// Экран статистики
type statisticsDisplay struct {
	writer      io.Writer
	temperature []float64
	humidity    []float64
	pressure    []float64
	weatherData subject
}

// SetWriter Установить writer
func (s *statisticsDisplay) SetWriter(writer io.Writer) {
	s.writer = writer
}

// Update Обновить данные
func (s *statisticsDisplay) Update(data *measurements) error {
	s.temperature = append(s.temperature, data.temperature)
	s.humidity = append(s.humidity, data.humidity)
	s.pressure = append(s.pressure, data.pressure)

	err := s.Display()

	return err
}

// Display Вывести информацию
func (s *statisticsDisplay) Display() error {
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

	_, err := fmt.Fprintln(s.writer, text)

	return err
}

// NewStatisticsDisplay Создать экран статистики
func NewStatisticsDisplay(weatherData subject) DisplayElement {
	display := new(statisticsDisplay)
	display.writer = os.Stdout
	display.weatherData = weatherData

	display.weatherData.RegisterObserver(display)

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
