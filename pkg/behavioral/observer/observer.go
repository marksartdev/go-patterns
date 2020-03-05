package observer

// Интерфейс субъекта
type subject interface {
	RegisterObserver()
	RemoveObserver()
	NotifyObservers()
}

// Интерфейс наблюдателя
type observer interface {
	Update(measurements)
}

// Измерения
type measurements struct {
	temperature float64
	humidity    float64
	pressure    float64
}
