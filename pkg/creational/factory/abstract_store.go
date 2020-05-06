package factory

// PizzaStore Интерфейс пиццерии.
type PizzaStore interface {
	OrderPizza(string) Pizza
}

// Абстрактная пиццерия.
type abstractPizzaStore struct {
	createPizza func(string) Pizza
}

// OrderPizza Заказать пиццу.
func (a *abstractPizzaStore) OrderPizza(pizzaType string) Pizza {
	pizza := a.createPizza(pizzaType)

	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()

	return pizza
}
