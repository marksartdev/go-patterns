package factory

// Нью-Йоркская фабрика ингредиентов.
type nyPizzaIngredientFactory struct{}

func (c nyPizzaIngredientFactory) createDough() dough {
	return newThinCrustDough()
}

func (c nyPizzaIngredientFactory) createSauce() sauce {
	return newMarinaraSauce()
}

func (c nyPizzaIngredientFactory) createCheese() []cheese {
	cheeses := make([]cheese, 0, 1)
	cheeses = append(cheeses, newReggianoCheese())

	return cheeses
}

func (c nyPizzaIngredientFactory) createVeggies(pizzaType string) []veggie {
	veggies := make([]veggie, 0, 3)

	switch pizzaType {
	case CheesePizza:
		veggies = append(veggies, newGarlic())
	case PepperoniPizza:
		veggies = append(veggies, newMushroom())
		veggies = append(veggies, newOnion())
		veggies = append(veggies, newRedPepper())
	case VeggiePizza:
		veggies = append(veggies, newMushroom())
		veggies = append(veggies, newOnion())
		veggies = append(veggies, newRedPepper())
	}

	return veggies
}

func (c nyPizzaIngredientFactory) createPepperoni() pepperoni {
	return newSlicedPepperoni()
}

func (c nyPizzaIngredientFactory) createClam() clams {
	return newFreshClams()
}
