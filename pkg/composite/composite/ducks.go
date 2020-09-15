// Package composite Составной паттерн.
package composite

import "github.com/marksartdev/go-patterns/pkg/common"

// Quackable interface.
type Quackable interface {
	Quack()
	quackObservable
	common.CustomWriterSetter
}

// Кряква.
type mallardDuck struct {
	observable quackObservable
	common.CustomWriter
}

// Quack Крякнуть.
func (m mallardDuck) Quack() {
	m.Write("Quack")
	m.notifyObserver()
}

// RegisterObserver Зарегистрировать наблюдателя.
func (m mallardDuck) RegisterObserver(observer observer) {
	m.observable.RegisterObserver(observer)
}

// Оповестить наблюдателей.
func (m mallardDuck) notifyObserver() {
	m.observable.notifyObserver()
}

func (m mallardDuck) String() string {
	return "Mallard Duck"
}

// Красноголовка.
type redheadDuck struct {
	observable quackObservable
	common.CustomWriter
}

// Quack Крякнуть.
func (r redheadDuck) Quack() {
	r.Write("Quack")
	r.notifyObserver()
}

// RegisterObserver Зарегистрировать наблюдателя.
func (r redheadDuck) RegisterObserver(observer observer) {
	r.observable.RegisterObserver(observer)
}

// Оповестить наблюдателей.
func (r redheadDuck) notifyObserver() {
	r.observable.notifyObserver()
}

func (r redheadDuck) String() string {
	return "Redhead Duck"
}

// Манок.
type duckCall struct {
	observable quackObservable
	common.CustomWriter
}

// Quack Крякнуть.
func (d duckCall) Quack() {
	d.Write("Kwak")
	d.notifyObserver()
}

// RegisterObserver Зарегистрировать наблюдателя.
func (d duckCall) RegisterObserver(observer observer) {
	d.observable.RegisterObserver(observer)
}

// Оповестить наблюдателей.
func (d duckCall) notifyObserver() {
	d.observable.notifyObserver()
}

func (d duckCall) String() string {
	return "Duck Call"
}

// Резиновая уточка.
type rubberDuck struct {
	observable quackObservable
	common.CustomWriter
}

// Quack Крякнуть.
func (r rubberDuck) Quack() {
	r.Write("Squeak")
	r.notifyObserver()
}

// RegisterObserver Зарегистрировать наблюдателя.
func (r rubberDuck) RegisterObserver(observer observer) {
	r.observable.RegisterObserver(observer)
}

// Оповестить наблюдателей.
func (r rubberDuck) notifyObserver() {
	r.observable.notifyObserver()
}

func (r rubberDuck) String() string {
	return "Rubber Duck"
}
