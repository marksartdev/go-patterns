package observer

// Интерфейс субъекта
type subject interface {
	RegisterObserver(observer)
	RemoveObserver(observer)
	NotifyObservers(*measurements) string
	SetChanged()
	HasChanged() bool
	ClearChanged()
	GetTemperature() float64
	GetHumidity() float64
	GetPressure() float64
}

// Интерфейс наблюдателя
type observer interface {
	Update(subject, *measurements) string
}

// Измерения
type measurements struct {
	temperature float64
	humidity    float64
	pressure    float64
}
