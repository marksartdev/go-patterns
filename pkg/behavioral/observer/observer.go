package observer

// Интерфейс субъекта.
type subject interface {
	RegisterObserver(observer)
	RemoveObserver(observer)
	NotifyObservers(*Measurements) string
	SetChanged()
	HasChanged() bool
	ClearChanged()
	GetTemperature() float64
	GetHumidity() float64
	GetPressure() float64
}

// Интерфейс наблюдателя.
type observer interface {
	Update(subject, *Measurements) string
}

// Measurements Измерения.
type Measurements struct {
	Temperature float64
	Humidity    float64
	Pressure    float64
}
