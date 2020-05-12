package simplefactory

const (
	chicagoDough     = "Extra Thick Crust Dough"
	chicagoSauce     = "Plum Tomato Sauce"
	chicagoSliceType = "square"
)

// Сырная пицца в Чикагском стиле.
type chicagoCheesePizza struct {
	abstractPizza
}

// Создать сырную пиццу в Чикагском стиле.
func newChicagoCheesePizza() Pizza {
	pizza := new(chicagoCheesePizza)
	pizza.name = "Chicago Style Deep Dish Cheese Pizza"
	pizza.dough = chicagoDough
	pizza.sauce = chicagoSauce

	pizza.toppings = append(pizza.toppings, "Shredded Mozzarella Cheese")

	pizza.sliceType = chicagoSliceType

	return pizza
}

// Пицца "Пепперони" в Чикагском стиле.
type chicagoPepperoniPizza struct {
	abstractPizza
}

// Создать пиццу "Пепперони" в Чикагском стиле.
func newChicagoPepperoniPizza() Pizza {
	pizza := new(chicagoPepperoniPizza)
	pizza.name = "Chicago Style Deep Dish Pepperoni Pizza"
	pizza.dough = chicagoDough
	pizza.sauce = chicagoSauce

	pizza.toppings = append(pizza.toppings, "Shredded Mozzarella Cheese")
	pizza.toppings = append(pizza.toppings, "Onions")
	pizza.toppings = append(pizza.toppings, "Peppers")

	pizza.sliceType = chicagoSliceType

	return pizza
}

// Пицца с мидиями в Чикагском стиле.
type chicagoClamPizza struct {
	abstractPizza
}

// Создать пиццу с мидиями в Чикагском стиле.
func newChicagoClamPizza() Pizza {
	pizza := new(chicagoClamPizza)
	pizza.name = "Chicago Style Deep Dish Clam Pizza"
	pizza.dough = chicagoDough
	pizza.sauce = chicagoSauce

	pizza.toppings = append(pizza.toppings, "Shredded Mozzarella Cheese")
	pizza.toppings = append(pizza.toppings, "Clams")

	pizza.sliceType = chicagoSliceType

	return pizza
}

// Вегетарианская пицца в Чикагском стиле.
type chicagoVeggiePizza struct {
	abstractPizza
}

// Создать вегетарианскую пиццу в Чикагском стиле.
func newChicagoVeggiePizza() Pizza {
	pizza := new(chicagoVeggiePizza)
	pizza.name = "Chicago Style Deep Dish Veggie Pizza"
	pizza.dough = chicagoDough
	pizza.sauce = chicagoSauce

	pizza.toppings = append(pizza.toppings, "Shredded Mozzarella Cheese")
	pizza.toppings = append(pizza.toppings, "Red Peppers")
	pizza.toppings = append(pizza.toppings, "Self-Rising Flour")
	pizza.toppings = append(pizza.toppings, "Dried Oregano")

	pizza.sliceType = chicagoSliceType

	return pizza
}
