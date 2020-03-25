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
	bv := new(houseBlend)
	bv.description = "Кофе \"Домашняя смесь\""
	bv.cost = .89

	return bv
}

// Кофе "Темная обжарка"
type darkRoast struct {
	beverage
}

// NewDarkRoast Создать кофе "Темная обжарка"
func NewDarkRoast() Beverage {
	bv := new(darkRoast)
	bv.description = "Кофе \"Темная обжарка\""
	bv.cost = .99

	return bv
}

// Кофе "Без кофеина"
type decaf struct {
	beverage
}

// NewDecaf Создать кофе "Без кофеина"
func NewDecaf() Beverage {
	bv := new(decaf)
	bv.description = "Кофе \"Без кофеина\""
	bv.cost = 1.05

	return bv
}

// Кофе "Эспрессо"
type espresso struct {
	beverage
}

// NewEspresso Создать кофе "Эспрессо"
func NewEspresso() Beverage {
	bv := new(espresso)
	bv.description = "Кофе \"Эспрессо\""
	bv.cost = 1.99

	return bv
}
