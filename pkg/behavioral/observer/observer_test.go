package observer_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/marksartdev/go-patterns/pkg/behavioral/observer"
)

const (
	currentConditionsPhrase = "Current conditions:\n"
	errStringF              = "Некорректный результат. Ожидалось %f, получено %f."
	errStringS              = "Некорректный результат. Ожидалось %s, получено %s."
)

func TestNewWeatherData(t *testing.T) {
	temperature, humidity, pressure := 20.0, 60.0, 600.0

	wd := observer.NewWeatherData()
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
	wd := observer.NewWeatherData()

	data := new(observer.Measurements)
	data.Temperature = 25.0
	data.Humidity = 65.0

	expected := currentConditionsPhrase
	expected += fmt.Sprintf("\tTemperature: %.1f\n", data.Temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", data.Humidity)

	display := observer.NewCurrentConditionsDisplay(wd)

	result := display.Update(wd, data)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestNewStatisticsDisplay(t *testing.T) {
	wd := observer.NewWeatherData()

	dataMin := new(observer.Measurements)
	dataMin.Temperature = 27.8
	dataMin.Humidity = 70.0
	dataMin.Pressure = 700.0

	dataMax := new(observer.Measurements)
	dataMax.Temperature = 32.2
	dataMax.Humidity = 80.0
	dataMax.Pressure = 800.0

	dataAvg := new(observer.Measurements)
	dataAvg.Temperature = (dataMin.Temperature + dataMax.Temperature) / 2
	dataAvg.Humidity = (dataMin.Humidity + dataMax.Humidity) / 2
	dataAvg.Pressure = (dataMin.Pressure + dataMax.Pressure) / 2

	expected := "Statistics:\n"
	expected += fmt.Sprintf(
		"\tTemperature (min/max/avg): %.1f/%.1f/%.1f\n",
		dataMin.Temperature,
		dataMax.Temperature,
		dataAvg.Temperature,
	)
	expected += fmt.Sprintf(
		"\tHumidity (min/max/avg): %.1f/%.1f/%.1f\n",
		dataMin.Humidity,
		dataMax.Humidity,
		dataAvg.Humidity,
	)
	expected += fmt.Sprintf(
		"\tPressure (min/max/avg): %.1f/%.1f/%.1f\n",
		dataMin.Pressure,
		dataMax.Pressure,
		dataAvg.Pressure,
	)

	display := observer.NewStatisticsDisplay(wd)

	_ = display.Update(wd, dataMin)
	result := display.Update(wd, dataMax)

	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestNewForecastDisplay(t *testing.T) {
	wd := observer.NewWeatherData()

	data := new(observer.Measurements)
	data.Temperature = 10.0
	data.Humidity = 40.0
	data.Pressure = 550.0

	rand.Seed(512)

	expected := "Forecast:\n"
	expected += fmt.Sprintf("\tTemperature: %.1f\n", getCoefficient()*data.Temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", getCoefficient()*data.Humidity)
	expected += fmt.Sprintf("\tPressure: %.1f\n", getCoefficient()*data.Pressure)

	display := observer.NewForecastDisplay(wd)

	result := display.Update(wd, data)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func getCoefficient() float64 {
	// nolint:gosec // Example
	return 0.7 + rand.Float64()*(1.3-0.7)
}

func TestNewHeatIndexDisplay_Above27C(t *testing.T) {
	wd := observer.NewWeatherData()

	data := new(observer.Measurements)
	data.Temperature = 30.0
	data.Humidity = 50.0

	expected := "Heat index: 31.0\n"

	display := observer.NewHeatIndexDisplay(wd)

	result := display.Update(wd, data)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestNewHeatIndexDisplay_Below27C(t *testing.T) {
	wd := observer.NewWeatherData()

	data := new(observer.Measurements)
	data.Temperature = 20.0
	data.Humidity = 50.0

	expected := ""

	display := observer.NewHeatIndexDisplay(wd)

	result := display.Update(wd, data)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_NotifyObservers_ActiveDelivery(t *testing.T) {
	data := new(observer.Measurements)
	data.Temperature = 35.0
	data.Humidity = 80.0

	expected := currentConditionsPhrase
	expected += fmt.Sprintf("\tTemperature: %.1f\n", data.Temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", data.Humidity)

	wd := observer.NewWeatherData()
	_ = observer.NewCurrentConditionsDisplay(wd)

	wd.SetChanged()

	result := wd.NotifyObservers(data)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_NotifyObservers_CurrentConditionsDisplay(t *testing.T) {
	temperature, humidity, pressure := 24.0, 62.0, 630.0

	expected := currentConditionsPhrase
	expected += fmt.Sprintf("\tTemperature: %.1f\n", temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", humidity)

	wd := observer.NewWeatherData()
	_ = observer.NewCurrentConditionsDisplay(wd)

	wd.SetChanged()

	result := wd.SetMeasurements(temperature, humidity, pressure)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_NotifyObservers_StatisticsDisplay(t *testing.T) {
	dataMin := new(observer.Measurements)
	dataMin.Temperature = 26.8
	dataMin.Humidity = 65.0
	dataMin.Pressure = 650.0

	dataMax := new(observer.Measurements)
	dataMax.Temperature = 32.2
	dataMax.Humidity = 77.0
	dataMax.Pressure = 770.0

	dataAvg := new(observer.Measurements)
	dataAvg.Temperature = (dataMin.Temperature + dataMax.Temperature) / 2
	dataAvg.Humidity = (dataMin.Humidity + dataMax.Humidity) / 2
	dataAvg.Pressure = (dataMin.Pressure + dataMax.Pressure) / 2

	expected := "Statistics:\n"
	expected += fmt.Sprintf(
		"\tTemperature (min/max/avg): %.1f/%.1f/%.1f\n",
		dataMin.Temperature,
		dataMax.Temperature,
		dataAvg.Temperature,
	)
	expected += fmt.Sprintf(
		"\tHumidity (min/max/avg): %.1f/%.1f/%.1f\n",
		dataMin.Humidity,
		dataMax.Humidity,
		dataAvg.Humidity,
	)
	expected += fmt.Sprintf(
		"\tPressure (min/max/avg): %.1f/%.1f/%.1f\n",
		dataMin.Pressure,
		dataMax.Pressure,
		dataAvg.Pressure,
	)

	wd := observer.NewWeatherData()
	_ = observer.NewStatisticsDisplay(wd)

	wd.SetChanged()
	_ = wd.SetMeasurements(dataMin.Temperature, dataMin.Humidity, dataMin.Pressure)

	wd.SetChanged()

	result := wd.SetMeasurements(dataMax.Temperature, dataMax.Humidity, dataMax.Pressure)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_NotifyObservers_ForecastDisplay(t *testing.T) {
	temperature, humidity, pressure := 22.0, 41.0, 670.0

	rand.Seed(observer.RandSeed)

	expected := "Forecast:\n"
	expected += fmt.Sprintf("\tTemperature: %.1f\n", getCoefficient()*temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", getCoefficient()*humidity)
	expected += fmt.Sprintf("\tPressure: %.1f\n", getCoefficient()*pressure)

	wd := observer.NewWeatherData()
	_ = observer.NewForecastDisplay(wd)

	wd.SetChanged()

	result := wd.SetMeasurements(temperature, humidity, pressure)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_NotifyObservers_HeatIndexDisplay(t *testing.T) {
	temperature, humidity, pressure := 35.0, 60.0, 720.0

	expected := "Heat index: 45.1\n"

	wd := observer.NewWeatherData()
	_ = observer.NewHeatIndexDisplay(wd)

	wd.SetChanged()

	result := wd.SetMeasurements(temperature, humidity, pressure)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_NotifyObservers_WithoutSetChanged(t *testing.T) {
	temperature, humidity, pressure := 15.0, 10.0, 590.0

	expected := ""

	wd := observer.NewWeatherData()
	_ = observer.NewCurrentConditionsDisplay(wd)

	result := wd.SetMeasurements(temperature, humidity, pressure)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}

func TestWeatherData_RemoveObserver(t *testing.T) {
	temperature, humidity, pressure := 10.1, 35.0, 630.0
	expected := ""

	wd := observer.NewWeatherData()
	display := observer.NewCurrentConditionsDisplay(wd)

	wd.RemoveObserver(display)

	result := wd.SetMeasurements(temperature, humidity, pressure)
	if result != expected {
		t.Errorf(errStringS, expected, result)
	}
}
