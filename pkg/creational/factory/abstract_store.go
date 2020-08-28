// Package factory Паттерны "Фабричный метод и Абстрактная фабрика".
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

	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()

	return pizza, nil
}
