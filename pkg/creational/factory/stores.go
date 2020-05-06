package factory

// Пиццерия в Нью-Йоркском стиле.
type nyPizzaStore struct {
	abstractPizzaStore
}

// Создать пиццу в Нью-Йоркском стиле.
func createNYPizza(pizzaType string) SimplePizza {
	switch pizzaType {
	case Cheese:
		return newSimpleNYCheesePizza()
	case Pepperoni:
		return newSimpleNYPepperoniPizza()
	case Clam:
		return newSimpleNYClamPizza()
	case Veggie:
		return newSimpleNYVeggiePizza()
	}

	return nil
}

// NewNYPizzaStore Создать пиццерию в Нью-Йоркском стиле.
func NewNYPizzaStore() PizzaStore {
	store := new(nyPizzaStore)
	store.createPizza = createNYPizza

	return store
}

// Пиццерия в Чикагском стиле.
type chicagoPizzaStore struct {
	abstractPizzaStore
}

// Создать пиццу в Чикагском стиле.
func createChicagoPizza(pizzaType string) SimplePizza {
	switch pizzaType {
	case Cheese:
		return newSimpleChicagoCheesePizza()
	case Pepperoni:
		return newSimpleChicagoPepperoniPizza()
	case Clam:
		return newSimpleChicagoClamPizza()
	case Veggie:
		return newSimpleChicagoVeggiePizza()
	}

	return nil
}

// NewChicagoPizzaStore Создать пиццерию в Чикагском стиле.
func NewChicagoPizzaStore() PizzaStore {
	store := new(chicagoPizzaStore)
	store.createPizza = createChicagoPizza

	return store
}
