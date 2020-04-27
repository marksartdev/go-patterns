package observer

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/gonum/floats"
)

const (
	// RandSeed Индекс для генератора случайных чисел
	RandSeed                = 512
	minHeatIndexTemperature = 27
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

// Update Обновить данные.
func (c *currentConditionsDisplay) Update(observable subject, data *Measurements) string {
	if data != nil {
		c.temperature = data.Temperature
		c.humidity = data.Humidity
	} else {
		wd, ok := observable.(WeatherDater)
		if ok {
			c.temperature = wd.GetTemperature()
			c.humidity = wd.GetHumidity()
		}
	}

	return c.Display()
}

// Display Вывести информацию.
func (c *currentConditionsDisplay) Display() string {
	text := "Current conditions:\n"
	text += fmt.Sprintf("\tTemperature: %.1f\n", c.temperature)
	text += fmt.Sprintf("\tHumidity: %.1f\n", c.humidity)

	return text
}

// NewCurrentConditionsDisplay Создать визуальный элемент текущего состояния.
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

// Update Обновить данные.
func (s *statisticsDisplay) Update(observable subject, data *Measurements) string {
	if data != nil {
		s.temperature = append(s.temperature, data.Temperature)
		s.humidity = append(s.humidity, data.Humidity)
		s.pressure = append(s.pressure, data.Pressure)
	} else {
		wd, ok := observable.(WeatherDater)
		if ok {
			s.temperature = append(s.temperature, wd.GetTemperature())
			s.humidity = append(s.humidity, wd.GetHumidity())
			s.pressure = append(s.pressure, wd.GetPressure())
		}
	}

	return s.Display()
}

// Display Вывести информацию.
func (s *statisticsDisplay) Display() string {
	var (
		temperatureMax, temperatureMin, temperatureAvg float64
		humidityMax, humidityMin, humidityAvg          float64
		pressureMax, pressureMin, pressureAvg          float64
	)

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
	text += fmt.Sprintf("\tTemperature (min/max/avg): %.1f/%.1f/%.1f\n", temperatureMin, temperatureMax, temperatureAvg)
	text += fmt.Sprintf("\tHumidity (min/max/avg): %.1f/%.1f/%.1f\n", humidityMin, humidityMax, humidityAvg)
	text += fmt.Sprintf("\tPressure (min/max/avg): %.1f/%.1f/%.1f\n", pressureMin, pressureMax, pressureAvg)

	return text
}

// NewStatisticsDisplay Создать визуальный элемент статистики.
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

// Update Обновить данные.
func (f *forecastDisplay) Update(observable subject, data *Measurements) string {
	if data != nil {
		f.temperature = data.Temperature
		f.humidity = data.Humidity
		f.pressure = data.Pressure
	} else {
		wd, ok := observable.(WeatherDater)
		if ok {
			f.temperature = wd.GetTemperature()
			f.humidity = wd.GetHumidity()
			f.pressure = wd.GetPressure()
		}
	}

	f.makeForecast()

	return f.Display()
}

// Display Вывести информацию.
func (f *forecastDisplay) Display() string {
	text := "Forecast:\n"
	text += fmt.Sprintf("\tTemperature: %.1f\n", f.temperature)
	text += fmt.Sprintf("\tHumidity: %.1f\n", f.humidity)
	text += fmt.Sprintf("\tPressure: %.1f\n", f.pressure)

	return text
}

// Сделать прогноз.
func (f *forecastDisplay) makeForecast() {
	rand.Seed(RandSeed)

	f.temperature = f.getCoefficient() * f.temperature
	f.humidity = f.getCoefficient() * f.humidity
	f.pressure = f.getCoefficient() * f.pressure
}

// Получить коэффициент для прогноза.
func (f *forecastDisplay) getCoefficient() float64 {
	return 0.7 + rand.Float64()*(1.3-0.7)
}

// NewForecastDisplay Создать визуальный элемент прогноза.
func NewForecastDisplay(weatherData subject) DisplayElement {
	display := new(forecastDisplay)
	display.weatherData = weatherData

	display.weatherData.RegisterObserver(display)

	return display
}

// Визуальный элемент теплового индекса
type heatIndexDisplay struct {
	temperature float64
	humidity    float64
	weatherData subject
}

// Update Обновить данные.
func (h *heatIndexDisplay) Update(observable subject, data *Measurements) string {
	if data != nil {
		h.temperature = data.Temperature
		h.humidity = data.Humidity
	} else {
		wd, ok := observable.(WeatherDater)
		if ok {
			h.temperature = wd.GetTemperature()
			h.humidity = wd.GetHumidity()
		}
	}

	return h.Display()
}

// Display Вывести информацию.
func (h *heatIndexDisplay) Display() string {
	if h.temperature < minHeatIndexTemperature {
		return ""
	}

	c1 := -8.78469475556
	c2 := 1.61139411
	c3 := 2.33854883889
	c4 := -0.14611605
	c5 := -0.012308094
	c6 := -0.0164248277778
	c7 := 0.002211732
	c8 := 0.00072546
	c9 := -0.000003582

	heatIndex := c1 +
		c2*h.temperature +
		c3*h.humidity +
		c4*h.temperature*h.humidity +
		c5*math.Pow(h.temperature, 2) +
		c6*math.Pow(h.humidity, 2) +
		c7*math.Pow(h.temperature, 2)*h.humidity +
		c8*h.temperature*math.Pow(h.humidity, 2) +
		c9*math.Pow(h.temperature, 2)*math.Pow(h.humidity, 2)

	return fmt.Sprintf("Heat index: %.1f\n", heatIndex)
}

// NewHeatIndexDisplay Создать визуальный элемент теплового индекса.
func NewHeatIndexDisplay(weatherData subject) DisplayElement {
	display := new(heatIndexDisplay)
	display.weatherData = weatherData

	display.weatherData.RegisterObserver(display)

	return display
}
