package composite

import "os"

// AbstractGooseFactory Абстрактная фабрика гусей.
type AbstractGooseFactory interface {
	CreateGoose() Quackable
}

// GooseFactory Фабрика гусей.
type GooseFactory struct{}

// CreateGoose Создать гуся.
func (f GooseFactory) CreateGoose() Quackable {
	g := goose{}
	g.observable = newObservable(g)
	g.SetWriter(os.Stdout)

	return gooseAdapter{g}
}

// CountingGooseFactory Фабрика гусей с подсчетом криков.
type CountingGooseFactory struct{}

// CreateGoose Создать гуся.
func (f CountingGooseFactory) CreateGoose() Quackable {
	g := goose{}
	g.observable = newObservable(g)
	g.SetWriter(os.Stdout)

	ga := gooseAdapter{g}

	return newQuackCounter(ga)
}
