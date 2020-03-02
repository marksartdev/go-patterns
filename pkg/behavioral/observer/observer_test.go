package observer

import (
	"bytes"
	"testing"
)

var errString = "Некорректный ответ (ожидалось %f, получено %f)"

func TestWeatherData_SetTemperature(t *testing.T) {
	reader := bytes.NewReader([]byte("1.1"))

	wd := NewWeatherData()
	wd.SetTemperature(reader)

	if wd.GetTemperature() != 1.1 {
		t.Errorf(errString, 1.1, wd.GetTemperature())
	}
}

func TestWeatherData_SetHumidity(t *testing.T) {
	reader := bytes.NewReader([]byte("2.2"))

	wd := NewWeatherData()
	wd.SetHumidity(reader)

	if wd.GetHumidity() != 2.2 {
		t.Errorf(errString, 2.2, wd.GetHumidity())
	}
}

func TestWeatherData_SetPressure(t *testing.T) {
	reader := bytes.NewReader([]byte("3.3"))

	wd := NewWeatherData()
	wd.SetPressure(reader)

	if wd.GetPressure() != 3.3 {
		t.Errorf(errString, 3.3, wd.GetPressure())
	}
}
