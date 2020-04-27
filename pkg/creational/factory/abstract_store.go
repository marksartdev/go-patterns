package factory

// PizzaStore Интерфейс пиццерии
type PizzaStore interface {
	OrderPizza(string) SimplePizza
}

// Абстрактная пиццерия
type abstractPizzaStore struct {
	createPizza func(string) SimplePizza
}

// OrderPizza Заказать пиццу.
func (a *abstractPizzaStore) OrderPizza(pizzaType string) SimplePizza {
	pizza := a.createPizza(pizzaType)

	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()

	return pizza
}
