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
	data.temperature = 2.1
	data.humidity = 2.2

	expected := "Current conditions:\n"
	expected += fmt.Sprintf("\tTemperature: %.1f\n", data.temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", data.humidity)

	display := NewCurrentConditionsDisplay(wd)

	result := display.Update(nil, data)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestNewStatisticsDisplay(t *testing.T) {
	wd := NewWeatherData()

	expected := "Statistics:\n"
	expected += fmt.Sprintf("\tTemperature (max/min/avg): %.1f/%.1f/%.1f\n", 3.7, 3.1, 3.4)
	expected += fmt.Sprintf("\tHumidity (max/min/avg): %.1f/%.1f/%.1f\n", 3.8, 3.2, 3.5)
	expected += fmt.Sprintf("\tPressure (max/min/avg): %.1f/%.1f/%.1f\n", 3.9, 3.3, 3.6)

	display := NewStatisticsDisplay(wd)

	data := new(measurements)
	data.temperature = 3.1
	data.humidity = 3.2
	data.pressure = 3.3
	result := display.Update(nil, data)

	data.temperature = 3.7
	data.humidity = 3.8
	data.pressure = 3.9
	result = display.Update(nil, data)

	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestNewForecastDisplay(t *testing.T) {
	wd := NewWeatherData()

	data := new(measurements)
	data.temperature = 4.1
	data.humidity = 4.2
	data.pressure = 4.3

	rand.Seed(512)

	expected := "Forecast:\n"
	expected += fmt.Sprintf("\tTemperature: %.1f\n", getCoefficient()*data.temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", getCoefficient()*data.humidity)
	expected += fmt.Sprintf("\tPressure: %.1f\n", getCoefficient()*data.pressure)

	display := NewForecastDisplay(wd)

	result := display.Update(nil, data)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func getCoefficient() float64 {
	return 0.7 + rand.Float64()*(1.3-0.7)
}

func TestWeatherData_NotifyObservers_ActiveDelivery(t *testing.T) {
	newMeasurements := new(measurements)
	newMeasurements.temperature = 5.1
	newMeasurements.humidity = 5.2

	expected := "Current conditions:\n"
	expected += fmt.Sprintf("\tTemperature: %.1f\n", newMeasurements.temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", newMeasurements.humidity)

	wd := NewWeatherData()
	_ = NewCurrentConditionsDisplay(wd)

	wd.SetChanged()
	result := wd.NotifyObservers(newMeasurements)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_NotifyObservers_CurrentConditionsDisplay(t *testing.T) {
	temperature, humidity, pressure := 6.1, 6.2, 6.3

	expected := "Current conditions:\n"
	expected += fmt.Sprintf("\tTemperature: %.1f\n", temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", humidity)

	wd := NewWeatherData()
	_ = NewCurrentConditionsDisplay(wd)

	wd.SetChanged()
	result := wd.SetMeasurements(temperature, humidity, pressure)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_NotifyObservers_StatisticsDisplay(t *testing.T) {
	expected := "Statistics:\n"
	expected += fmt.Sprintf("\tTemperature (max/min/avg): %.1f/%.1f/%.1f\n", 7.7, 7.1, 7.4)
	expected += fmt.Sprintf("\tHumidity (max/min/avg): %.1f/%.1f/%.1f\n", 7.8, 7.2, 7.5)
	expected += fmt.Sprintf("\tPressure (max/min/avg): %.1f/%.1f/%.1f\n", 7.9, 7.3, 7.6)

	wd := NewWeatherData()
	_ = NewStatisticsDisplay(wd)

	wd.SetChanged()
	_ = wd.SetMeasurements(7.1, 7.2, 7.3)

	wd.SetChanged()
	result := wd.SetMeasurements(7.7, 7.8, 7.9)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_NotifyObservers_ForecastDisplay(t *testing.T) {
	temperature, humidity, pressure := 8.1, 8.2, 8.3
	rand.Seed(512)

	expected := "Forecast:\n"
	expected += fmt.Sprintf("\tTemperature: %.1f\n", getCoefficient()*temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", getCoefficient()*humidity)
	expected += fmt.Sprintf("\tPressure: %.1f\n", getCoefficient()*pressure)

	wd := NewWeatherData()
	_ = NewForecastDisplay(wd)

	wd.SetChanged()
	result := wd.SetMeasurements(temperature, humidity, pressure)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_NotifyObservers_WithoutSetChanged(t *testing.T) {
	temperature, humidity, pressure := 9.1, 9.2, 9.3

	expected := ""

	wd := NewWeatherData()
	_ = NewCurrentConditionsDisplay(wd)

	result := wd.SetMeasurements(temperature, humidity, pressure)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_RemoveObserver(t *testing.T) {
	temperature, humidity, pressure := 10.1, 10.2, 10.3
	expected := ""

	wd := NewWeatherData()
	display := NewCurrentConditionsDisplay(wd)

	wd.RemoveObserver(display)

	result := wd.SetMeasurements(temperature, humidity, pressure)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}
