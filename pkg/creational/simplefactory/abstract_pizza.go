// Package simplefactory Простая фабрика.
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
	prepare()
	bake()
	cut()
	box()
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
func (a *abstractPizza) prepare() {
	a.log = append(a.log,
		fmt.Sprintf("Preparing %s", a.name),
		fmt.Sprintf("Tossing dough... %s", a.dough),
		fmt.Sprintf("Adding sauce... %s", a.sauce),
		"Adding toppings:",
	)

	for _, topping := range a.toppings {
		a.log = append(a.log, fmt.Sprintf("    %s", topping))
	}
}

// Испечь пиццу.
func (a *abstractPizza) bake() {
	a.log = append(a.log, "Bake for 25 minutes at 350")
}

// Разрезать пиццу.
func (a *abstractPizza) cut() {
	a.log = append(a.log, fmt.Sprintf("Cutting the pizza into %s slices", a.sliceType))
}

// Упаковать пиццу.
func (a *abstractPizza) box() {
	a.log = append(a.log, "Place pizza in official PizzaStore box")
}

// GetName Получить название пиццы.
func (a *abstractPizza) GetName() string {
	return a.name
}

// GetLog Получить лог приготовления пиццы.
func (a *abstractPizza) GetLog() []string {
	return a.log
}
