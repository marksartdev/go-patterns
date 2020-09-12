// Package composite Составной паттерн.
package composite

import "os"

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

// NewMallardDuck Создать крякву.
func NewMallardDuck() Quackable {
	d := mallardDuck{}
	d.SetWriter(os.Stdout)

	return d
}

// NewRedHeadDuck Создать красноголовку.
func NewRedHeadDuck() Quackable {
	d := redheadDuck{}
	d.SetWriter(os.Stdout)

	return d
}

// NewDuckCall Создать манок.
func NewDuckCall() Quackable {
	d := duckCall{}
	d.SetWriter(os.Stdout)

	return d
}

// NewRubberDuck Создать резиновую уточку.
func NewRubberDuck() Quackable {
	d := rubberDuck{}
	d.SetWriter(os.Stdout)

	return d
}
