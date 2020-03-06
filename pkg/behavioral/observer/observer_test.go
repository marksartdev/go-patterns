package observer

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"testing"
)

var errStringF = "Некорректный ответ (ожидалось %f, получено %f)"
var errStringS = "Некорректный ответ (ожидалось %s, получено %s)"

// Битый reader
type badReader struct {
}

func (r *badReader) Read(_ []byte) (n int, err error) {
	return 0, os.ErrInvalid
}

// Битый writer
type badWriter struct {
}

// Write Записать в writer
func (w *badWriter) Write(_ []byte) (n int, err error) {
	return 0, os.ErrInvalid
}

func TestWeatherData_SetTemperature(t *testing.T) {
	reader := bytes.NewReader([]byte("1.1"))
	buffer := bytes.NewBuffer(make([]byte, 0))
	expected := "Введите температуру\n"

	wd := NewWeatherData()
	wd.SetInput(reader)
	wd.SetOutput(buffer)

	if err := wd.SetTemperature(); err != nil {
		t.Error(err)
	}

	if buffer.String() != expected {
		t.Errorf(errStringS, expected, buffer)
	}

	if wd.getTemperature() != 1.1 {
		t.Errorf(errStringF, 1.1, wd.getTemperature())
	}
}

func TestWeatherData_SetHumidity(t *testing.T) {
	reader := bytes.NewReader([]byte("2.2"))
	buffer := bytes.NewBuffer(make([]byte, 0))
	expected := "Введите влажность\n"

	wd := NewWeatherData()
	wd.SetInput(reader)
	wd.SetOutput(buffer)

	if err := wd.SetHumidity(); err != nil {
		t.Error(err)
	}

	if buffer.String() != expected {
		t.Errorf(errStringS, expected, buffer)
	}

	if wd.getHumidity() != 2.2 {
		t.Errorf(errStringF, 2.2, wd.getHumidity())
	}
}

func TestWeatherData_SetPressure(t *testing.T) {
	reader := bytes.NewReader([]byte("3.3"))
	buffer := bytes.NewBuffer(make([]byte, 0))
	expected := "Введите давление\n"

	wd := NewWeatherData()
	wd.SetInput(reader)
	wd.SetOutput(buffer)

	if err := wd.SetPressure(); err != nil {
		t.Error(err)
	}

	if buffer.String() != expected {
		t.Errorf(errStringS, expected, buffer)
	}

	if wd.getPressure() != 3.3 {
		t.Errorf(errStringF, 3.3, wd.getPressure())
	}
}

func TestWeatherData_InputNaN(t *testing.T) {
	reader := bytes.NewReader([]byte("NotNumber"))
	buffer := bytes.NewBuffer(make([]byte, 0))

	wd := NewWeatherData()
	wd.SetInput(reader)
	wd.SetOutput(buffer)

	err := wd.SetTemperature()
	if err == nil {
		t.Error("Ожидалась ошибка при вводе строки вместо числа")
	} else if err.Error() != "strconv.ParseFloat: parsing \"NotNumber\": invalid syntax" {
		t.Error(err)
	}
}

func TestWeatherData_ReadErr(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0))

	wd := NewWeatherData()
	wd.SetInput(new(badReader))
	wd.SetOutput(buffer)

	err := wd.SetHumidity()
	if err == nil {
		t.Error("Ожидалась ошибка при использовании битого Reader")
	} else if err != os.ErrInvalid {
		t.Error(err)
	}
}

func TestWeatherData_WriteErr(t *testing.T) {
	wd := NewWeatherData()
	wd.SetOutput(new(badWriter))

	err := wd.SetPressure()
	if err == nil {
		t.Error("Ожидалась ошибка при использовании битого Writer")
	} else if err != os.ErrInvalid {
		t.Error(err)
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
	display.SetOutput(buffer)

	display.Update(data)

	err := display.Display()
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
	display.SetOutput(buffer)

	data := new(measurements)
	data.temperature = 20.0
	data.humidity = 50.0
	data.pressure = 600.0
	display.Update(data)

	data.temperature = 30.0
	data.humidity = 70.0
	data.pressure = 700.0
	display.Update(data)

	err := display.Display()
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
	display.SetOutput(buffer)

	display.Update(data)

	err := display.Display()
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

func TestCurrentConditionsDisplay_DisplayErr(t *testing.T) {
	display := NewCurrentConditionsDisplay()
	display.SetOutput(new(badWriter))

	err := display.Display()
	if err == nil {
		t.Error("Ожидалась ошибка при отображении текущего состояния")
	} else if err != os.ErrInvalid {
		t.Error(err)
	}
}

func TestStatisticsDisplay_DisplayErr(t *testing.T) {
	display := NewStatisticsDisplay()
	display.SetOutput(new(badWriter))

	err := display.Display()
	if err == nil {
		t.Error("Ожидалась ошибка при отображении статистики")
	} else if err != os.ErrInvalid {
		t.Error(err)
	}
}

func TestForecastDisplay_DisplayErr(t *testing.T) {
	display := NewForecastDisplay()
	display.SetOutput(new(badWriter))

	err := display.Display()
	if err == nil {
		t.Error("Ожидалась ошибка при отображении прогноза")
	} else if err != os.ErrInvalid {
		t.Error(err)
	}
}
