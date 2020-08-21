package composite

import (
	"fmt"

	"github.com/marksartdev/go-patterns/pkg/common"
)

// Меню.
type menu struct {
	menuComponent
	menuComponents common.ArrayList
	name           string
	description    string
	iterator       common.Iterator
}

// Add Добавляет дочерний компонент.
func (m *menu) Add(component MenuComponent) error {
	m.menuComponents.Add(component)

	return nil
}

// Remove Удаляет дочерний компонент.
func (m *menu) Remove(i int) error {
	m.menuComponents.Remove(i)

	return nil
}

// GetChild Возвращает дочерние компоненты.
func (m *menu) GetChild(i int) (MenuComponent, error) {
	if item, ok := m.menuComponents.Get(i).(MenuComponent); ok {
		return item, nil
	}

	return nil, nil
}

// GetName Возвращает название меню.
func (m *menu) GetName() (string, error) {
	return m.name, nil
}

// GetDescription Возвращает описание меню.
func (m *menu) GetDescription() (string, error) {
	return m.description, nil
}

// Print Печатает меню.
func (m *menu) Print() error {
	msg := fmt.Sprintf("\n%s,  %s\n--------------------", m.name, m.description)

	err := m.write(msg)
	if err != nil {
		return err
	}

	iterator := m.menuComponents.Iterator()
	for iterator.HasNext() {
		if component, ok := iterator.Next().(MenuComponent); ok {
			err = component.Print()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CreateIterator Создает итератор для компонента меню.
func (m *menu) CreateIterator() common.Iterator {
	if m.iterator == nil {
		m.iterator = newCompositeIterator(m.menuComponents.Iterator())
	}

	return m.iterator
}

// NewMenu Создает меню.
func NewMenu(name, description string) MenuComponent {
	return &menu{newMenuComponent(), common.NewArrayList(), name, description, nil}
}
