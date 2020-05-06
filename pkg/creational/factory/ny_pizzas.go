package factory

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
	pizza := new(nyCheesePizza)
	pizza.name = "NY Style Sauce and Cheese Pizza"
	pizza.dough = nyDough
	pizza.sauce = nySauce

	pizza.toppings = append(pizza.toppings, "Grated Reggiano Cheese")

	pizza.sliceType = nySliceType

	return pizza
}

// Пицца "Пепперони" в Нью-Йоркском стиле.
type nyPepperoniPizza struct {
	abstractPizza
}

// Создать пиццу "Пепперони" в Нью-Йоркском стиле.
func newNYPepperoniPizza() Pizza {
	pizza := new(nyPepperoniPizza)
	pizza.name = "NY Style Sauce and Pepperoni Pizza"
	pizza.dough = nyDough
	pizza.sauce = nySauce

	pizza.toppings = append(pizza.toppings, "Grated Reggiano Cheese")
	pizza.toppings = append(pizza.toppings, "Onions")
	pizza.toppings = append(pizza.toppings, "Peppers")

	pizza.sliceType = nySliceType

	return pizza
}

// Пицца с мидиями в Нью-Йоркском стиле.
type nyClamPizza struct {
	abstractPizza
}

// Создать пиццу с мидиями в Нью-Йоркском стиле.
func newNYClamPizza() Pizza {
	pizza := new(nyClamPizza)
	pizza.name = "NY Style Sauce and Calm Pizza"
	pizza.dough = nyDough
	pizza.sauce = nySauce

	pizza.toppings = append(pizza.toppings, "Grated Reggiano Cheese")
	pizza.toppings = append(pizza.toppings, "Calms")

	pizza.sliceType = nySliceType

	return pizza
}

// Вегетарианская пицца в Нью-Йоркском стиле.
type nyVeggiePizza struct {
	abstractPizza
}

// Создать вегетарианскую пиццу в Нью-Йоркском стиле.
func newSimpleNYVeggiePizza() Pizza {
	pizza := new(nyVeggiePizza)
	pizza.name = "NY Style Sauce and Veggie Pizza"
	pizza.dough = nyDough
	pizza.sauce = nySauce

	pizza.toppings = append(pizza.toppings, "Grated Parmesan Cheese")
	pizza.toppings = append(pizza.toppings, "Red Peppers")
	pizza.toppings = append(pizza.toppings, "Dried Oregano")

	pizza.sliceType = nySliceType

	return pizza
}
