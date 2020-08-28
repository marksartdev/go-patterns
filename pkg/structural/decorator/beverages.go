// Package decorator Паттерн "Декоратор".
package decorator

// Small Маленький стакан.
const Small = 0

// Medium Средний стакан.
const Medium = 1

// Large Большой стакан.
const Large = 2

// Интерфейс описания напитка.
type descriptor interface {
	GetDescription() string
}

// Интерфейс стоимости напитка.
type coster interface {
	Cost() float64
}

// Интерфейс размеров стакана.
type sizer interface {
	GetSize() int
	SetSize(int)
}

// Beverage Интерфейс напитка.
type Beverage interface {
	descriptor
	coster
	sizer
}

// Базовая структура напитка.
type beverage struct {
	description string
	cost        [3]float64
	size        int
}

// GetDescription Получить описание.
func (b *beverage) GetDescription() string {
	return b.description
}

// Cost Рассчитать стоимость.
func (b *beverage) Cost() float64 {
	return b.cost[b.size]
}

// GetSize Получить размер стакана.
func (b *beverage) GetSize() int {
	return b.size
}

// SetSize Установить размер стакана.
func (b *beverage) SetSize(size int) {
	b.size = size
}

// Кофе "Домашняя смесь".
type houseBlend struct {
	beverage
}

// NewHouseBlend Создать кофе "Домашняя смесь".
func NewHouseBlend() Beverage {
	b := new(houseBlend)
	b.description = "Кофе \"Домашняя смесь\""
	b.cost = [3]float64{.89, .99, 1.09}
	b.size = Small

	return b
}

// Кофе "Темная обжарка".
type darkRoast struct {
	beverage
}

// NewDarkRoast Создать кофе "Темная обжарка".
func NewDarkRoast() Beverage {
	b := new(darkRoast)
	b.description = "Кофе \"Темная обжарка\""
	b.cost = [3]float64{.99, 1.09, 1.19}
	b.size = Small

	return b
}

// Кофе "Без кофеина".
type decaf struct {
	beverage
}

// NewDecaf Создать кофе "Без кофеина".
func NewDecaf() Beverage {
	b := new(decaf)
	b.description = "Кофе \"Без кофеина\""
	b.cost = [3]float64{1.05, 1.15, 1.25}
	b.size = Small

	return b
}

// Кофе "Эспрессо".
type espresso struct {
	beverage
}

// NewEspresso Создать кофе "Эспрессо".
func NewEspresso() Beverage {
	b := new(espresso)
	b.description = "Кофе \"Эспрессо\""
	b.cost = [3]float64{1.99, 2.09, 2.19}
	b.size = Small

	return b
}
