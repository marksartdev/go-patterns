package observer

import (
	"bytes"
	"testing"
)

var errStringF = "Некорректный ответ (ожидалось %f, получено %f)"
var errStringS = "Некорректный ответ (ожидалось %s, получено %s)"

func TestWeatherData_SetTemperature(t *testing.T) {
	reader := bytes.NewReader([]byte("1.1"))

	wd := NewWeatherData()
	wd.SetTemperature(reader)

	if wd.getTemperature() != 1.1 {
		t.Errorf(errStringF, 1.1, wd.getTemperature())
	}
}

func TestWeatherData_SetHumidity(t *testing.T) {
	reader := bytes.NewReader([]byte("2.2"))

	wd := NewWeatherData()
	wd.SetHumidity(reader)

	if wd.getHumidity() != 2.2 {
		t.Errorf(errStringF, 2.2, wd.getHumidity())
	}
}

func TestWeatherData_SetPressure(t *testing.T) {
	reader := bytes.NewReader([]byte("3.3"))

	wd := NewWeatherData()
	wd.SetPressure(reader)

	if wd.getPressure() != 3.3 {
		t.Errorf(errStringF, 3.3, wd.getPressure())
	}
}

func TestNewCurrentConditionDisplay(t *testing.T) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	expected := "Current condition:\n" +
		"\tTemperature: 0.0\n" +
		"\tHumidity: 0.0\n" +
		"\tPressure: 0.0\n\n"

	display := NewCurrentConditionDisplay()
	err := display.Display(buffer)
	if err != nil {
		t.Error(err)
	}

	if buffer.String() != expected {
		t.Errorf(errStringS, expected, buffer.String())
	}
}
