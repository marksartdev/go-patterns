package factory

// Пиццерия в Нью-Йоркском стиле.
type nyPizzaStore struct {
	abstractPizzaStore
}

// Создать пиццу в Нью-Йоркском стиле.
func createNYPizza(pizzaType string) Pizza {
	var pizza Pizza

	ingredientFactory := new(nyPizzaIngredientFactory)

	switch pizzaType {
	case CheesePizza:
		pizza = newCheesePizza(ingredientFactory, "diagonal")
		pizza.setName("New-York Style Cheese Pizza")

	case PepperoniPizza:
		pizza = newPepperoniPizza(ingredientFactory, "diagonal")
		pizza.setName("New-York Style Pepperoni Pizza")

	case ClamPizza:
		pizza = newClamPizza(ingredientFactory, "diagonal")
		pizza.setName("New-York Style Clam Pizza")

	case VeggiePizza:
		pizza = newVeggiePizza(ingredientFactory, "diagonal")
		pizza.setName("New-York Style Cheese Pizza")
	}

	return pizza
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
	var pizza Pizza

	ingredientFactory := new(chicagoPizzaIngredientFactory)

	switch pizzaType {
	case CheesePizza:
		pizza = newCheesePizza(ingredientFactory, "square")
		pizza.setName("Chicago Style Cheese Pizza")

	case PepperoniPizza:
		pizza = newPepperoniPizza(ingredientFactory, "square")
		pizza.setName("Chicago Style Pepperoni Pizza")

	case ClamPizza:
		pizza = newClamPizza(ingredientFactory, "square")
		pizza.setName("Chicago Style Clam Pizza")

	case VeggiePizza:
		pizza = newVeggiePizza(ingredientFactory, "square")
		pizza.setName("Chicago Style Cheese Pizza")
	}

	return pizza
}

// NewChicagoPizzaStore Создать пиццерию в Чикагском стиле.
func NewChicagoPizzaStore() PizzaStore {
	store := new(chicagoPizzaStore)
	store.createPizza = createChicagoPizza

	return store
}
