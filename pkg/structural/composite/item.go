package composite

import "fmt"

// Блюдо.
type menuItem struct {
	component
	name        string
	description string
	vegetarian  bool
	price       float64
}

// Возвращает название блюда.
func (m *menuItem) getName() (string, error) {
	return m.name, nil
}

// Возвращает описание блюда.
func (m *menuItem) getDescription() (string, error) {
	return m.description, nil
}

// Возвращает стоимость блюда.
func (m *menuItem) getPrice() (float64, error) {
	return m.price, nil
}

// Проверяет, является ли блюдо вегетарианским.
func (m *menuItem) isVegetarian() (bool, error) {
	return m.vegetarian, nil
}

// Печатает блюдо.
func (m *menuItem) print() error {
	msg := fmt.Sprintf("   %s", m.name)
	if m.vegetarian {
		msg += "(v)"
	}

	msg += fmt.Sprintf(", %.2f\n      -- %s", m.price, m.description)

	return m.write(msg)
}

// Создает блюдо.
func newMenuItem(name, description string, vegetarian bool, price float64) menuComponent {
	return &menuItem{newMenuComponent(), name, description, vegetarian, price}
}
