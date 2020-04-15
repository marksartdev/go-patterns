package observer

import (
	"fmt"
	"math/rand"
	"testing"
)

var errStringF = "Некорректный результат. Ожидалось %f, получено %f."
var errStringS = "Некорректный результат. Ожидалось %s, получено %s."

func TestNewWeatherData(t *testing.T) {
	temperature, humidity, pressure := 20.0, 60.0, 600.0

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
	data.temperature = 25.0
	data.humidity = 65.0

	expected := "Current conditions:\n"
	expected += fmt.Sprintf("\tTemperature: %.1f\n", data.temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", data.humidity)

	display := NewCurrentConditionsDisplay(wd)

	result := display.Update(wd, data)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestNewStatisticsDisplay(t *testing.T) {
	wd := NewWeatherData()

	dataMin := new(measurements)
	dataMin.temperature = 27.8
	dataMin.humidity = 70.0
	dataMin.pressure = 700.0

	dataMax := new(measurements)
	dataMax.temperature = 32.2
	dataMax.humidity = 80.0
	dataMax.pressure = 800.0

	dataAvg := new(measurements)
	dataAvg.temperature = (dataMin.temperature + dataMax.temperature) / 2
	dataAvg.humidity = (dataMin.humidity + dataMax.humidity) / 2
	dataAvg.pressure = (dataMin.pressure + dataMax.pressure) / 2

	expected := "Statistics:\n"
	expected += fmt.Sprintf(
		"\tTemperature (min/max/avg): %.1f/%.1f/%.1f\n",
		dataMin.temperature,
		dataMax.temperature,
		dataAvg.temperature,
	)
	expected += fmt.Sprintf(
		"\tHumidity (min/max/avg): %.1f/%.1f/%.1f\n",
		dataMin.humidity,
		dataMax.humidity,
		dataAvg.humidity,
	)
	expected += fmt.Sprintf(
		"\tPressure (min/max/avg): %.1f/%.1f/%.1f\n",
		dataMin.pressure,
		dataMax.pressure,
		dataAvg.pressure,
	)

	display := NewStatisticsDisplay(wd)

	_ = display.Update(wd, dataMin)
	result := display.Update(wd, dataMax)

	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestNewForecastDisplay(t *testing.T) {
	wd := NewWeatherData()

	data := new(measurements)
	data.temperature = 10.0
	data.humidity = 40.0
	data.pressure = 550.0

	rand.Seed(512)

	expected := "Forecast:\n"
	expected += fmt.Sprintf("\tTemperature: %.1f\n", getCoefficient()*data.temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", getCoefficient()*data.humidity)
	expected += fmt.Sprintf("\tPressure: %.1f\n", getCoefficient()*data.pressure)

	display := NewForecastDisplay(wd)

	result := display.Update(wd, data)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func getCoefficient() float64 {
	return 0.7 + rand.Float64()*(1.3-0.7)
}

func TestNewHeatIndexDisplay_Above27C(t *testing.T) {
	wd := NewWeatherData()

	data := new(measurements)
	data.temperature = 30.0
	data.humidity = 50.0

	expected := "Heat index: 31.0\n"

	display := NewHeatIndexDisplay(wd)

	result := display.Update(wd, data)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestNewHeatIndexDisplay_Below27C(t *testing.T) {
	wd := NewWeatherData()

	data := new(measurements)
	data.temperature = 20.0
	data.humidity = 50.0

	expected := ""

	display := NewHeatIndexDisplay(wd)

	result := display.Update(wd, data)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_NotifyObservers_ActiveDelivery(t *testing.T) {
	data := new(measurements)
	data.temperature = 35.0
	data.humidity = 80.0

	expected := "Current conditions:\n"
	expected += fmt.Sprintf("\tTemperature: %.1f\n", data.temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", data.humidity)

	wd := NewWeatherData()
	_ = NewCurrentConditionsDisplay(wd)

	wd.SetChanged()
	result := wd.NotifyObservers(data)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_NotifyObservers_CurrentConditionsDisplay(t *testing.T) {
	temperature, humidity, pressure := 24.0, 62.0, 630.0

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
	dataMin := new(measurements)
	dataMin.temperature = 26.8
	dataMin.humidity = 65.0
	dataMin.pressure = 650.0

	dataMax := new(measurements)
	dataMax.temperature = 32.2
	dataMax.humidity = 77.0
	dataMax.pressure = 770.0

	dataAvg := new(measurements)
	dataAvg.temperature = (dataMin.temperature + dataMax.temperature) / 2
	dataAvg.humidity = (dataMin.humidity + dataMax.humidity) / 2
	dataAvg.pressure = (dataMin.pressure + dataMax.pressure) / 2

	expected := "Statistics:\n"
	expected += fmt.Sprintf(
		"\tTemperature (min/max/avg): %.1f/%.1f/%.1f\n",
		dataMin.temperature,
		dataMax.temperature,
		dataAvg.temperature,
	)
	expected += fmt.Sprintf(
		"\tHumidity (min/max/avg): %.1f/%.1f/%.1f\n",
		dataMin.humidity,
		dataMax.humidity,
		dataAvg.humidity,
	)
	expected += fmt.Sprintf(
		"\tPressure (min/max/avg): %.1f/%.1f/%.1f\n",
		dataMin.pressure,
		dataMax.pressure,
		dataAvg.pressure,
	)

	wd := NewWeatherData()
	_ = NewStatisticsDisplay(wd)

	wd.SetChanged()
	_ = wd.SetMeasurements(dataMin.temperature, dataMin.humidity, dataMin.pressure)

	wd.SetChanged()
	result := wd.SetMeasurements(dataMax.temperature, dataMax.humidity, dataMax.pressure)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_NotifyObservers_ForecastDisplay(t *testing.T) {
	temperature, humidity, pressure := 22.0, 41.0, 670.0
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

func TestWeatherData_NotifyObservers_HeatIndexDisplay(t *testing.T) {
	temperature, humidity, pressure := 35.0, 60.0, 720.0

	expected := "Heat index: 45.1\n"

	wd := NewWeatherData()
	_ = NewHeatIndexDisplay(wd)

	wd.SetChanged()
	result := wd.SetMeasurements(temperature, humidity, pressure)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_NotifyObservers_WithoutSetChanged(t *testing.T) {
	temperature, humidity, pressure := 15.0, 10.0, 590.0

	expected := ""

	wd := NewWeatherData()
	_ = NewCurrentConditionsDisplay(wd)

	result := wd.SetMeasurements(temperature, humidity, pressure)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_RemoveObserver(t *testing.T) {
	temperature, humidity, pressure := 10.1, 35.0, 630.0
	expected := ""

	wd := NewWeatherData()
	display := NewCurrentConditionsDisplay(wd)

	wd.RemoveObserver(display)

	result := wd.SetMeasurements(temperature, humidity, pressure)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}
