package composite

import "os"

// AbstractDuckFactory Абстрактная фабрика по созданию уток.
type AbstractDuckFactory interface {
	CreateMallardDuck() Quackable
	CreateRedHeatDuck() Quackable
	CreateDuckCall() Quackable
	CreateRubberDuck() Quackable
}

// DuckFactory Фабрика по созданию уток.
type DuckFactory struct{}

// CreateMallardDuck Создать крякву.
func (f DuckFactory) CreateMallardDuck() Quackable {
	d := mallardDuck{}
	d.SetWriter(os.Stdout)

	return d
}

// CreateRedHeatDuck Создать красноголовку.
func (f DuckFactory) CreateRedHeatDuck() Quackable {
	d := redheadDuck{}
	d.SetWriter(os.Stdout)

	return d
}

// CreateDuckCall Создать манок.
func (f DuckFactory) CreateDuckCall() Quackable {
	d := duckCall{}
	d.SetWriter(os.Stdout)

	return d
}

// CreateRubberDuck Создать резиновую уточку.
func (f DuckFactory) CreateRubberDuck() Quackable {
	d := rubberDuck{}
	d.SetWriter(os.Stdout)

	return d
}

// CountingDuckFactory Фабрика по созданию уток с подсчетом кряков.
type CountingDuckFactory struct{}

// CreateMallardDuck Создать крякву.
func (f CountingDuckFactory) CreateMallardDuck() Quackable {
	d := mallardDuck{}
	d.SetWriter(os.Stdout)

	return newQuackCounter(d)
}

// CreateRedHeatDuck Создать красноголовку.
func (f CountingDuckFactory) CreateRedHeatDuck() Quackable {
	d := redheadDuck{}
	d.SetWriter(os.Stdout)

	return newQuackCounter(d)
}

// CreateDuckCall Создать манок.
func (f CountingDuckFactory) CreateDuckCall() Quackable {
	d := duckCall{}
	d.SetWriter(os.Stdout)

	return newQuackCounter(d)
}

// CreateRubberDuck Создать резиновую уточку.
func (f CountingDuckFactory) CreateRubberDuck() Quackable {
	d := rubberDuck{}
	d.SetWriter(os.Stdout)

	return newQuackCounter(d)
}
