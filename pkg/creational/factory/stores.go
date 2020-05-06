package factory

// Пиццерия в Нью-Йоркском стиле.
type nyPizzaStore struct {
	abstractPizzaStore
}

// Создать пиццу в Нью-Йоркском стиле.
func createNYPizza(pizzaType string) Pizza {
	switch pizzaType {
	case Cheese:
		return newNYCheesePizza()
	case Pepperoni:
		return newNYPepperoniPizza()
	case Clam:
		return newNYClamPizza()
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
func createChicagoPizza(pizzaType string) Pizza {
	switch pizzaType {
	case Cheese:
		return newChicagoCheesePizza()
	case Pepperoni:
		return newChicagoPepperoniPizza()
	case Clam:
		return newChicagoClamPizza()
	case Veggie:
		return newChicagoVeggiePizza()
	}

	return nil
}

// NewChicagoPizzaStore Создать пиццерию в Чикагском стиле.
func NewChicagoPizzaStore() PizzaStore {
	store := new(chicagoPizzaStore)
	store.createPizza = createChicagoPizza

	return store
}
