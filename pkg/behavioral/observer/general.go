package observer

import "io"

// Интерфейс субъекта
type subject interface {
	RegisterObserver(observer)
	RemoveObserver(observer)
	NotifyObservers()
}

// Интерфейс наблюдателя
type observer interface {
	Update(*measurements)
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
