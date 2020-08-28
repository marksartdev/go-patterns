package simplefactory

const (
	chicagoDough     = "Extra Thick Crust Dough"
	chicagoSauce     = "Plum Tomato Sauce"
	chicagoSliceType = "square"
)

// Сырная пицца в Чикагском стиле.
type chicagoCheesePizza struct {
	abstractPizza
}

// Создать сырную пиццу в Чикагском стиле.
func newChicagoCheesePizza() Pizza {
	return &chicagoCheesePizza{
		abstractPizza{
			name:  "Chicago Style Deep Dish Cheese Pizza",
			dough: chicagoDough,
			sauce: chicagoSauce,
			toppings: []string{
				"Shredded Mozzarella Cheese",
			},
			sliceType: chicagoSliceType,
		},
	}
}

// Пицца "Пепперони" в Чикагском стиле.
type chicagoPepperoniPizza struct {
	abstractPizza
}

// Создать пиццу "Пепперони" в Чикагском стиле.
func newChicagoPepperoniPizza() Pizza {
	return &chicagoPepperoniPizza{
		abstractPizza{
			name:  "Chicago Style Deep Dish Pepperoni Pizza",
			dough: chicagoDough,
			sauce: chicagoSauce,
			toppings: []string{
				"Shredded Mozzarella Cheese",
				"Onions",
				"Peppers",
			},
			sliceType: chicagoSliceType,
		},
	}
}

// Пицца с мидиями в Чикагском стиле.
type chicagoClamPizza struct {
	abstractPizza
}

// Создать пиццу с мидиями в Чикагском стиле.
func newChicagoClamPizza() Pizza {
	return &chicagoClamPizza{
		abstractPizza{
			name:  "Chicago Style Deep Dish Clam Pizza",
			dough: chicagoDough,
			sauce: chicagoSauce,
			toppings: []string{
				"Shredded Mozzarella Cheese",
				"Clams",
			},
			sliceType: chicagoSliceType,
		},
	}
}

// Вегетарианская пицца в Чикагском стиле.
type chicagoVeggiePizza struct {
	abstractPizza
}

// Создать вегетарианскую пиццу в Чикагском стиле.
func newChicagoVeggiePizza() Pizza {
	return &chicagoVeggiePizza{
		abstractPizza{
			name:  "Chicago Style Deep Dish Veggie Pizza",
			dough: chicagoDough,
			sauce: chicagoSauce,
			toppings: []string{
				"Shredded Mozzarella Cheese",
				"Red Peppers",
				"Self-Rising Flour",
				"Dried Oregano",
			},
			sliceType: chicagoSliceType,
		},
	}
}
