package factory

// SimplePizzaStore Интерфейс простой пиццерии
type SimplePizzaStore interface {
	OrderPizza(string) SimplePizza
}

// Простая пиццерия
type simplePizzaStore struct {
	factory SimplePizzaFactory
}

// OrderPizza Заказать пиццу
func (p *simplePizzaStore) OrderPizza(pizzaType string) SimplePizza {
	pizza := p.factory.createPizza(pizzaType)
	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()

	return pizza
}

// NewSimplePizzaStore Создать простую пиццерию
func NewSimplePizzaStore(factory SimplePizzaFactory) SimplePizzaStore {
	store := new(simplePizzaStore)
	store.factory = factory

	return store
}
