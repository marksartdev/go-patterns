package observer

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

var errStringF = "Некорректный результат (ожидалось %.1f, получено %.1f)"
var errStringS = "Некорректный результат (ожидалось %s, получено %s)"

// Битый reader
type badReader struct{}

func (r *badReader) Read(_ []byte) (n int, err error) {
	return 0, os.ErrInvalid
}

// Битый writer
type badWriter struct{}

// Write Записать в writer
func (w *badWriter) Write(_ []byte) (n int, err error) {
	return 0, os.ErrInvalid
}

func TestNewWeatherData(t *testing.T) {
	testReader := bytes.NewReader([]byte("1.1 1.2 1.3"))
	testWriter := bytes.NewBuffer(make([]byte, 0))
	expected := "Введите новые данные через пробел (temperature humidity pressure):\n"

	wd := NewWeatherData()
	wd.SetReader(testReader)
	wd.SetWriter(testWriter)

	if err := wd.SetMeasurements(); err != nil {
		t.Error(err)
	}

	if testWriter.String() != expected {
		t.Errorf(errStringS, expected, testWriter)
	}

	if wd.GetTemperature() != 1.1 {
		t.Errorf(errStringF, 1.1, wd.GetTemperature())
	}

	if wd.GetHumidity() != 1.2 {
		t.Errorf(errStringF, 1.2, wd.GetHumidity())
	}

	if wd.GetPressure() != 1.3 {
		t.Errorf(errStringF, 1.3, wd.GetPressure())
	}
}

func TestNewWeatherDataBadReader(t *testing.T) {
	testReader := new(badReader)
	testWriter := bytes.NewBuffer(make([]byte, 0))

	wd := NewWeatherData()
	wd.SetReader(testReader)
	wd.SetWriter(testWriter)

	err := wd.SetMeasurements()
	if err == nil {
		t.Error("Ожидалась ошибка при использовании битого Reader")
	} else if err != os.ErrInvalid {
		t.Error(err)
	}
}

func TestNewWeatherDataBadWriter(t *testing.T) {
	testReader := bytes.NewReader([]byte("1.1 1.2 1.3"))
	testWriter := new(badWriter)

	wd := NewWeatherData()
	wd.SetReader(testReader)
	wd.SetWriter(testWriter)

	err := wd.SetMeasurements()
	if err == nil {
		t.Error("Ожидалась ошибка при использовании битого Writer")
	} else if err != os.ErrInvalid {
		t.Error(err)
	}
}

func TestNewWeatherDataLessMeasurements(t *testing.T) {
	testReader := bytes.NewReader([]byte("1.1 1.2"))
	testWriter := bytes.NewBuffer(make([]byte, 0))

	wd := NewWeatherData()
	wd.SetReader(testReader)
	wd.SetWriter(testWriter)

	err := wd.SetMeasurements()
	if err == nil {
		t.Error("Ожидалась ошибка при вводе меньшего количества показателей")
	} else if err != io.EOF {
		t.Error(err)
	}
}

func TestNewWeatherDataNaN(t *testing.T) {
	testReader := bytes.NewReader([]byte("1.1 1.2 NotNumber"))
	testWriter := bytes.NewBuffer(make([]byte, 0))

	wd := NewWeatherData()
	wd.SetReader(testReader)
	wd.SetWriter(testWriter)

	err := wd.SetMeasurements()
	if err == nil {
		t.Error("Ожидалась ошибка при вводе не числового значения")
	} else if err.Error() != "strconv.ParseFloat: parsing \"NotNumber\": invalid syntax" {
		t.Error(err)
	}
}

func TestNewCurrentConditionsDisplay(t *testing.T) {
	testWriter := bytes.NewBuffer(make([]byte, 0))
	wd := NewWeatherData()

	data := new(measurements)
	data.temperature = 20.5
	data.humidity = 91.0
	data.pressure = 650.0

	expected := "Current conditions:\n"
	expected += fmt.Sprintf("\tTemperature: %.1f\n", data.temperature)
	expected += fmt.Sprintf("\tHumidity: %.1f\n\n", data.humidity)

	display := NewCurrentConditionsDisplay(wd)
	display.SetWriter(testWriter)

	err := display.Update(data)
	if err != nil {
		t.Error(err)
	}

	if testWriter.String() != expected {
		t.Errorf(errStringS, expected, testWriter.String())
	}
}

func TestNewCurrentConditionsDisplay_BadWriter(t *testing.T) {
	testWriter := new(badWriter)
	wd := NewWeatherData()
	data := new(measurements)

	display := NewCurrentConditionsDisplay(wd)
	display.SetWriter(testWriter)

	err := display.Update(data)
	if err == nil {
		t.Error("Ожидалась ошибка при использовании битого Writer")
	} else if err != os.ErrInvalid {
		t.Error(err)
	}
}

//
//func TestNewStatisticsDisplay(t *testing.T) {
//	buffer := bytes.NewBuffer(make([]byte, 0))
//	expected := "Statistics:\n"
//	expected += fmt.Sprintf("\tTemperature (max/min/avg): %.1f/%.1f/%.1f\n", 30.0, 20.0, 25.0)
//	expected += fmt.Sprintf("\tHumidity (max/min/avg): %.1f/%.1f/%.1f\n", 70.0, 50.0, 60.0)
//	expected += fmt.Sprintf("\tPressure (max/min/avg): %.1f/%.1f/%.1f\n\n", 700.0, 600.0, 650.0)
//
//	display := NewStatisticsDisplay()
//	display.SetOutput(buffer)
//
//	data := new(measurements)
//	data.temperature = 20.0
//	data.humidity = 50.0
//	data.pressure = 600.0
//	display.Update(data)
//
//	data.temperature = 30.0
//	data.humidity = 70.0
//	data.pressure = 700.0
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
