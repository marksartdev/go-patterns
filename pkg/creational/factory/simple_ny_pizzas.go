package factory

const nyPizzaStyle = "Нью-Йорк"

// Сырная пицца в Нью-Йоркском стиле
type simpleNYStyleCheesePizza struct {
	simplePizza
}

// Создать сырную пиццу в Нью-Йоркском стиле
func newSimpleNYStyleCheesePizza() SimplePizza {
	pizza := new(simpleNYStyleCheesePizza)
	pizza.Name = "Сырная"
	pizza.Style = nyPizzaStyle

	return pizza
}

// Пицца "Пепперони" в Нью-Йоркском стиле
type simpleNYStylePepperoniPizza struct {
	simplePizza
}

// Создать пиццу "Пепперони" в Нью-Йоркском стиле
func newSimpleNYStylePepperoniPizza() SimplePizza {
	pizza := new(simpleNYStylePepperoniPizza)
	pizza.Name = "Пепперони"
	pizza.Style = nyPizzaStyle

	return pizza
}

// Пицца с мидиями в Нью-Йоркском стиле
type simpleNYStyleClamPizza struct {
	simplePizza
}

// Создать пиццу с мидиями в Нью-Йоркском стиле
func newSimpleNYStyleClamPizza() SimplePizza {
	pizza := new(simpleNYStyleClamPizza)
	pizza.Name = "С мидиями"
	pizza.Style = nyPizzaStyle

	return pizza
}

// Вегетарианская пицца в Нью-Йоркском стиле
type simpleNYStyleVeggiePizza struct {
	simplePizza
}

// Создать вегетарианскую пиццу в Нью-Йоркском стиле
func newSimpleNYStyleVeggiePizza() SimplePizza {
	pizza := new(simpleNYStyleVeggiePizza)
	pizza.Name = "Вегетарианская"
	pizza.Style = nyPizzaStyle

	return pizza
}
