package decorator

// Интерфейс описания напитка
type descriptor interface {
	GetDescription() string
}

// Интерфейс стоимости напитка
type coster interface {
	Cost() float64
}

// Beverage Интерфейс напитка
type Beverage interface {
	descriptor
	coster
}

// Базовая структура напитка
type beverage struct {
	description string
	cost        float64
}

// GetDescription Получить описание
func (b *beverage) GetDescription() string {
	return b.description
}

// Cost Рассчитать стоимость
func (b *beverage) Cost() float64 {
	return b.cost
}

// Кофе "Домашняя смесь"
type houseBlend struct {
	beverage
}

// NewHouseBlend Создать кофе "Домашняя смесь"
func NewHouseBlend() Beverage {
	b := new(houseBlend)
	b.description = "Кофе \"Домашняя смесь\""
	b.cost = .89

	return b
}

// Кофе "Темная обжарка"
type darkRoast struct {
	beverage
}

// NewDarkRoast Создать кофе "Темная обжарка"
func NewDarkRoast() Beverage {
	b := new(darkRoast)
	b.description = "Кофе \"Темная обжарка\""
	b.cost = .99

	return b
}

// Кофе "Без кофеина"
type decaf struct {
	beverage
}

// NewDecaf Создать кофе "Без кофеина"
func NewDecaf() Beverage {
	b := new(decaf)
	b.description = "Кофе \"Без кофеина\""
	b.cost = 1.05

	return b
}

// Кофе "Эспрессо"
type espresso struct {
	beverage
}

// NewEspresso Создать кофе "Эспрессо"
func NewEspresso() Beverage {
	b := new(espresso)
	b.description = "Кофе \"Эспрессо\""
	b.cost = 1.99

	return b
}
