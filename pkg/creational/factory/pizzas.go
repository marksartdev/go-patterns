package factory

import "fmt"

// Сырная пицца.
type cheesePizza struct {
	abstractPizza
}

// Создать сырную пиццу.
func newCheesePizza(ingredientFactory pizzaIngredientFactory, sliceType string) Pizza {
	return cheesePizza{
		abstractPizza{
			sliceType: sliceType,
			abstractPrepare: func(pizza abstractPizza) abstractPizza {
				pizza.log = append(pizza.log, fmt.Sprintf("Preparing %s", pizza.name))

				pizza.dough = ingredientFactory.createDough()
				pizza.sauce = ingredientFactory.createSauce()
				pizza.cheese = ingredientFactory.createCheese()
				pizza.veggies = ingredientFactory.createVeggies(CheesePizza)

				return pizza
			},
		},
	}
}

// Пицца "Пепперони".
type pepperoniPizza struct {
	abstractPizza
}

// Создать пиццу "Пепперони".
func newPepperoniPizza(ingredientFactory pizzaIngredientFactory, sliceType string) Pizza {
	return pepperoniPizza{
		abstractPizza{
			sliceType: sliceType,
			abstractPrepare: func(pizza abstractPizza) abstractPizza {
				pizza.log = append(pizza.log, fmt.Sprintf("Preparing %s", pizza.name))

				pizza.dough = ingredientFactory.createDough()
				pizza.sauce = ingredientFactory.createSauce()
				pizza.cheese = ingredientFactory.createCheese()
				pizza.veggies = ingredientFactory.createVeggies(PepperoniPizza)
				pizza.pepperoni = ingredientFactory.createPepperoni()

				return pizza
			},
		},
	}
}

// Пицца с мидиями.
type clamPizza struct {
	abstractPizza
}

// Создать пиццу с мидиями.
func newClamPizza(ingredientFactory pizzaIngredientFactory, sliceType string) Pizza {
	return clamPizza{
		abstractPizza{
			sliceType: sliceType,
			abstractPrepare: func(pizza abstractPizza) abstractPizza {
				pizza.log = append(pizza.log, fmt.Sprintf("Preparing %s", pizza.name))

				pizza.dough = ingredientFactory.createDough()
				pizza.sauce = ingredientFactory.createSauce()
				pizza.cheese = ingredientFactory.createCheese()
				pizza.clams = ingredientFactory.createClam()

				return pizza
			},
		},
	}
}

// Вегетарианская пицца.
type veggiePizza struct {
	abstractPizza
}

// Создать вегетарианскую пиццу.
func newVeggiePizza(ingredientFactory pizzaIngredientFactory, sliceType string) Pizza {
	return veggiePizza{
		abstractPizza{
			sliceType: sliceType,
			abstractPrepare: func(pizza abstractPizza) abstractPizza {
				pizza.log = append(pizza.log, fmt.Sprintf("Preparing %s", pizza.name))

				pizza.dough = ingredientFactory.createDough()
				pizza.sauce = ingredientFactory.createSauce()
				pizza.cheese = ingredientFactory.createCheese()
				pizza.veggies = ingredientFactory.createVeggies(VeggiePizza)

				return pizza
			},
		},
	}
}
