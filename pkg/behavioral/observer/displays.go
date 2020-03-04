package observer

import (
	"fmt"
	"io"
)

// Displayer Интерфейс дисплея
type Displayer interface {
	Display(io.Writer) error
}

// Дисплей "Текущее состояние"
type currentConditionsDisplay struct {
	temperature float64
	humidity    float64
	pressure    float64
}

// Display Отобразить дисплей
func (d *currentConditionsDisplay) Display(writer io.Writer) error {
	text := "Current conditions:\n"
	text += fmt.Sprintf("\tTemperature: %.1f\n", d.temperature)
	text += fmt.Sprintf("\tHumidity: %.1f\n", d.humidity)
	text += fmt.Sprintf("\tPressure: %.1f\n", d.pressure)

	_, err := fmt.Fprintln(writer, text)
	return err
}

// NewCurrentConditionsDisplay Создать дисплей "Текущее состояние"
func NewCurrentConditionsDisplay() Displayer {
	return &currentConditionsDisplay{}
}
