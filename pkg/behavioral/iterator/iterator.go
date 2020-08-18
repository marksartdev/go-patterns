package iterator

import (
	"time"

	"github.com/marksartdev/go-patterns/pkg/common"
)

// Menu Интерфейс меню.
type Menu interface {
	CreateIterator() common.Iterator
}

// Итератор для меню закусочной.
// nolint:unused
type dinerMenuIterator struct {
	items    []MenuItem
	position int
}

// HasNext Проверяет, есть ли еще элемент в коллекции.
func (d *dinerMenuIterator) HasNext() bool {
	return d.position < len(d.items) && d.items[d.position] != nil
}

// Next Возвращает следующий элемент.
func (d *dinerMenuIterator) Next() interface{} {
	item := d.items[d.position]
	d.position++

	return item
}

// Remove Удаляет текущий элемент.
func (d *dinerMenuIterator) Remove() error {
	if d.position <= 0 {
		return common.IllegalStateError{}
	}

	if d.items[d.position-1] != nil {
		for i := d.position - 1; i < len(d.items)-1; i++ {
			d.items[i] = d.items[i+1]
		}

		d.items[len(d.items)-1] = nil
	}

	return nil
}

// Создает итератор для меню закусочной.
// nolint:unused,deadcode
func newDinerMenuIterator(items []MenuItem) common.Iterator {
	return &dinerMenuIterator{items, 0}
}

// Альтернативный итератор для меню закусочной.
type alternatingDinerMenuIterator struct {
	items    []MenuItem
	position int
}

// HasNext Проверяет, есть ли еще элемент в коллекции.
func (a *alternatingDinerMenuIterator) HasNext() bool {
	return a.position < len(a.items) && a.items[a.position] != nil
}

// Next Возвращает следующий элемент.
func (a *alternatingDinerMenuIterator) Next() interface{} {
	item := a.items[a.position]
	a.position += 2

	return item
}

// Remove Удаляет текущий элемент.
func (a *alternatingDinerMenuIterator) Remove() error {
	return common.UnsupportedOperationError{}
}

// Создает альтернативный итератор для меню закусочной.
func newAlternatingDinerMenuIterator(items []MenuItem) common.Iterator {
	menuIterator := &alternatingDinerMenuIterator{items, 0}
	weekday := int(time.Now().Weekday())
	// nolint:gomnd
	menuIterator.position = weekday % 2

	return menuIterator
}
