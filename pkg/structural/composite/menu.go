package composite

import (
	"fmt"
	"io"

	"github.com/marksartdev/go-patterns/pkg/common"
)

// Меню.
type menu struct {
	component
	menuComponents common.ArrayList
	name           string
	description    string
	iterator       common.Iterator
}

// Добавляет дочерний компонент.
func (m *menu) add(component menuComponent) error {
	m.menuComponents.Add(component)

	return nil
}

// Удаляет дочерний компонент.
func (m *menu) remove(i int) error {
	m.menuComponents.Remove(i)

	return nil
}

// Возвращает дочерние компоненты.
func (m *menu) getChild(i int) (menuComponent, error) {
	if item, ok := m.menuComponents.Get(i).(menuComponent); ok {
		return item, nil
	}

	return nil, nil
}

// Возвращает название меню.
func (m *menu) getName() (string, error) {
	return m.name, nil
}

// Возвращает описание меню.
func (m *menu) getDescription() (string, error) {
	return m.description, nil
}

// Печатает меню.
func (m *menu) print() error {
	msg := fmt.Sprintf("\n%s,  %s\n--------------------", m.name, m.description)

	err := m.write(msg)
	if err != nil {
		return err
	}

	iterator := m.menuComponents.Iterator()
	for iterator.HasNext() {
		if next, ok := iterator.Next().(menuComponent); ok {
			err = next.print()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Создает итератор для компонента меню.
func (m *menu) createIterator() common.Iterator {
	if m.iterator == nil {
		m.iterator = newCompositeIterator(m.menuComponents.Iterator())
	}

	return m.iterator
}

// Устанавливает writer.
func (m *menu) setWriter(writer io.Writer) {
	m.writer = writer

	iterator := m.menuComponents.Iterator()
	for iterator.HasNext() {
		if next, ok := iterator.Next().(menuComponent); ok {
			next.setWriter(writer)
		}
	}
}

// Создает меню.
func newMenu(name, description string) menuComponent {
	return &menu{newMenuComponent(), common.NewArrayList(), name, description, nil}
}
