package factory

import "fmt"

// Pizza Интерфейс пиццы
type Pizza interface {
	Prepare() string
	Bake() string
	Cut() string
	Box() string
}

// Базовая структура пиццы
type pizza struct {
	name string
}

// Prepare Приготовить пиццу
func (p *pizza) Prepare() string {
	return fmt.Sprintf("Готовим пиццу %q", p.name)
}

// Bake Испечь пиццу
func (p *pizza) Bake() string {
	return fmt.Sprintf("Выпекаем пиццу %q", p.name)
}

// Cut Разрезать пиццу
func (p *pizza) Cut() string {
	return fmt.Sprintf("Разрезаем пиццу %q", p.name)
}

// Box Упаковать пиццу
func (p *pizza) Box() string {
	return fmt.Sprintf("Упаковываем пиццу %q", p.name)
}

// Сырная пицца
type cheesePizza struct {
	pizza
}

// Создать сырную пиццу
func newCheesePizza() Pizza {
	p := new(cheesePizza)
	p.name = "Сырная"

	return p
}

// Пицца "Пепперони"
type pepperoniPizza struct {
	pizza
}

// Создать пиццу "Пепперони"
func newPepperoniPizza() Pizza {
	p := new(pepperoniPizza)
	p.name = "Пепперони"

	return p
}

// Пицца с мидиями
type clamPizza struct {
	pizza
}

// Создать пиццу с мидиями
func newClamPizza() Pizza {
	p := new(clamPizza)
	p.name = "С мидиями"

	return p
}

// Вегетарианская пицца
type veggiePizza struct {
	pizza
}

// Создать вегетарианскую пиццу
func newVeggiePizza() Pizza {
	p := new(veggiePizza)
	p.name = "Вегетарианская"

	return p
}
