package factory

import "fmt"

// Сырная пицца.
type cheesePizza struct {
	abstractPizza
}

// Создать сырную пиццу.
func newCheesePizza(ingredientFactory pizzaIngredientFactory, sliceType string) Pizza {
	pizza := new(cheesePizza)
	pizza.sliceType = sliceType

	pizza.abstractPrepare = func(a *abstractPizza) {
		a.log = append(a.log, fmt.Sprintf("Preparing %s", a.name))

		a.dough = ingredientFactory.createDough()
		a.sauce = ingredientFactory.createSauce()
		a.cheese = ingredientFactory.createCheese()
		a.veggies = ingredientFactory.createVeggies(CheesePizza)
	}

	return pizza
}

// Пицца "Пепперони".
type pepperoniPizza struct {
	abstractPizza
}

// Создать пиццу "Пепперони".
func newPepperoniPizza(ingredientFactory pizzaIngredientFactory, sliceType string) Pizza {
	pizza := new(pepperoniPizza)
	pizza.sliceType = sliceType

	pizza.abstractPrepare = func(a *abstractPizza) {
		a.log = append(a.log, fmt.Sprintf("Preparing %s", a.name))

		a.dough = ingredientFactory.createDough()
		a.sauce = ingredientFactory.createSauce()
		a.cheese = ingredientFactory.createCheese()
		a.veggies = ingredientFactory.createVeggies(PepperoniPizza)
		a.pepperoni = ingredientFactory.createPepperoni()
	}

	return pizza
}

// Пицца с мидиями.
type clamPizza struct {
	abstractPizza
}

// Создать пиццу с мидиями.
func newClamPizza(ingredientFactory pizzaIngredientFactory, sliceType string) Pizza {
	pizza := new(clamPizza)
	pizza.sliceType = sliceType

	pizza.abstractPrepare = func(a *abstractPizza) {
		a.log = append(a.log, fmt.Sprintf("Preparing %s", a.name))

		a.dough = ingredientFactory.createDough()
		a.sauce = ingredientFactory.createSauce()
		a.cheese = ingredientFactory.createCheese()
		a.clams = ingredientFactory.createClam()
	}

	return pizza
}

// Вегетарианская пицца.
type veggiePizza struct {
	abstractPizza
}

// Создать вегетарианскую пиццу.
func newVeggiePizza(ingredientFactory pizzaIngredientFactory, sliceType string) Pizza {
	pizza := new(veggiePizza)
	pizza.sliceType = sliceType

	pizza.abstractPrepare = func(a *abstractPizza) {
		a.log = append(a.log, fmt.Sprintf("Preparing %s", a.name))

		a.dough = ingredientFactory.createDough()
		a.sauce = ingredientFactory.createSauce()
		a.cheese = ingredientFactory.createCheese()
		a.veggies = ingredientFactory.createVeggies(VeggiePizza)
	}

	return pizza
}
