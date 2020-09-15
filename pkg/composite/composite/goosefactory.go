package composite

import "github.com/marksartdev/go-patterns/pkg/common"

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
	g.CustomWriter = common.NewCustomWriter()

	return gooseAdapter{g}
}

// CountingGooseFactory Фабрика гусей с подсчетом криков.
type CountingGooseFactory struct{}

// CreateGoose Создать гуся.
func (f CountingGooseFactory) CreateGoose() Quackable {
	g := goose{}
	g.observable = newObservable(g)
	g.CustomWriter = common.NewCustomWriter()

	ga := gooseAdapter{g}

	return newQuackCounter(ga)
}
