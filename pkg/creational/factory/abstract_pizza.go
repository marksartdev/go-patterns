package factory

import "fmt"

const (
	// Cheese Сырная пицца
	Cheese = "cheese"
	// Pepperoni Пицца "Пепперони"
	Pepperoni = "pepperoni"
	// Clam Пицца с мидиями
	Clam = "clam"
	// Veggie Вегетарианская пицца
	Veggie = "veggie"
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
func (s *abstractPizza) prepare() {
	s.log = append(s.log, fmt.Sprintf("Preparing %s", s.name))
	s.log = append(s.log, fmt.Sprintf("Tossing dough... %s", s.dough))
	s.log = append(s.log, fmt.Sprintf("Adding sauce... %s", s.sauce))

	s.log = append(s.log, "Adding toppings:")
	for _, topping := range s.toppings {
		s.log = append(s.log, fmt.Sprintf("    %s", topping))
	}
}

// Испечь пиццу.
func (s *abstractPizza) bake() {
	s.log = append(s.log, "Bake for 25 minutes at 350")
}

// Разрезать пиццу.
func (s *abstractPizza) cut() {
	s.log = append(s.log, fmt.Sprintf("Cutting the pizza into %s slices", s.sliceType))
}

// Упаковать пиццу.
func (s *abstractPizza) box() {
	s.log = append(s.log, "Place pizza in official PizzaStore box")
}

// GetName Получить название пиццы.
func (s *abstractPizza) GetName() string {
	return s.name
}

// GetLog Получить лог приготовления пиццы.
func (s *abstractPizza) GetLog() []string {
	return s.log
}
