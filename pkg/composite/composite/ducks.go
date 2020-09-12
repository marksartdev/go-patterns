// Package composite Составной паттерн.
package composite

// Quackable interface.
type Quackable interface {
	Quack()
}

// Кряква.
type mallardDuck struct {
	customWriter
}

// Quack Крякнуть.
func (m mallardDuck) Quack() {
	m.write("Quack")
}

// Красноголовка.
type redheadDuck struct {
	customWriter
}

// Quack Крякнуть.
func (r redheadDuck) Quack() {
	r.write("Quack")
}

// Манок.
type duckCall struct {
	customWriter
}

// Quack Крякнуть.
func (d duckCall) Quack() {
	d.write("Kwak")
}

// Резиновая уточка.
type rubberDuck struct {
	customWriter
}

// Quack Крякнуть.
func (r rubberDuck) Quack() {
	r.write("Squeak")
}
