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
}

// GetDescription Получить описание
func (b *beverage) GetDescription() string {
	return b.description
}

// Напиток "Домашний кофе"
type houseBlend struct {
	beverage
}

// Cost Рассчитать стоимость
func (h *houseBlend) Cost() float64 {
	return 50.0
}

// NewHouseBlend Создать напиток "Домашний кофе"
func NewHouseBlend() Beverage {
	bv := new(houseBlend)
	bv.description = "Домашний кофе"

	return bv
}

// Напиток "Темный жаренный кофе"
type darkRoast struct {
	beverage
}

// Cost Рассчитать стоимость
func (d *darkRoast) Cost() float64 {
	return 60.0
}

// NewDarkRoast Создать напиток "Темный жаренный кофе"
func NewDarkRoast() Beverage {
	bv := new(darkRoast)
	bv.description = "Темный жаренный кофе"

	return bv
}

// Напиток "Без кофеина"
type decaf struct {
	beverage
}

// Cost Рассчитать стоимость
func (d *decaf) Cost() float64 {
	return 40.0
}

// NewDecaf Создать напиток "Без кофеина"
func NewDecaf() Beverage {
	bv := new(decaf)
	bv.description = "Без кофеина"

	return bv
}

// Напиток "Эспрессо"
type espresso struct {
	beverage
}

// Cost Рассчитать стоимость
func (e *espresso) Cost() float64 {
	return 30.0
}

// NewEspresso Создать напиток "Эспрессо"
func NewEspresso() Beverage {
	bv := new(espresso)
	bv.description = "Эспрессо"

	return bv
}
