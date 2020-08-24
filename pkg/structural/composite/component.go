package composite

import (
	"fmt"
	"io"
	"os"

	"github.com/marksartdev/go-patterns/pkg/common"
)

// Интерфейс компонента меню.
type menuComponent interface {
	add(component menuComponent) error
	remove(i int) error
	getChild(i int) (menuComponent, error)
	getName() (string, error)
	getDescription() (string, error)
	getPrice() (float64, error)
	isVegetarian() (bool, error)
	print() error
	createIterator() common.Iterator
	setWriter(writer io.Writer)
}

// Компонент меню.
type component struct {
	writer io.Writer
}

// Добавляет дочерний компонент.
func (m *component) add(menuComponent) error {
	return common.UnsupportedOperationError{}
}

// Удаляет дочерний компонент.
func (m *component) remove(int) error {
	return common.UnsupportedOperationError{}
}

// Возвращает дочерние компоненты.
func (m *component) getChild(int) (menuComponent, error) {
	return nil, common.UnsupportedOperationError{}
}

// Возвращает название.
func (m *component) getName() (string, error) {
	return "", common.UnsupportedOperationError{}
}

// Возвращает описание.
func (m *component) getDescription() (string, error) {
	return "", common.UnsupportedOperationError{}
}

// Возвращает стоимость.
func (m *component) getPrice() (float64, error) {
	return 0, common.UnsupportedOperationError{}
}

// Проверяет, является ли блюдо вегетарианским.
func (m *component) isVegetarian() (bool, error) {
	return false, common.UnsupportedOperationError{}
}

// Печатает компонент.
func (m *component) print() error {
	return common.UnsupportedOperationError{}
}

// Возвращает итератор компонента.
func (m *component) createIterator() common.Iterator {
	return common.NewNullIterator()
}

// Устанавливает writer.
func (m *component) setWriter(writer io.Writer) {
	m.writer = writer
}

// Печатает.
func (m *component) write(msg string) error {
	_, err := fmt.Fprintln(m.writer, msg)

	return err
}

// Создает базовую реализацию компонента меню.
func newMenuComponent() component {
	return component{os.Stdout}
}
