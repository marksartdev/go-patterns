package factory

// SimplePizzaFactory Интерфейс простой фабрики по изготовлению пиццы
type SimplePizzaFactory interface {
	createPizza(string) SimplePizza
}

// Простая фабрика по изготовлению пиццы в Нью-Йоркском стиле
type simpleNYPizzaFactory struct{}

// Создать пиццу.
func (s *simpleNYPizzaFactory) createPizza(pizzaType string) SimplePizza {
	var pizza SimplePizza

	switch pizzaType {
	case Cheese:
		pizza = newSimpleNYCheesePizza()

	case Pepperoni:
		pizza = newSimpleNYPepperoniPizza()

	case Clam:
		pizza = newSimpleNYClamPizza()

	case Veggie:
		pizza = newSimpleNYVeggiePizza()

	default:
		pizza = nil
	}

	return pizza
}

// NewSimpleNYPizzaFactory Создать простую фабрику по изготовлению пиццы в Нью-Йоркском стиле.
func NewSimpleNYPizzaFactory() SimplePizzaFactory {
	return new(simpleNYPizzaFactory)
}

// Простая фабрика по изготовлению пиццы в Нью-Йоркском стиле
type simpleChicagoPizzaFactory struct{}

// Создать пиццу.
func (s *simpleChicagoPizzaFactory) createPizza(pizzaType string) SimplePizza {
	var pizza SimplePizza

	switch pizzaType {
	case Cheese:
		pizza = newSimpleChicagoCheesePizza()

	case Pepperoni:
		pizza = newSimpleChicagoPepperoniPizza()

	case Clam:
		pizza = newSimpleChicagoClamPizza()

	case Veggie:
		pizza = newSimpleChicagoVeggiePizza()

	default:
		pizza = nil
	}

	return pizza
}

// NewSimpleChicagoPizzaFactory Создать простую фабрику по изготовлению пиццы в Чикагском стиле.
func NewSimpleChicagoPizzaFactory() SimplePizzaFactory {
	return new(simpleChicagoPizzaFactory)
}
