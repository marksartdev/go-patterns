package factory

// Интерфейс фабрики по производству ингредиентов.
type pizzaIngredientFactory interface {
	createDough() dough
	createSauce() sauce
	createCheese() []cheese
	createVeggies(pizzaType string) []veggie
	createPepperoni() pepperoni
	createClam() clams
}
