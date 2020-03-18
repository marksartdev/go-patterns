package decorator

// Интерфейс описания напитка
type descriptor interface {
	GetDescription() string
}

// Интерфейс стоимости напитка
type coster interface {
	Cost() float64
}

// Интерфейс дополнений
type addition interface {
	HasMilk() bool
	SetMilk()
	HasSoy() bool
	SetSoy()
	HasMocha() bool
	SetMocha()
	HasWhip() bool
	SetWhip()
}

// Beverage Интерфейс напитка
type Beverage interface {
	descriptor
	coster
	addition
}

// Базовая структура напитка
type beverage struct {
	description string
	milk        bool
	soy         bool
	mocha       bool
	whip        bool
}

// GetDescription Получить описание
func (b *beverage) GetDescription() string {
	return b.description
}

// HasMilk Проверить наличие молока
func (b *beverage) HasMilk() bool {
	return b.milk
}

// SetMilk Задать молоко
func (b *beverage) SetMilk() {
	b.milk = true
}

// HasSoy Проверить наличие сои
func (b *beverage) HasSoy() bool {
	return b.soy
}

// SetSoy Задать сою
func (b *beverage) SetSoy() {
	b.soy = true
}

// HasMocha Проверить наличие шоколада
func (b *beverage) HasMocha() bool {
	return b.mocha
}

// SetMocha Задать шоколад
func (b *beverage) SetMocha() {
	b.mocha = true
}

// HasWhip Проверить наличие взбитых сливок
func (b *beverage) HasWhip() bool {
	return b.whip
}

// SetWhip Задать взбитые сливки
func (b *beverage) SetWhip() {
	b.whip = true
}

// Cost Рассчитать стоимость
func (b *beverage) Cost() float64 {
	cost := 0.0

	if b.HasMilk() {
		cost += 10.5
	}
	if b.HasSoy() {
		cost += 5.5
	}
	if b.HasMocha() {
		cost += 15.0
	}
	if b.HasWhip() {
		cost += 7.5
	}

	return cost
}

// Напиток "Домашний кофе"
type houseBlend struct {
	beverage
}

// Cost Рассчитать стоимость
func (h *houseBlend) Cost() float64 {
	return 50.0 + h.beverage.Cost()
}

// NewHouseBlend Создать напиток "Домашний кофе"
func NewHouseBlend() Beverage {
	bv := new(houseBlend)
	bv.description = "Домашний кофе"

	return bv
}

// Напиток "Кофе темной обжарки"
type darkRoast struct {
	beverage
}

// Cost Рассчитать стоимость
func (d *darkRoast) Cost() float64 {
	return 60.0 + d.beverage.Cost()
}

// NewDarkRoast Создать напиток "Кофе темной обжарки"
func NewDarkRoast() Beverage {
	bv := new(darkRoast)
	bv.description = "Кофе темной обжарки"

	return bv
}

// Напиток "Кофе без кофеина"
type decaf struct {
	beverage
}

// Cost Рассчитать стоимость
func (d *decaf) Cost() float64 {
	return 40.0 + d.beverage.Cost()
}

// NewDecaf Создать напиток "Кофе без кофеина"
func NewDecaf() Beverage {
	bv := new(decaf)
	bv.description = "Кофе без кофеина"

	return bv
}

// Напиток "Эспрессо"
type espresso struct {
	beverage
}

// Cost Рассчитать стоимость
func (e *espresso) Cost() float64 {
	return 30.0 + e.beverage.Cost()
}

// NewEspresso Создать напиток "Эспрессо"
func NewEspresso() Beverage {
	bv := new(espresso)
	bv.description = "Эспрессо"

	return bv
}
