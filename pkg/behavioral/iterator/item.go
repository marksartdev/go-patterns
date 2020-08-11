package iterator

// MenuItem Интерфейс блюда.
type MenuItem interface {
	GetName() string
	GetDescription() string
	GetPrice() float64
	IsVegetarian() bool
}

// Блюдо.
type menuItem struct {
	name        string
	description string
	vegetarian  bool
	price       float64
}

// GetName Возвращает название блюда.
func (m menuItem) GetName() string {
	return m.name
}

// GetDescription Возвращает описание блюда.
func (m menuItem) GetDescription() string {
	return m.description
}

// GetPrice Возвращает стоимость блюда.
func (m menuItem) GetPrice() float64 {
	return m.price
}

// IsVegetarian Проверяет, является ли блюдо вегетарианским.
func (m menuItem) IsVegetarian() bool {
	return m.vegetarian
}

// Создает блюдо.
func newMenuItem(name, description string, vegetarian bool, price float64) MenuItem {
	return menuItem{name, description, vegetarian, price}
}
