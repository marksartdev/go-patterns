package observer

import (
	"bytes"
	"fmt"
	"math/rand"
	"regexp"
	"testing"
)

var errStringF = "Некорректный ответ (ожидалось %f, получено %f)"
var errStringS = "Некорректный ответ (ожидалось %s, получено %s)"

func TestWeatherData_SetTemperature(t *testing.T) {
	reader := bytes.NewReader([]byte("1.1"))

	wd := NewWeatherData()
	wd.SetReader(reader)
	wd.SetTemperature()

	if wd.getTemperature() != 1.1 {
		t.Errorf(errStringF, 1.1, wd.getTemperature())
	}
}

func TestWeatherData_SetHumidity(t *testing.T) {
	reader := bytes.NewReader([]byte("2.2"))

	wd := NewWeatherData()
	wd.SetReader(reader)
	wd.SetHumidity()

	if wd.getHumidity() != 2.2 {
		t.Errorf(errStringF, 2.2, wd.getHumidity())
	}
}

func TestWeatherData_SetPressure(t *testing.T) {
	reader := bytes.NewReader([]byte("3.3"))

	wd := NewWeatherData()
	wd.SetReader(reader)
	wd.SetPressure()

	if wd.getPressure() != 3.3 {
		t.Errorf(errStringF, 3.3, wd.getPressure())
	}
}

func TestWeatherData_InputNaN(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	reader := bytes.NewReader([]byte("NotNumber"))
	var wd *weatherData
	var ok bool

	regErr, err := regexp.Compile(`msg=[A-z]+|msg="[А-я ]+"`)
	if err != nil {
		t.Errorf(err.Error())

		return
	}

	if wd, ok = NewWeatherData().(*weatherData); !ok {
		t.Errorf("Полученная структура не имплементирует интерфейс WeatherDater")

		return
	}

	wd.logger.SetOutput(buffer)
	wd.SetReader(reader)
	wd.SetTemperature()
	matches := regErr.FindAllString(buffer.String(), -1)
	if matches[0] != "msg=\"Введено некорректное значение\"" || matches[1] != "msg=EOF" {
		t.Errorf("Получена неправильная последовательность ошибок")
	}
}

func TestNewCurrentConditionsDisplay(t *testing.T) {
	data := new(measurements)
	data.temperature = 18.5
	data.humidity = 90.0
	data.pressure = 650.0

	buffer := bytes.NewBuffer(make([]byte, 0))
	expected := "Current conditions:\n"
	expected += fmt.Sprintf("\tTemperature: %.1f\n", data.temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", data.humidity)
	expected += fmt.Sprintf("\tPressure: %.1f\n\n", data.pressure)

	display := NewCurrentConditionsDisplay()
	display.Update(data)

	err := display.Display(buffer)
	if err != nil {
		t.Error(err)
	}

	if buffer.String() != expected {
		t.Errorf(errStringS, expected, buffer.String())
	}
}

func TestNewStatisticsDisplay(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	expected := "Statistics:\n"
	expected += fmt.Sprintf("\tTemperature (max/min/avg): %.1f/%.1f/%.1f\n", 30.0, 20.0, 25.0)
	expected += fmt.Sprintf("\tHumidity (max/min/avg): %.1f/%.1f/%.1f\n", 70.0, 50.0, 60.0)
	expected += fmt.Sprintf("\tPressure (max/min/avg): %.1f/%.1f/%.1f\n\n", 700.0, 600.0, 650.0)

	display := NewStatisticsDisplay()

	data := new(measurements)
	data.temperature = 20.0
	data.humidity = 50.0
	data.pressure = 600.0
	display.Update(data)

	data.temperature = 30.0
	data.humidity = 70.0
	data.pressure = 700.0
	display.Update(data)

	err := display.Display(buffer)
	if err != nil {
		t.Error(err)
	}

	if buffer.String() != expected {
		t.Errorf(errStringS, expected, buffer.String())
	}
}

func TestNewForecastDisplay(t *testing.T) {
	data := new(measurements)
	data.temperature = 20.0
	data.humidity = 60.0
	data.pressure = 650.0

	rand.Seed(512)

	buffer := bytes.NewBuffer(make([]byte, 0))
	expected := "Forecast:\n"
	expected += fmt.Sprintf("\tTemperature: %.1f\n", getCoefficient()*data.temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n", getCoefficient()*data.humidity)
	expected += fmt.Sprintf("\tPressure: %.1f\n\n", getCoefficient()*data.pressure)

	display := NewForecastDisplay()
	display.Update(data)

	err := display.Display(buffer)
	if err != nil {
		t.Error(err)
	}

	if buffer.String() != expected {
		t.Errorf(errStringS, expected, buffer.String())
	}
}

func getCoefficient() float64 {
	return 0.7 + rand.Float64()*(1.3-0.7)
}
