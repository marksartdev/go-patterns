package observer

// Интерфейс работы с измерениями
type measurementsHandler interface {
	MeasurementsChanged() string
	SetMeasurements(float64, float64, float64) string
}

// WeatherDater Интерфейс источника метеоданных, способного быть субъектом
type WeatherDater interface {
	measurementsHandler
	subject
}

// Источник метеоданных
type weatherData struct {
	changed     bool
	observers   map[observer]struct{}
	temperature float64
	humidity    float64
	pressure    float64
}

// MeasurementsChanged Вызывается при каждом обновлении показаний датчиков
func (w *weatherData) MeasurementsChanged() string {
	return w.NotifyObservers(nil)
}

// SetMeasurements Задать измерения
func (w *weatherData) SetMeasurements(temperature, humidity, pressure float64) string {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure

	return w.MeasurementsChanged()
}

// RegisterObserver Регистрация нового наблюдателя
func (w *weatherData) RegisterObserver(newObserver observer) {
	w.observers[newObserver] = struct{}{}
}

// RemoveObserver Удалить наблюдателя из списка
func (w *weatherData) RemoveObserver(removedObserver observer) {
	delete(w.observers, removedObserver)
}

// NotifyObservers Оповестить наблюдателей
func (w *weatherData) NotifyObservers(newMeasurements *measurements) string {
	var result string

	if w.HasChanged() {
		for currentObserver := range w.observers {
			result += currentObserver.Update(w, newMeasurements)
		}

		w.ClearChanged()
	}

	return result
}

// SetChanged Зафиксировать изменения
func (w *weatherData) SetChanged() {
	w.changed = true
}

// HasChanged Проверить фиксацию изменений
func (w *weatherData) HasChanged() bool {
	return w.changed
}

// ClearChanged Снять фиксацию изменений
func (w *weatherData) ClearChanged() {
	w.changed = false
}

// GetTemperature Получить текущую температуру
func (w *weatherData) GetTemperature() float64 {
	return w.temperature
}

// GetHumidity Получить текущую влажность
func (w *weatherData) GetHumidity() float64 {
	return w.humidity
}

// GetPressure Получить текущее давление
func (w *weatherData) GetPressure() float64 {
	return w.pressure
}

// NewWeatherData Создать weatherData
func NewWeatherData() WeatherDater {
	wd := new(weatherData)
	wd.observers = make(map[observer]struct{})

	return wd
}
