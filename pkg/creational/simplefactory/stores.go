package simplefactory

// PizzaStore Интерфейс пиццерии.
type PizzaStore interface {
	OrderPizza(string) (Pizza, error)
}

// Простая пиццерия.
type pizzaStore struct {
	factory PizzaFactory
}

// OrderPizza Заказать пиццу.
func (p *pizzaStore) OrderPizza(pizzaType string) (Pizza, error) {
	pizza, err := p.factory.createPizza(pizzaType)
	if err != nil {
		return nil, err
	}

	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()

	return pizza, nil
}

// NewPizzaStore Создать пиццерию.
func NewPizzaStore(factory PizzaFactory) PizzaStore {
	store := new(pizzaStore)
	store.factory = factory

	return store
}
