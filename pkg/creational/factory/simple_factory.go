package factory

// SimplePizzaFactory Интерфейс простой фабрики по изготовлению пиццы.
type SimplePizzaFactory interface {
	createPizza(string) Pizza
}

// Простая фабрика по изготовлению пиццы в Нью-Йоркском стиле.
type simpleNYPizzaFactory struct{}

// Создать пиццу.
func (s *simpleNYPizzaFactory) createPizza(pizzaType string) Pizza {
	var pizza Pizza

	switch pizzaType {
	case Cheese:
		pizza = newNYCheesePizza()

	case Pepperoni:
		pizza = newNYPepperoniPizza()

	case Clam:
		pizza = newNYClamPizza()

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

// Простая фабрика по изготовлению пиццы в Нью-Йоркском стиле.
type simpleChicagoPizzaFactory struct{}

// Создать пиццу.
func (s *simpleChicagoPizzaFactory) createPizza(pizzaType string) Pizza {
	var pizza Pizza

	switch pizzaType {
	case Cheese:
		pizza = newChicagoCheesePizza()

	case Pepperoni:
		pizza = newChicagoPepperoniPizza()

	case Clam:
		pizza = newChicagoClamPizza()

	case Veggie:
		pizza = newChicagoVeggiePizza()

	default:
		pizza = nil
	}

	return pizza
}

// NewSimpleChicagoPizzaFactory Создать простую фабрику по изготовлению пиццы в Чикагском стиле.
func NewSimpleChicagoPizzaFactory() SimplePizzaFactory {
	return new(simpleChicagoPizzaFactory)
}
