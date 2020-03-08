package observer

import (
	"fmt"
	"math/rand"
	"testing"
)

var errStringF = "Некорректный результат (ожидалось %.1f, получено %.1f)"
var errStringS = "Некорректный результат (ожидалось %s, получено %s)"

func TestNewWeatherData(t *testing.T) {
	temperature, humidity, pressure := 1.1, 1.2, 1.3

	wd := NewWeatherData()
	wd.SetMeasurements(temperature, humidity, pressure)

	if wd.GetTemperature() != temperature {
		t.Errorf(errStringF, temperature, wd.GetTemperature())
	}

	if wd.GetHumidity() != humidity {
		t.Errorf(errStringF, humidity, wd.GetHumidity())
	}

	if wd.GetPressure() != pressure {
		t.Errorf(errStringF, pressure, wd.GetPressure())
	}
}

func TestNewCurrentConditionsDisplay(t *testing.T) {
	wd := NewWeatherData()

	data := new(measurements)
	data.temperature = 20.5
	data.humidity = 91.0

	expected := "Current conditions:\n"
	expected += fmt.Sprintf("\tTemperature: %.1f\n", data.temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", data.humidity)

	display := NewCurrentConditionsDisplay(wd)

	result := display.Update(data)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestNewStatisticsDisplay(t *testing.T) {
	wd := NewWeatherData()

	expected := "Statistics:\n"
	expected += fmt.Sprintf("\tTemperature (max/min/avg): %.1f/%.1f/%.1f\n", 30.0, 20.0, 25.0)
	expected += fmt.Sprintf("\tHumidity (max/min/avg): %.1f/%.1f/%.1f\n", 70.0, 50.0, 60.0)
	expected += fmt.Sprintf("\tPressure (max/min/avg): %.1f/%.1f/%.1f\n", 700.0, 600.0, 650.0)

	display := NewStatisticsDisplay(wd)

	data := new(measurements)
	data.temperature = 20.0
	data.humidity = 50.0
	data.pressure = 600.0
	result := display.Update(data)

	data.temperature = 30.0
	data.humidity = 70.0
	data.pressure = 700.0
	result = display.Update(data)

	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestNewForecastDisplay(t *testing.T) {
	wd := NewWeatherData()

	data := new(measurements)
	data.temperature = 20.0
	data.humidity = 60.0
	data.pressure = 650.0

	rand.Seed(512)

	expected := "Forecast:\n"
	expected += fmt.Sprintf("\tTemperature: %.1f\n", getCoefficient()*data.temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", getCoefficient()*data.humidity)
	expected += fmt.Sprintf("\tPressure: %.1f\n", getCoefficient()*data.pressure)

	display := NewForecastDisplay(wd)

	result := display.Update(data)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func getCoefficient() float64 {
	return 0.7 + rand.Float64()*(1.3-0.7)
}

func TestWeatherData_NotifyObservers(t *testing.T) {
	temperature, humidity, pressure := 2.1, 2.2, 2.3

	expected := "Current conditions:\n"
	expected += fmt.Sprintf("\tTemperature: %.1f\n", temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", humidity)

	wd := NewWeatherData()
	_ = NewCurrentConditionsDisplay(wd)

	result := wd.SetMeasurements(temperature, humidity, pressure)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_RemoveObserver(t *testing.T) {
	temperature, humidity, pressure := 2.1, 2.2, 2.3
	expected := ""

	wd := NewWeatherData()
	display := NewCurrentConditionsDisplay(wd)

	wd.RemoveObserver(display)

	result := wd.SetMeasurements(temperature, humidity, pressure)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}
