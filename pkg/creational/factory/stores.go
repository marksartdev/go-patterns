package factory

// PizzaStore Интерфейс пиццерии
type PizzaStore interface {
	OrderPizza(int) Pizza
}

// Пиццерия
type pizzaStore struct {
	factory SimplePizzaFactory
}

// OrderPizza Заказать пиццу
func (p *pizzaStore) OrderPizza(pizzaType int) Pizza {
	pizza := p.factory.createPizza(pizzaType)
	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()

	return pizza
}

// NewPizzaStore Создать пиццерию
func NewPizzaStore(factory SimplePizzaFactory) PizzaStore {
	store := new(pizzaStore)
	store.factory = factory

	return store
}
