package composite

import (
	"fmt"
	"io"
	"os"

	"github.com/marksartdev/go-patterns/pkg/common"
)

// MenuComponent Интерфейс компонента меню.
type MenuComponent interface {
	Add(component MenuComponent) error
	Remove(i int) error
	GetChild(i int) (MenuComponent, error)
	GetName() (string, error)
	GetDescription() (string, error)
	GetPrice() (float64, error)
	IsVegetarian() (bool, error)
	Print() error
	CreateIterator() common.Iterator
	SetWriter(writer io.Writer)
}

// Компонент меню.
type menuComponent struct {
	writer io.Writer
}

// Add Добавляет дочерний компонент.
func (m *menuComponent) Add(component MenuComponent) error {
	return common.UnsupportedOperationError{}
}

// Remove Удаляет дочерний компонент.
func (m *menuComponent) Remove(i int) error {
	return common.UnsupportedOperationError{}
}

// GetChild Возвращает дочерние компоненты.
func (m *menuComponent) GetChild(i int) (MenuComponent, error) {
	return nil, common.UnsupportedOperationError{}
}

// GetName Возвращает название.
func (m *menuComponent) GetName() (string, error) {
	return "", common.UnsupportedOperationError{}
}

// GetDescription Возвращает описание.
func (m *menuComponent) GetDescription() (string, error) {
	return "", common.UnsupportedOperationError{}
}

// GetPrice Возвращает стоимость.
func (m *menuComponent) GetPrice() (float64, error) {
	return 0, common.UnsupportedOperationError{}
}

// IsVegetarian Проверяет, является ли блюдо вегетарианским.
func (m *menuComponent) IsVegetarian() (bool, error) {
	return false, common.UnsupportedOperationError{}
}

// Print Печатает компонент.
func (m *menuComponent) Print() error {
	return common.UnsupportedOperationError{}
}

// CreateIterator Возвращает итератор компонента.
func (m *menuComponent) CreateIterator() common.Iterator {
	return common.NewNullIterator()
}

// SetWriter Устанавливает writer.
func (m *menuComponent) SetWriter(writer io.Writer) {
	m.writer = writer
}

// Печатает.
func (m *menuComponent) write(msg string) error {
	_, err := fmt.Fprintln(m.writer, msg)

	return err
}

// Создает базовую реализацию компонента меню.
func newMenuComponent() menuComponent {
	return menuComponent{os.Stdout}
}
