package simplefactory

// PizzaStore Интерфейс пиццерии.
type PizzaStore interface {
	OrderPizza(string) Pizza
}

// Простая пиццерия.
type pizzaStore struct {
	factory PizzaFactory
}

// OrderPizza Заказать пиццу.
func (p *pizzaStore) OrderPizza(pizzaType string) Pizza {
	pizza := p.factory.createPizza(pizzaType)
	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()

	return pizza
}

// NewPizzaStore Создать пиццерию.
func NewPizzaStore(factory PizzaFactory) PizzaStore {
	store := new(pizzaStore)
	store.factory = factory

	return store
}
