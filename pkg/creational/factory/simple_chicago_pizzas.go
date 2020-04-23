package factory

// Сырная пицца в Чикагском стиле
type simpleChicagoStyleCheesePizza struct {
	simplePizza
}

// Создать сырную пиццу в Чикагском стиле
func newSimpleChicagoStyleCheesePizza() SimplePizza {
	pizza := new(simpleChicagoStyleCheesePizza)
	pizza.Name = "Сырная"
	pizza.Style = "Чикаго"

	return pizza
}

// Пицца "Пепперони" в Чикагском стиле
type simpleChicagoStylePepperoniPizza struct {
	simplePizza
}

// Создать пиццу "Пепперони" в Чикагском стиле
func newSimpleChicagoStylePepperoniPizza() SimplePizza {
	pizza := new(simpleChicagoStylePepperoniPizza)
	pizza.Name = "Пепперони"
	pizza.Style = "Чикаго"

	return pizza
}

// Пицца с мидиями в Чикагском стиле
type simpleChicagoStyleClamPizza struct {
	simplePizza
}

// Создать пиццу с мидиями в Чикагском стиле
func newSimpleChicagoStyleClamPizza() SimplePizza {
	pizza := new(simpleChicagoStyleClamPizza)
	pizza.Name = "С мидиями"
	pizza.Style = "Чикаго"

	return pizza
}

// Вегетарианская пицца в Чикагском стиле
type simpleChicagoStyleVeggiePizza struct {
	simplePizza
}

// Создать вегетарианскую пиццу в Чикагском стиле
func newSimpleChicagoStyleVeggiePizza() SimplePizza {
	pizza := new(simpleChicagoStyleVeggiePizza)
	pizza.Name = "Вегетарианская"
	pizza.Style = "Чикаго"

	return pizza
}
