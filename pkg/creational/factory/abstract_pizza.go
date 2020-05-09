package factory

import "fmt"

const (
	// CheesePizza Сырная пицца
	CheesePizza = "cheese"
	// PepperoniPizza Пицца "Пепперони"
	PepperoniPizza = "pepperoni"
	// ClamPizza Пицца с мидиями
	ClamPizza = "clam"
	// VeggiePizza Вегетарианская пицца
	VeggiePizza = "veggie"
)

// Pizza Интерфейс пиццы.
type Pizza interface {
	prepare()
	bake()
	cut()
	box()
	setName(string)
	GetName() string
	GetLog() []string
}

// Абстрактная пицца.
type abstractPizza struct {
	name            string
	dough           *dough
	sauce           *sauce
	cheese          []*cheese
	veggies         []*veggie
	pepperoni       *pepperoni
	clams           *clams
	sliceType       string
	abstractPrepare func(*abstractPizza)
	log             []string
}

// Приготовить пиццу.
func (a *abstractPizza) prepare() {
	a.abstractPrepare(a)

	a.log = append(a.log, fmt.Sprintf("Tossing %s", a.dough))
	a.log = append(a.log, fmt.Sprintf("Adding %s", a.sauce))

	a.log = append(a.log, "Adding cheese:")
	for _, item := range a.cheese {
		a.log = append(a.log, fmt.Sprintf("    %s", item))
	}

	a.log = append(a.log, "Adding toppings:")
	for _, item := range a.veggies {
		a.log = append(a.log, fmt.Sprintf("    %s", item))
	}

	if a.pepperoni != nil {
		a.log = append(a.log, fmt.Sprintf("    %s", a.pepperoni))
	}

	if a.clams != nil {
		a.log = append(a.log, fmt.Sprintf("    %s", a.clams))
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

// Установить название пиццы.
func (a *abstractPizza) setName(name string) {
	a.name = name
}

// GetName Получить название пиццы.
func (a *abstractPizza) GetName() string {
	return a.name
}

// GetLog Получить лог приготовления пиццы.
func (a *abstractPizza) GetLog() []string {
	return a.log
}
