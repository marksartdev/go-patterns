package composite

import (
	"io"

	"github.com/marksartdev/go-patterns/pkg/common"
)

// Гусь.
type goose struct {
	observable quackObservable
	common.CustomWriter
}

// Honk Крикнуть.
func (g goose) Honk() {
	g.Write("Honk")
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

// Оповестить наблюдателей (заглушка, т.к. это делает goose при крике).
func (g gooseAdapter) notifyObserver() {}

// SetWriter Установить writer.
func (g gooseAdapter) SetWriter(writer io.Writer) {
	g.goose.SetWriter(writer)
}
