package composite

import "fmt"

// Гусь.
type goose struct {
	observable quackObservable
	customWriter
}

// Honk Крикнуть.
func (g goose) Honk() {
	g.write("Honk")
	g.notifyObserver()
}

// RegisterObserver Зарегистрировать наблюдателя.
func (g goose) RegisterObserver(observer observer) {
	g.observable.RegisterObserver(observer)
}

// Оповестить наблюдателей.
func (g goose) notifyObserver() {
	g.observable.notifyObserver()
}

func (g goose) String() string {
	return "Goose"
}

// Адаптер для гуся.
type gooseAdapter struct {
	goose goose
}

// Quack Крякнуть.
func (g gooseAdapter) Quack() {
	g.goose.Honk()
}

// RegisterObserver Зарегистрировать наблюдателя.
func (g gooseAdapter) RegisterObserver(observer observer) {
	g.goose.RegisterObserver(observer)
}

// Оповестить наблюдателей.
func (g gooseAdapter) notifyObserver() {
	g.goose.notifyObserver()
}

func (g gooseAdapter) String() string {
	return fmt.Sprintf("%s like Duck", g.goose)
}
