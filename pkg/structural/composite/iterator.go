package composite

import "github.com/marksartdev/go-patterns/pkg/common"

// Итератор комбинации.
type compositeIterator struct {
	stack []common.Iterator
}

// HasNext Проверяет, есть ли еще элемент в коллекции.
func (c *compositeIterator) HasNext() bool {
	if len(c.stack) == 0 {
		return false
	}

	n := len(c.stack) - 1

	iterator := c.stack[n]
	if !iterator.HasNext() {
		c.stack = c.stack[:n]

		return c.HasNext()
	}

	return true
}

// Next Возвращает следующий элемент.
func (c *compositeIterator) Next() interface{} {
	if c.HasNext() {
		n := len(c.stack) - 1
		iterator := c.stack[n]

		component := iterator.Next()
		if m, ok := component.(*menu); ok {
			c.stack = append(c.stack, m.CreateIterator())
		}

		return component
	}

	return nil
}

// Remove Удаляет текущий элемент.
func (c *compositeIterator) Remove() error {
	return common.UnsupportedOperationError{}
}

// Создает итератор комбинации.
func newCompositeIterator(iterator common.Iterator) common.Iterator {
	newIterator := &compositeIterator{}
	newIterator.stack = append(newIterator.stack, iterator)

	return newIterator
}
