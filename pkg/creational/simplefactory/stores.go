package simplefactory

// PizzaStore Интерфейс пиццерии.
type PizzaStore interface {
	OrderPizza(pizzaType string) (Pizza, error)
}

// Простая пиццерия.
type pizzaStore struct {
	factory PizzaFactory
}

// OrderPizza Заказать пиццу.
func (p pizzaStore) OrderPizza(pizzaType string) (Pizza, error) {
	pizza, err := p.factory.createPizza(pizzaType)
	if err != nil {
		return nil, err
	}

	pizza = pizza.prepare()
	pizza = pizza.bake()
	pizza = pizza.cut()
	pizza = pizza.box()

	return pizza, nil
}

// NewPizzaStore Создать пиццерию.
func NewPizzaStore(factory PizzaFactory) PizzaStore {
	store := pizzaStore{factory: factory}

	return store
}
