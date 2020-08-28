package simplefactory

const (
	nyDough     = "Thin Crust Dough"
	nySauce     = "Marinara Sauce"
	nySliceType = "diagonal"
)

// Сырная пицца в Нью-Йоркском стиле.
type nyCheesePizza struct {
	abstractPizza
}

// Создать сырную пиццу в Нью-Йоркском стиле.
func newNYCheesePizza() Pizza {
	return &nyCheesePizza{
		abstractPizza{
			name:  "NY Style Sauce and Cheese Pizza",
			dough: nyDough,
			sauce: nySauce,
			toppings: []string{
				"Grated Reggiano Cheese",
			},
			sliceType: nySliceType,
		},
	}
}

// Пицца "Пепперони" в Нью-Йоркском стиле.
type nyPepperoniPizza struct {
	abstractPizza
}

// Создать пиццу "Пепперони" в Нью-Йоркском стиле.
func newNYPepperoniPizza() Pizza {
	return &nyPepperoniPizza{
		abstractPizza{
			name:  "NY Style Sauce and Pepperoni Pizza",
			dough: nyDough,
			sauce: nySauce,
			toppings: []string{
				"Grated Reggiano Cheese",
				"Onions",
				"Peppers",
			},
			sliceType: nySliceType,
		},
	}
}

// Пицца с мидиями в Нью-Йоркском стиле.
type nyClamPizza struct {
	abstractPizza
}

// Создать пиццу с мидиями в Нью-Йоркском стиле.
func newNYClamPizza() Pizza {
	return &nyClamPizza{
		abstractPizza{
			name:  "NY Style Sauce and Clam Pizza",
			dough: nyDough,
			sauce: nySauce,
			toppings: []string{
				"Grated Reggiano Cheese",
				"Clams",
			},
			sliceType: nySliceType,
		},
	}
}

// Вегетарианская пицца в Нью-Йоркском стиле.
type nyVeggiePizza struct {
	abstractPizza
}

// Создать вегетарианскую пиццу в Нью-Йоркском стиле.
func newNYVeggiePizza() Pizza {
	return &nyVeggiePizza{
		abstractPizza{
			name:  "NY Style Sauce and Veggie Pizza",
			dough: nyDough,
			sauce: nySauce,
			toppings: []string{
				"Grated Parmesan Cheese",
				"Red Peppers",
				"Dried Oregano",
			},
			sliceType: nySliceType,
		},
	}
}
