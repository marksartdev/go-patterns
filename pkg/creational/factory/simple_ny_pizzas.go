package factory

const nyPizzaStyle = "Нью-Йорк"

// Сырная пицца в Нью-Йоркском стиле.
type simpleNYCheesePizza struct {
	simpleAbstractPizza
}

// Создать сырную пиццу в Нью-Йоркском стиле.
func newSimpleNYCheesePizza() SimplePizza {
	pizza := new(simpleNYCheesePizza)
	pizza.Name = "Сырная"
	pizza.Style = nyPizzaStyle

	return pizza
}

// Пицца "Пепперони" в Нью-Йоркском стиле.
type simpleNYPepperoniPizza struct {
	simpleAbstractPizza
}

// Создать пиццу "Пепперони" в Нью-Йоркском стиле.
func newSimpleNYPepperoniPizza() SimplePizza {
	pizza := new(simpleNYPepperoniPizza)
	pizza.Name = "Пепперони"
	pizza.Style = nyPizzaStyle

	return pizza
}

// Пицца с мидиями в Нью-Йоркском стиле.
type simpleNYClamPizza struct {
	simpleAbstractPizza
}

// Создать пиццу с мидиями в Нью-Йоркском стиле.
func newSimpleNYClamPizza() SimplePizza {
	pizza := new(simpleNYClamPizza)
	pizza.Name = "С мидиями"
	pizza.Style = nyPizzaStyle

	return pizza
}

// Вегетарианская пицца в Нью-Йоркском стиле.
type simpleNYVeggiePizza struct {
	simpleAbstractPizza
}

// Создать вегетарианскую пиццу в Нью-Йоркском стиле.
func newSimpleNYVeggiePizza() SimplePizza {
	pizza := new(simpleNYVeggiePizza)
	pizza.Name = "Вегетарианская"
	pizza.Style = nyPizzaStyle

	return pizza
}
