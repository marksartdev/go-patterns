package factory

const (
	// CHEESE Сырная
	CHEESE = 0

	// PEPPERONI Пепперони
	PEPPERONI = 1

	// CLAM С мидиями
	CLAM = 2

	// VEGGIE Вегетарианская
	VEGGIE = 3
)

// SimplePizzaFactory Интерфейс простой фабрики по изготовлению пиццы
type SimplePizzaFactory interface {
	createPizza(int) Pizza
}

// Простая фабрика по изготовлению пиццы
type customPizzaFactory struct{}

// Создать пиццу
func (c *customPizzaFactory) createPizza(pizzaType int) Pizza {
	var pizza Pizza

	switch pizzaType {
	default:
		fallthrough
	case 0:
		pizza = newCheesePizza()

	case 1:
		pizza = newPepperoniPizza()

	case 2:
		pizza = newClamPizza()

	case 3:
		pizza = newVeggiePizza()
	}

	return pizza
}

// NewCustomPizzaFactory Создать простую фабрику по изготовлению пиццы
func NewCustomPizzaFactory() SimplePizzaFactory {
	return new(customPizzaFactory)
}
