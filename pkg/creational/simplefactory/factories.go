package simplefactory

// PizzaFactory Интерфейс простой фабрики по изготовлению пиццы.
type PizzaFactory interface {
	createPizza(string) (Pizza, error)
}

// Простая фабрика по изготовлению пиццы в Нью-Йоркском стиле.
type nyPizzaFactory struct{}

// Создать пиццу.
func (s nyPizzaFactory) createPizza(pizzaType string) (Pizza, error) {
	var (
		pizza Pizza
		err   error
	)

	switch pizzaType {
	case CheesePizza:
		pizza = newNYCheesePizza()

	case PepperoniPizza:
		pizza = newNYPepperoniPizza()

	case ClamPizza:
		pizza = newNYClamPizza()

	case VeggiePizza:
		pizza = newNYVeggiePizza()

	default:
		err = newPizzaTypeError(pizzaType)
	}

	return pizza, err
}

// NewNYPizzaFactory Создать простую фабрику по изготовлению пиццы в Нью-Йоркском стиле.
func NewNYPizzaFactory() PizzaFactory {
	return nyPizzaFactory{}
}

// Простая фабрика по изготовлению пиццы в Нью-Йоркском стиле.
type chicagoPizzaFactory struct{}

// Создать пиццу.
func (s chicagoPizzaFactory) createPizza(pizzaType string) (Pizza, error) {
	var (
		pizza Pizza
		err   error
	)

	switch pizzaType {
	case CheesePizza:
		pizza = newChicagoCheesePizza()

	case PepperoniPizza:
		pizza = newChicagoPepperoniPizza()

	case ClamPizza:
		pizza = newChicagoClamPizza()

	case VeggiePizza:
		pizza = newChicagoVeggiePizza()

	default:
		err = newPizzaTypeError(pizzaType)
	}

	return pizza, err
}

// NewChicagoPizzaFactory Создать простую фабрику по изготовлению пиццы в Чикагском стиле.
func NewChicagoPizzaFactory() PizzaFactory {
	return chicagoPizzaFactory{}
}
