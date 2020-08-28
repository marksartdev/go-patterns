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
	return []cheese{newReggianoCheese()}
}

func (c nyPizzaIngredientFactory) createVeggies(pizzaType string) []veggie {
	switch pizzaType {
	case CheesePizza:
		return []veggie{newGarlic()}
	case PepperoniPizza:
		return []veggie{newMushroom(), newOnion(), newRedPepper()}
	case VeggiePizza:
		return []veggie{newMushroom(), newOnion(), newRedPepper()}
	default:
		return []veggie{}
	}
}

func (c nyPizzaIngredientFactory) createPepperoni() pepperoni {
	return newSlicedPepperoni()
}

func (c nyPizzaIngredientFactory) createClam() clams {
	return newFreshClams()
}
