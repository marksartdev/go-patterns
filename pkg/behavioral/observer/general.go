package observer

// Интерфейс субъекта
type subject interface {
	RegisterObserver(observer)
	RemoveObserver(observer)
	NotifyObservers()
	GetTemperature() float64
	GetHumidity() float64
	GetPressure() float64
}

// Интерфейс наблюдателя
type observer interface {
	Update(*measurements) string
}

// Измерения
type measurements struct {
	temperature float64
	humidity    float64
	pressure    float64
}
