package factory

import "fmt"

const (
	// CheesePizza Сырная пицца.
	CheesePizza = "cheese"
	// PepperoniPizza Пицца "Пепперони".
	PepperoniPizza = "pepperoni"
	// ClamPizza Пицца с мидиями.
	ClamPizza = "clam"
	// VeggiePizza Вегетарианская пицца.
	VeggiePizza = "veggie"
)

// Pizza Интерфейс пиццы.
type Pizza interface {
	prepare() abstractPizza
	bake() abstractPizza
	cut() abstractPizza
	box() abstractPizza
	setName(name string) abstractPizza
	GetName() string
	GetLog() []string
}

// Абстрактная пицца.
type abstractPizza struct {
	name            string
	dough           dough
	sauce           sauce
	cheese          []cheese
	veggies         []veggie
	pepperoni       pepperoni
	clams           clams
	sliceType       string
	abstractPrepare func(abstractPizza) abstractPizza
	log             []string
}

// Приготовить пиццу.
func (a abstractPizza) prepare() abstractPizza {
	newA := a.abstractPrepare(a)

	newA.log = append(newA.log, fmt.Sprintf("Tossing %s", newA.dough))
	newA.log = append(newA.log, fmt.Sprintf("Adding %s", newA.sauce))

	newA.log = append(newA.log, "Adding cheese:")
	for _, item := range newA.cheese {
		newA.log = append(newA.log, fmt.Sprintf("    %s", item))
	}

	newA.log = append(newA.log, "Adding toppings:")
	for _, item := range newA.veggies {
		newA.log = append(newA.log, fmt.Sprintf("    %s", item))
	}

	if newA.pepperoni != (pepperoni{}) {
		newA.log = append(newA.log, fmt.Sprintf("    %s", newA.pepperoni))
	}

	if newA.clams != (clams{}) {
		newA.log = append(newA.log, fmt.Sprintf("    %s", newA.clams))
	}

	return newA
}

// Испечь пиццу.
func (a abstractPizza) bake() abstractPizza {
	a.log = append(a.log, "Bake for 25 minutes at 350")

	return a
}

// Разрезать пиццу.
func (a abstractPizza) cut() abstractPizza {
	a.log = append(a.log, fmt.Sprintf("Cutting the pizza into %s slices", a.sliceType))

	return a
}

// Упаковать пиццу.
func (a abstractPizza) box() abstractPizza {
	a.log = append(a.log, "Place pizza in official PizzaStore box")

	return a
}

// Установить название пиццы.
func (a abstractPizza) setName(name string) abstractPizza {
	a.name = name

	return a
}

// GetName Получить название пиццы.
func (a abstractPizza) GetName() string {
	return a.name
}

// GetLog Получить лог приготовления пиццы.
func (a abstractPizza) GetLog() []string {
	return a.log
}
