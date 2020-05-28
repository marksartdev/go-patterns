package factory

// PizzaStore Интерфейс пиццерии.
type PizzaStore interface {
	OrderPizza(pizzaType string) (Pizza, error)
}

// Абстрактная пиццерия.
type abstractPizzaStore struct {
	createPizza func(pizzaType string) (Pizza, error)
}

// OrderPizza Заказать пиццу.
func (a abstractPizzaStore) OrderPizza(pizzaType string) (Pizza, error) {
	pizza, err := a.createPizza(pizzaType)
	if err != nil {
		return pizza, err
	}

	pizza = pizza.prepare()
	pizza = pizza.bake()
	pizza = pizza.cut()
	pizza = pizza.box()

	return pizza, nil
}
