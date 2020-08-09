package simplefactory

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
	GetName() string
	GetLog() []string
}

// Абстрактная пицца.
type abstractPizza struct {
	name      string
	dough     string
	sauce     string
	toppings  []string
	sliceType string
	log       []string
}

// Приготовить пиццу.
func (a abstractPizza) prepare() abstractPizza {
	a.log = append(a.log, fmt.Sprintf("Preparing %s", a.name))
	a.log = append(a.log, fmt.Sprintf("Tossing dough... %s", a.dough))
	a.log = append(a.log, fmt.Sprintf("Adding sauce... %s", a.sauce))

	a.log = append(a.log, "Adding toppings:")
	for _, topping := range a.toppings {
		a.log = append(a.log, fmt.Sprintf("    %s", topping))
	}

	return a
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

// GetName Получить название пиццы.
func (a abstractPizza) GetName() string {
	return a.name
}

// GetLog Получить лог приготовления пиццы.
func (a abstractPizza) GetLog() []string {
	return a.log
}
