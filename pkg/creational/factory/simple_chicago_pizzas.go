package factory

const chicagoPizzaStyle = "Чикаго"

// Сырная пицца в Чикагском стиле.
type simpleChicagoCheesePizza struct {
	simpleAbstractPizza
}

// Создать сырную пиццу в Чикагском стиле.
func newSimpleChicagoCheesePizza() SimplePizza {
	pizza := new(simpleChicagoCheesePizza)
	pizza.Name = "Сырная"
	pizza.Style = chicagoPizzaStyle

	return pizza
}

// Пицца "Пепперони" в Чикагском стиле.
type simpleChicagoPepperoniPizza struct {
	simpleAbstractPizza
}

// Создать пиццу "Пепперони" в Чикагском стиле.
func newSimpleChicagoPepperoniPizza() SimplePizza {
	pizza := new(simpleChicagoPepperoniPizza)
	pizza.Name = "Пепперони"
	pizza.Style = chicagoPizzaStyle

	return pizza
}

// Пицца с мидиями в Чикагском стиле.
type simpleChicagoClamPizza struct {
	simpleAbstractPizza
}

// Создать пиццу с мидиями в Чикагском стиле.
func newSimpleChicagoClamPizza() SimplePizza {
	pizza := new(simpleChicagoClamPizza)
	pizza.Name = "С мидиями"
	pizza.Style = chicagoPizzaStyle

	return pizza
}

// Вегетарианская пицца в Чикагском стиле.
type simpleChicagoVeggiePizza struct {
	simpleAbstractPizza
}

// Создать вегетарианскую пиццу в Чикагском стиле.
func newSimpleChicagoVeggiePizza() SimplePizza {
	pizza := new(simpleChicagoVeggiePizza)
	pizza.Name = "Вегетарианская"
	pizza.Style = chicagoPizzaStyle

	return pizza
}
