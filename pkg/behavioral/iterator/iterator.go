package iterator

import "github.com/marksartdev/go-patterns/pkg/common"

// Menu Интерфейс меню.
type Menu interface {
	CreateIterator() Iterator
}

// Iterator Интерфейс итератора.
type Iterator interface {
	HasNext() bool
	Next() interface{}
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

// NewDinerMenuIterator Создает итератор для меню закусочной.
func NewDinerMenuIterator(items [maxItems]MenuItem) Iterator {
	return &dinerMenuIterator{items, 0}
}

// Итератор для меню блинной.
type pancakeHouseMenuIterator struct {
	items    common.ArrayList
	position int
}

// HasNext Проверяет, есть ли еще элемент в коллекции.
func (p *pancakeHouseMenuIterator) HasNext() bool {
	return p.position < p.items.Size()
}

// Next Возвращает следующий элемент.
func (p *pancakeHouseMenuIterator) Next() interface{} {
	item := p.items.Get(p.position)
	p.position++

	return item
}

// NewPancakeHouseMenuIterator Создает итератор для меню блинной.
func NewPancakeHouseMenuIterator(items common.ArrayList) Iterator {
	return &pancakeHouseMenuIterator{items, 0}
}
