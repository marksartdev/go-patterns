// Package adapter Паттерн "Адаптер".
package adapter

import (
	"math/rand"
	"strings"

	"github.com/marksartdev/go-patterns/pkg/common"
)

const flyConverter = 3

// Утиный адаптер для индюшки.
type turkeyAdapter struct {
	turkey Turkey
}

// Quack Крякает.
func (t turkeyAdapter) Quack() string {
	return t.turkey.Gobble()
}

// Fly Летает.
func (t turkeyAdapter) Fly() string {
	log := make([]string, flyConverter)

	for i := 0; i < flyConverter; i++ {
		log[i] = t.turkey.Fly()
	}

	return strings.Join(log, "\n")
}

// NewTurkeyAdapter Создает утиный адаптер для индюшки.
func NewTurkeyAdapter(turkey Turkey) Duck {
	return turkeyAdapter{turkey}
}

// Индюшачий адаптер для утки.
type duckAdapter struct {
	duck   Duck
	random *rand.Rand
}

// Gobble Кулдыкает.
func (d duckAdapter) Gobble() string {
	return d.duck.Quack()
}

// Fly Летает.
func (d duckAdapter) Fly() string {
	if d.random.Intn(flyConverter) == 0 {
		return d.duck.Fly()
	}

	return ""
}

// NewDuckAdapter Создает индюшачий адаптер для утки.
func NewDuckAdapter(duck Duck, seed int64) Turkey {
	source := rand.NewSource(seed)
	// nolint:gosec // Example
	random := rand.New(source)

	return duckAdapter{duck, random}
}

// Адаптер итератора для перечисления.
type enumerationIterator struct {
	enumeration Enumeration
}

// HasNext Проверяет наличие следующего элемента.
func (e *enumerationIterator) HasNext() bool {
	return e.enumeration.HasMoreElements()
}

// Next Возвращает следующий элемент.
func (e *enumerationIterator) Next() interface{} {
	return e.enumeration.NextElement()
}

// Remove Удаляет текущий элемент.
func (e *enumerationIterator) Remove() error {
	return common.UnsupportedOperationError{}
}

// NewEnumerationIterator Создает адаптер итератора для перечисления.
func NewEnumerationIterator(enumeration Enumeration) common.Iterator {
	return &enumerationIterator{enumeration}
}

// Адаптер перечисления для итератора.
type iteratorEnumeration struct {
	iterator common.Iterator
}

// HasMoreElements Проверяет наличие следующего элемента.
func (i *iteratorEnumeration) HasMoreElements() bool {
	return i.iterator.HasNext()
}

// NextElement Возвращает следующий элемент.
func (i *iteratorEnumeration) NextElement() interface{} {
	return i.iterator.Next()
}

// NewIteratorEnumeration Создает адаптер перечисления для итератора.
func NewIteratorEnumeration(iterator common.Iterator) Enumeration {
	return &iteratorEnumeration{iterator}
}
