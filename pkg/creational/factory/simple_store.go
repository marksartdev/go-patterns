package factory

// Простая пиццерия.
type simplePizzaStore struct {
	factory SimplePizzaFactory
}

// OrderPizza Заказать пиццу.
func (p *simplePizzaStore) OrderPizza(pizzaType string) Pizza {
	pizza := p.factory.createPizza(pizzaType)
	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()

	return pizza
}

// NewSimplePizzaStore Создать простую пиццерию.
func NewSimplePizzaStore(factory SimplePizzaFactory) PizzaStore {
	store := new(simplePizzaStore)
	store.factory = factory

	return store
}
