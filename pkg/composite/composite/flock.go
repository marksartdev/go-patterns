package composite

import (
	"io"

	"github.com/marksartdev/go-patterns/pkg/common"
)

// Flock Интерфейс стаи.
type Flock interface {
	Add(quacker Quackable)
	Quackable
}

// Стая.
type flock struct {
	quackers common.ArrayList
}

// Add Добавить квакера в стаю.
func (f flock) Add(quacker Quackable) {
	f.quackers.Add(quacker)
}

// Quack Крякнуть.
func (f flock) Quack() {
	iterator := f.quackers.Iterator()
	for iterator.HasNext() {
		if quacker, ok := iterator.Next().(Quackable); ok {
			quacker.Quack()
		}
	}
}

// RegisterObserver Зарегистрировать наблюдателя.
func (f flock) RegisterObserver(observer observer) {
	iterator := f.quackers.Iterator()
	for iterator.HasNext() {
		if quacker, ok := iterator.Next().(Quackable); ok {
			quacker.RegisterObserver(observer)
		}
	}
}

// Оповестить наблюдателей (заглушка, т.к. это делают сами утки при кряканье).
func (f flock) notifyObserver() {}

// SetWriter Установить writer.
func (f flock) SetWriter(writer io.Writer) {
	iterator := f.quackers.Iterator()
	for iterator.HasNext() {
		if quacker, ok := iterator.Next().(Quackable); ok {
			quacker.SetWriter(writer)
		}
	}
}

// NewFlock Создать стаю.
func NewFlock() Flock {
	f := flock{}
	f.quackers = common.NewArrayList()

	return f
}
