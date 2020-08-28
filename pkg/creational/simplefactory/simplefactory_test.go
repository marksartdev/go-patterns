package simplefactory_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marksartdev/go-patterns/pkg/creational/simplefactory"
)

type testCase struct {
	pizzaType string
	name      string
	logs      []string
	err       string
}

func TestPizzaStore_OrderPizza_NY(t *testing.T) {
	pizzaFactory := simplefactory.NewNYPizzaFactory()
	cases := getNYTestCases1()
	cases = append(cases, getNYTestCases2()...)

	testCases(t, pizzaFactory, cases)
}

func getNYTestCases1() []testCase {
	return []testCase{
		{
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
}

func getNYTestCases2() []testCase {
	return []testCase{
		{
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
}

func TestPizzaStore_OrderPizza_Chicago(t *testing.T) {
	pizzaFactory := simplefactory.NewChicagoPizzaFactory()
	cases := getChicagoTestCases1()
	cases = append(cases, getChicagoTestCases2()...)

	testCases(t, pizzaFactory, cases)
}

func getChicagoTestCases1() []testCase {
	return []testCase{
		{
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
}

func getChicagoTestCases2() []testCase {
	return []testCase{
		{
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
}

func TestPizzaStore_OrderPizza_Error(t *testing.T) {
	pizzaFactory := simplefactory.NewNYPizzaFactory()
	cases := []testCase{
		{
			"New-York",
			"",
			nil,
			"this factory can't create New-York pizza",
		},
	}

	testCases(t, pizzaFactory, cases)

	pizzaFactory = simplefactory.NewChicagoPizzaFactory()
	cases = []testCase{
		{
			"Chicago",
			"",
			nil,
			"this factory can't create Chicago pizza",
		},
	}

	testCases(t, pizzaFactory, cases)
}

func testCases(t *testing.T, pizzaFactory simplefactory.PizzaFactory, cases []testCase) {
	for i := range cases {
		pizzaStore := simplefactory.NewPizzaStore(pizzaFactory)
		pizza, err := pizzaStore.OrderPizza(cases[i].pizzaType)

		if err != nil {
			assert.Equal(t, cases[i].err, err.Error(), fmt.Sprintf("case %d", i))
		} else {
			assert.Equal(t, cases[i].name, pizza.GetName(), fmt.Sprintf("case %d", i))
			assert.Equal(t, cases[i].logs, pizza.GetLog(), fmt.Sprintf("case %d", i))
		}
	}
}
