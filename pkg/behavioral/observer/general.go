package observer

import "io"

// Интерфейс субъекта
type subject interface {
	RegisterObserver(observer)
	RemoveObserver(observer)
	NotifyObservers() []error
	GetTemperature() float64
	GetHumidity() float64
	GetPressure() float64
}

// Интерфейс наблюдателя
type observer interface {
	Update(*measurements) error
}

// Измерения
type measurements struct {
	temperature float64
	humidity    float64
	pressure    float64
}

// Интерфейс чтения данных
type reader interface {
	SetReader(io.Reader)
}

// Интерфейс записи данных
type writer interface {
	SetWriter(io.Writer)
}
