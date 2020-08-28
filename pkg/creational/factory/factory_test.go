package factory_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marksartdev/go-patterns/pkg/creational/factory"
)

type testCase struct {
	pizzaType string
	name      string
	logs      []string
	err       string
}

func TestNewNYPizzaStore(t *testing.T) {
	pizzaStore := factory.NewNYPizzaStore()
	cases := getNYTestCases1()
	cases = append(cases, getNYTestCases2()...)

	testCases(t, pizzaStore, cases)
}

func getNYTestCases1() []testCase {
	return []testCase{
		{
			factory.CheesePizza,
			"New-York Style Cheese Pizza",
			[]string{
				"Preparing New-York Style Cheese Pizza",
				"Tossing Thin Crust Dough",
				"Adding Marinara Sauce",
				"Adding cheese:",
				"    Grated Reggiano Cheese",
				"Adding toppings:",
				"    garlic",
				"Bake for 25 minutes at 350",
				"Cutting the pizza into diagonal slices",
				"Place pizza in official PizzaStore box",
			},
			"",
		},
		{
			factory.PepperoniPizza,
			"New-York Style Pepperoni Pizza",
			[]string{
				"Preparing New-York Style Pepperoni Pizza",
				"Tossing Thin Crust Dough",
				"Adding Marinara Sauce",
				"Adding cheese:",
				"    Grated Reggiano Cheese",
				"Adding toppings:",
				"    mushroom",
				"    onion",
				"    red pepper",
				"    sliced pepperoni",
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
			factory.ClamPizza,
			"New-York Style Clam Pizza",
			[]string{
				"Preparing New-York Style Clam Pizza",
				"Tossing Thin Crust Dough",
				"Adding Marinara Sauce",
				"Adding cheese:",
				"    Grated Reggiano Cheese",
				"Adding toppings:",
				"    fresh clams",
				"Bake for 25 minutes at 350",
				"Cutting the pizza into diagonal slices",
				"Place pizza in official PizzaStore box",
			},
			"",
		},
		{
			factory.VeggiePizza,
			"New-York Style Veggie Pizza",
			[]string{
				"Preparing New-York Style Veggie Pizza",
				"Tossing Thin Crust Dough",
				"Adding Marinara Sauce",
				"Adding cheese:",
				"    Grated Reggiano Cheese",
				"Adding toppings:",
				"    mushroom",
				"    onion",
				"    red pepper",
				"Bake for 25 minutes at 350",
				"Cutting the pizza into diagonal slices",
				"Place pizza in official PizzaStore box",
			},
			"",
		},
	}
}

func TestNewChicagoPizzaStore(t *testing.T) {
	pizzaStore := factory.NewChicagoPizzaStore()
	cases := getChicagoTestCases1()
	cases = append(cases, getChicagoTestCases2()...)

	testCases(t, pizzaStore, cases)
}

func getChicagoTestCases1() []testCase {
	return []testCase{
		{
			factory.CheesePizza,
			"Chicago Style Cheese Pizza",
			[]string{
				"Preparing Chicago Style Cheese Pizza",
				"Tossing Thick Crust Dough",
				"Adding Plum Tomato Sauce",
				"Adding cheese:",
				"    Shredded Mozzarella Cheese",
				"    Parmesan Cheese",
				"Adding toppings:",
				"    oregano",
				"Bake for 25 minutes at 350",
				"Cutting the pizza into square slices",
				"Place pizza in official PizzaStore box",
			},
			"",
		},
		{
			factory.PepperoniPizza,
			"Chicago Style Pepperoni Pizza",
			[]string{
				"Preparing Chicago Style Pepperoni Pizza",
				"Tossing Thick Crust Dough",
				"Adding Plum Tomato Sauce",
				"Adding cheese:",
				"    Shredded Mozzarella Cheese",
				"    Parmesan Cheese",
				"Adding toppings:",
				"    eggplant",
				"    spinach",
				"    olives",
				"    sliced pepperoni",
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
			factory.ClamPizza,
			"Chicago Style Clam Pizza",
			[]string{
				"Preparing Chicago Style Clam Pizza",
				"Tossing Thick Crust Dough",
				"Adding Plum Tomato Sauce",
				"Adding cheese:",
				"    Shredded Mozzarella Cheese",
				"    Parmesan Cheese",
				"Adding toppings:",
				"    frozen clams",
				"Bake for 25 minutes at 350",
				"Cutting the pizza into square slices",
				"Place pizza in official PizzaStore box",
			},
			"",
		},
		{
			factory.VeggiePizza,
			"Chicago Style Veggie Pizza",
			[]string{
				"Preparing Chicago Style Veggie Pizza",
				"Tossing Thick Crust Dough",
				"Adding Plum Tomato Sauce",
				"Adding cheese:",
				"    Shredded Mozzarella Cheese",
				"    Parmesan Cheese",
				"Adding toppings:",
				"    eggplant",
				"    spinach",
				"    olives",
				"Bake for 25 minutes at 350",
				"Cutting the pizza into square slices",
				"Place pizza in official PizzaStore box",
			},
			"",
		},
	}
}

func TestPizzaStore_OrderPizza_Error(t *testing.T) {
	pizzaStore := factory.NewNYPizzaStore()
	cases := []testCase{
		{
			"New-York",
			"",
			nil,
			"this pizzaStore can't create New-York pizza",
		},
	}

	testCases(t, pizzaStore, cases)

	pizzaStore = factory.NewChicagoPizzaStore()
	cases = []testCase{
		{
			"Chicago",
			"",
			nil,
			"this pizzaStore can't create Chicago pizza",
		},
	}

	testCases(t, pizzaStore, cases)
}

func testCases(t *testing.T, pizzaStore factory.PizzaStore, cases []testCase) {
	for i := range cases {
		pizza, err := pizzaStore.OrderPizza(cases[i].pizzaType)

		if err != nil {
			assert.Equal(t, cases[i].err, err.Error(), fmt.Sprintf("case %d", i))
		} else {
			assert.Equal(t, cases[i].name, pizza.GetName(), fmt.Sprintf("case %d", i))
			assert.Equal(t, cases[i].logs, pizza.GetLog(), fmt.Sprintf("case %d", i))
		}
	}
}
