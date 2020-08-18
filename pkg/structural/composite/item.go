package composite

import "fmt"

// Блюдо.
type menuItem struct {
	menuComponent
	name        string
	description string
	vegetarian  bool
	price       float64
}

// GetName Возвращает название блюда.
func (m *menuItem) GetName() (string, error) {
	return m.name, nil
}

// GetDescription Возвращает описание блюда.
func (m *menuItem) GetDescription() (string, error) {
	return m.description, nil
}

// GetPrice Возвращает стоимость блюда.
func (m *menuItem) GetPrice() (float64, error) {
	return m.price, nil
}

// IsVegetarian Проверяет, является ли блюдо вегетарианским.
func (m *menuItem) IsVegetarian() (bool, error) {
	return m.vegetarian, nil
}

// Print Печатает блюдо.
func (m *menuItem) Print() error {
	msg := fmt.Sprintf("   %s", m.name)
	if m.vegetarian {
		msg += "(v)"
	}

	msg += fmt.Sprintf(", %.2f\n      -- %s", m.price, m.description)

	return m.write(msg)
}

// NewMenuItem Создает блюдо.
func NewMenuItem(name, description string, vegetarian bool, price float64) MenuComponent {
	return &menuItem{newMenuComponent(), name, description, vegetarian, price}
}
