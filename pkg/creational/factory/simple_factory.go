package factory

// SimplePizzaFactory Интерфейс простой фабрики по изготовлению пиццы
type SimplePizzaFactory interface {
	createPizza(string) SimplePizza
}

// Простая фабрика по изготовлению пиццы в Нью-Йоркском стиле
type simpleNYPizzaFactory struct{}

// Создать пиццу
func (s *simpleNYPizzaFactory) createPizza(pizzaType string) SimplePizza {
	var pizza SimplePizza

	switch pizzaType {
	case "cheese":
		pizza = newSimpleNYStyleCheesePizza()

	case "pepperoni":
		pizza = newSimpleNYStylePepperoniPizza()

	case "clam":
		pizza = newSimpleNYStyleClamPizza()

	case "veggie":
		pizza = newSimpleNYStyleVeggiePizza()

	default:
		pizza = nil
	}

	return pizza
}

// NewSimpleNYPizzaFactory Создать простую фабрику по изготовлению пиццы
func NewSimpleNYPizzaFactory() SimplePizzaFactory {
	return new(simpleNYPizzaFactory)
}

// Простая фабрика по изготовлению пиццы в Нью-Йоркском стиле
type simpleChicagoPizzaFactory struct{}

// Создать пиццу
func (s *simpleChicagoPizzaFactory) createPizza(pizzaType string) SimplePizza {
	var pizza SimplePizza

	switch pizzaType {
	case "cheese":
		pizza = newSimpleChicagoStyleCheesePizza()

	case "pepperoni":
		pizza = newSimpleChicagoStylePepperoniPizza()

	case "clam":
		pizza = newSimpleChicagoStyleClamPizza()

	case "veggie":
		pizza = newSimpleChicagoStyleVeggiePizza()

	default:
		pizza = nil
	}

	return pizza
}

// NewSimpleChicagoPizzaFactory Создать простую фабрику по изготовлению пиццы
func NewSimpleChicagoPizzaFactory() SimplePizzaFactory {
	return new(simpleChicagoPizzaFactory)
}
