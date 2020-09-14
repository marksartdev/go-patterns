// Package composite Составной паттерн.
package composite

// Quackable interface.
type Quackable interface {
	quackObservable
	Quack()
}

// Кряква.
type mallardDuck struct {
	observable quackObservable
	customWriter
}

// Quack Крякнуть.
func (m mallardDuck) Quack() {
	m.write("Quack")
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
	customWriter
}

// Quack Крякнуть.
func (r redheadDuck) Quack() {
	r.write("Quack")
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
	customWriter
}

// Quack Крякнуть.
func (d duckCall) Quack() {
	d.write("Kwak")
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
	customWriter
}

// Quack Крякнуть.
func (r rubberDuck) Quack() {
	r.write("Squeak")
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
