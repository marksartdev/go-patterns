package composite

import "github.com/marksartdev/go-patterns/pkg/common"

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

// NewFlock Создать стаю.
func NewFlock() Flock {
	f := flock{}
	f.quackers = common.NewArrayList()

	return f
}
