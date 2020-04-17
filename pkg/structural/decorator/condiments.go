package decorator

import "fmt"

// Базовая структура декоратора
type condimentDecorator struct {
	beverage    Beverage
	description string
	cost        float64
}

// GetDescription Получить описание
func (c *condimentDecorator) GetDescription() string {
	return fmt.Sprintf("%s, %s", c.beverage.GetDescription(), c.description)
}

// Cost Рассчитать стоимость
func (c *condimentDecorator) Cost() float64 {
	return c.beverage.Cost() + c.cost
}

// Молочная пена
type milkDecorator struct {
	condimentDecorator
}

// NewMilkDecorator Задекорировать напиток молочной пеной
func NewMilkDecorator(beverage Beverage) Beverage {
	d := new(milkDecorator)
	d.beverage = beverage
	d.description = "молочная пена"
	d.cost = .10

	return d
}

// Шоколад
type mochaDecorator struct {
	condimentDecorator
}

// NewMochaDecorator Задекорировать напиток шоколадом
func NewMochaDecorator(beverage Beverage) Beverage {
	d := new(mochaDecorator)
	d.beverage = beverage
	d.description = "шоколад"
	d.cost = .20

	return d
}

// Соя
type soyDecorator struct {
	condimentDecorator
}

// NewSoyDecorator Задекорировать напиток соей
func NewSoyDecorator(beverage Beverage) Beverage {
	d := new(soyDecorator)
	d.beverage = beverage
	d.description = "соя"
	d.cost = .15

	return d
}

// Взбитые сливки
type whipDecorator struct {
	condimentDecorator
}

// NewWhipDecorator Задекорировать напиток взбитыми сливками
func NewWhipDecorator(beverage Beverage) Beverage {
	d := new(whipDecorator)
	d.beverage = beverage
	d.description = "взбитые сливки"
	d.cost = .10

	return d
}
