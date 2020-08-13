package iterator

import "github.com/marksartdev/go-patterns/pkg/common"

// Menu Интерфейс меню.
type Menu interface {
	CreateIterator() common.Iterator
}

// Итератор для меню закусочной.
type dinerMenuIterator struct {
	items    [maxItems]MenuItem
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

// NewDinerMenuIterator Создает итератор для меню закусочной.
func NewDinerMenuIterator(items [maxItems]MenuItem) common.Iterator {
	return &dinerMenuIterator{items, 0}
}
