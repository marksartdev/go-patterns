package factory

// Нью-Йоркская фабрика ингредиентов.
type chicagoPizzaIngredientFactory struct{}

func (c *chicagoPizzaIngredientFactory) createDough() *dough {
	return newThickCrustDough()
}

func (c *chicagoPizzaIngredientFactory) createSauce() *sauce {
	return newPlumTomatoSauce()
}

func (c *chicagoPizzaIngredientFactory) createCheese() []*cheese {
	cheeses := make([]*cheese, 0, 2)
	cheeses = append(cheeses, newMozzarellaCheese())
	cheeses = append(cheeses, newParmesanCheese())

	return cheeses
}

func (c *chicagoPizzaIngredientFactory) createVeggies(pizzaType string) []*veggie {
	veggies := make([]*veggie, 0)

	switch pizzaType {
	case CheesePizza:
		veggies = append(veggies, newOregano())
	case PepperoniPizza:
		veggies = append(veggies, newEggplant())
		veggies = append(veggies, newSpinach())
		veggies = append(veggies, newOlives())
	case VeggiePizza:
		veggies = append(veggies, newEggplant())
		veggies = append(veggies, newSpinach())
		veggies = append(veggies, newOlives())
	}

	return veggies
}

func (c *chicagoPizzaIngredientFactory) createPepperoni() *pepperoni {
	return newSlicedPepperoni()
}

func (c *chicagoPizzaIngredientFactory) createClam() *clams {
	return newFrozenClams()
}
