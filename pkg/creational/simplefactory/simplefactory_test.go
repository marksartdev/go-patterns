package simplefactory_test

import (
	"fmt"
	"testing"

	"github.com/Mark-Sart/go-patterns/pkg/creational/simplefactory"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	factory   simplefactory.PizzaFactory
	pizzaType string
	name      string
	logs      []string
	err       string
}

func TestPizzaStore_OrderPizza_NY_CheesePepperoni(t *testing.T) {
	pizzaFactory := simplefactory.NewNYPizzaFactory()

	cases := []testCase{
		{
			pizzaFactory,
			simplefactory.CheesePizza,
			"NY Style Sauce and Cheese Pizza",
			[]string{
				"Preparing NY Style Sauce and Cheese Pizza",
				"Tossing dough... Thin Crust Dough",
				"Adding sauce... Marinara Sauce",
				"Adding toppings:",
				"    Grated Reggiano Cheese",
				"Bake for 25 minutes at 350",
				"Cutting the pizza into diagonal slices",
				"Place pizza in official PizzaStore box",
			},
			"",
		},
		{
			pizzaFactory,
			simplefactory.PepperoniPizza,
			"NY Style Sauce and Pepperoni Pizza",
			[]string{
				"Preparing NY Style Sauce and Pepperoni Pizza",
				"Tossing dough... Thin Crust Dough",
				"Adding sauce... Marinara Sauce",
				"Adding toppings:",
				"    Grated Reggiano Cheese",
				"    Onions",
				"    Peppers",
				"Bake for 25 minutes at 350",
				"Cutting the pizza into diagonal slices",
				"Place pizza in official PizzaStore box",
			},
			"",
		},
	}

	testCases(t, cases)
}

func TestPizzaStore_OrderPizza_NY_ClamVeggie(t *testing.T) {
	pizzaFactory := simplefactory.NewNYPizzaFactory()

	cases := []testCase{
		{
			pizzaFactory,
			simplefactory.ClamPizza,
			"NY Style Sauce and Clam Pizza",
			[]string{
				"Preparing NY Style Sauce and Clam Pizza",
				"Tossing dough... Thin Crust Dough",
				"Adding sauce... Marinara Sauce",
				"Adding toppings:",
				"    Grated Reggiano Cheese",
				"    Clams",
				"Bake for 25 minutes at 350",
				"Cutting the pizza into diagonal slices",
				"Place pizza in official PizzaStore box",
			},
			"",
		},
		{
			pizzaFactory,
			simplefactory.VeggiePizza,
			"NY Style Sauce and Veggie Pizza",
			[]string{
				"Preparing NY Style Sauce and Veggie Pizza",
				"Tossing dough... Thin Crust Dough",
				"Adding sauce... Marinara Sauce",
				"Adding toppings:",
				"    Grated Parmesan Cheese",
				"    Red Peppers",
				"    Dried Oregano",
				"Bake for 25 minutes at 350",
				"Cutting the pizza into diagonal slices",
				"Place pizza in official PizzaStore box",
			},
			"",
		},
	}

	testCases(t, cases)
}

func TestPizzaStore_OrderPizza_Chicago_CheesePepperoni(t *testing.T) {
	pizzaFactory := simplefactory.NewChicagoPizzaFactory()

	cases := []testCase{
		{
			pizzaFactory,
			simplefactory.CheesePizza,
			"Chicago Style Deep Dish Cheese Pizza",
			[]string{
				"Preparing Chicago Style Deep Dish Cheese Pizza",
				"Tossing dough... Extra Thick Crust Dough",
				"Adding sauce... Plum Tomato Sauce",
				"Adding toppings:",
				"    Shredded Mozzarella Cheese",
				"Bake for 25 minutes at 350",
				"Cutting the pizza into square slices",
				"Place pizza in official PizzaStore box",
			},
			"",
		},
		{
			pizzaFactory,
			simplefactory.PepperoniPizza,
			"Chicago Style Deep Dish Pepperoni Pizza",
			[]string{
				"Preparing Chicago Style Deep Dish Pepperoni Pizza",
				"Tossing dough... Extra Thick Crust Dough",
				"Adding sauce... Plum Tomato Sauce",
				"Adding toppings:",
				"    Shredded Mozzarella Cheese",
				"    Onions",
				"    Peppers",
				"Bake for 25 minutes at 350",
				"Cutting the pizza into square slices",
				"Place pizza in official PizzaStore box",
			},
			"",
		},
	}

	testCases(t, cases)
}

func TestPizzaStore_OrderPizza_Chicago_ClamVeggie(t *testing.T) {
	pizzaFactory := simplefactory.NewChicagoPizzaFactory()

	cases := []testCase{
		{
			pizzaFactory,
			simplefactory.ClamPizza,
			"Chicago Style Deep Dish Clam Pizza",
			[]string{
				"Preparing Chicago Style Deep Dish Clam Pizza",
				"Tossing dough... Extra Thick Crust Dough",
				"Adding sauce... Plum Tomato Sauce",
				"Adding toppings:",
				"    Shredded Mozzarella Cheese",
				"    Clams",
				"Bake for 25 minutes at 350",
				"Cutting the pizza into square slices",
				"Place pizza in official PizzaStore box",
			},
			"",
		},
		{
			pizzaFactory,
			simplefactory.VeggiePizza,
			"Chicago Style Deep Dish Veggie Pizza",
			[]string{
				"Preparing Chicago Style Deep Dish Veggie Pizza",
				"Tossing dough... Extra Thick Crust Dough",
				"Adding sauce... Plum Tomato Sauce",
				"Adding toppings:",
				"    Shredded Mozzarella Cheese",
				"    Red Peppers",
				"    Self-Rising Flour",
				"    Dried Oregano",
				"Bake for 25 minutes at 350",
				"Cutting the pizza into square slices",
				"Place pizza in official PizzaStore box",
			},
			"",
		},
	}

	testCases(t, cases)
}

func TestPizzaStore_OrderPizza_Error(t *testing.T) {
	cases := []testCase{
		{
			simplefactory.NewNYPizzaFactory(),
			"New-York",
			"",
			nil,
			"this factory can't create New-York pizza",
		},
		{
			simplefactory.NewChicagoPizzaFactory(),
			"Chicago",
			"",
			nil,
			"this factory can't create Chicago pizza",
		},
	}

	testCases(t, cases)
}

func testCases(t *testing.T, cases []testCase) {
	for i, item := range cases {
		pizzaStore := simplefactory.NewPizzaStore(item.factory)
		pizza, err := pizzaStore.OrderPizza(item.pizzaType)

		if err != nil {
			assert.Equal(t, item.err, err.Error(), fmt.Sprintf("case %d", i))
		} else {
			assert.Equal(t, item.name, pizza.GetName(), fmt.Sprintf("case %d", i))
			assert.Equal(t, item.logs, pizza.GetLog(), fmt.Sprintf("case %d", i))
		}
	}
}
