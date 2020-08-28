package factory

// Нью-Йоркская фабрика ингредиентов.
type chicagoPizzaIngredientFactory struct{}

func (c chicagoPizzaIngredientFactory) createDough() dough {
	return newThickCrustDough()
}

func (c chicagoPizzaIngredientFactory) createSauce() sauce {
	return newPlumTomatoSauce()
}

func (c chicagoPizzaIngredientFactory) createCheese() []cheese {
	return []cheese{newMozzarellaCheese(), newParmesanCheese()}
}

func (c chicagoPizzaIngredientFactory) createVeggies(pizzaType string) []veggie {
	switch pizzaType {
	case CheesePizza:
		return []veggie{newOregano()}
	case PepperoniPizza:
		return []veggie{newEggplant(), newSpinach(), newOlives()}
	case VeggiePizza:
		return []veggie{newEggplant(), newSpinach(), newOlives()}
	default:
		return []veggie{}
	}
}

func (c chicagoPizzaIngredientFactory) createPepperoni() pepperoni {
	return newSlicedPepperoni()
}

func (c chicagoPizzaIngredientFactory) createClam() clams {
	return newFrozenClams()
}
