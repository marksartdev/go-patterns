package observer

import (
	"fmt"
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

//
//func TestNewForecastDisplay(t *testing.T) {
//	data := new(measurements)
//	data.temperature = 20.0
//	data.humidity = 60.0
//	data.pressure = 650.0
//
//	rand.Seed(512)
//
//	buffer := bytes.NewBuffer(make([]byte, 0))
//	expected := "Forecast:\n"
//	expected += fmt.Sprintf("\tTemperature: %.1f\n", getCoefficient()*data.temperature)
//	expected += fmt.Sprintf("\tHumidity: %.1f\n", getCoefficient()*data.humidity)
//	expected += fmt.Sprintf("\tPressure: %.1f\n\n", getCoefficient()*data.pressure)
//
//	display := NewForecastDisplay()
//	display.SetOutput(buffer)
//
//	display.Update(data)
//
//	err := display.Display()
//	if err != nil {
//		t.Error(err)
//	}
//
//	if buffer.String() != expected {
//		t.Errorf(errStringS, expected, buffer.String())
//	}
//}
//
//func getCoefficient() float64 {
//	return 0.7 + rand.Float64()*(1.3-0.7)
//}
//
//func TestCurrentConditionsDisplay_DisplayErr(t *testing.T) {
//	display := NewCurrentConditionsDisplay()
//	display.SetOutput(new(badWriter))
//
//	err := display.Display()
//	if err == nil {
//		t.Error("Ожидалась ошибка при отображении текущего состояния")
//	} else if err != os.ErrInvalid {
//		t.Error(err)
//	}
//}
//
//func TestStatisticsDisplay_DisplayErr(t *testing.T) {
//	display := NewStatisticsDisplay()
//	display.SetOutput(new(badWriter))
//
//	err := display.Display()
//	if err == nil {
//		t.Error("Ожидалась ошибка при отображении статистики")
//	} else if err != os.ErrInvalid {
//		t.Error(err)
//	}
//}
//
//func TestForecastDisplay_DisplayErr(t *testing.T) {
//	display := NewForecastDisplay()
//	display.SetOutput(new(badWriter))
//
//	err := display.Display()
//	if err == nil {
//		t.Error("Ожидалась ошибка при отображении прогноза")
//	} else if err != os.ErrInvalid {
//		t.Error(err)
//	}
//}
