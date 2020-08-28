package composite_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marksartdev/go-patterns/pkg/structural/composite"
)

const delimiter = "--------------------\n"

func TestWaitress(t *testing.T) {
	waitress, err := composite.CreateWaitress()
	assert.NoError(t, err)

	t.Run("PrintMenu", testPrintMenu(waitress))
	t.Run("PrintVegetarianMenu", testPrintVegetarianMenu(waitress))
}

func testPrintMenu(waitress composite.Waitress) func(t *testing.T) {
	return func(t *testing.T) {
		expected := "\nALL MENUS,  All menus combined\n"
		expected += delimiter
		expected += "\nPANCAKE HOUSE MENU,  Breakfast\n"
		expected += delimiter
		expected += "   K&B's Pancake Breakfast(v), 2.99\n"
		expected += "      -- Pancakes with scrambled eggs, and toast\n"
		expected += "   Regular Pancake Breakfast, 2.99\n"
		expected += "      -- Pancakes with fried eggs, sausage\n"
		expected += "   Blueberry Pancakes(v), 3.49\n"
		expected += "      -- Pancakes made with fresh blueberries, and blueberry syrup\n"
		expected += "   Waffles(v), 3.59\n"
		expected += "      -- Waffles, with your choice of blueberries or strawberries\n"
		expected += "\nDINER MENU,  Lunch\n"
		expected += delimiter
		expected += "   Vegetarian BLT(v), 2.99\n"
		expected += "      -- (Fakin') Bacon with lettuce & tomato on whole wheat\n"
		expected += "   BLT, 2.99\n"
		expected += "      -- Bacon with lettuce & tomato on whole wheat\n"
		expected += "   Soup of the day, 3.29\n"
		expected += "      -- A bowl of the soup of the day, with a side of potato salad\n"
		expected += "   Hotdog, 3.05\n"
		expected += "      -- A hot dog, with sauerkraut, relish, onions, topped with cheese\n"
		expected += "   Steamed Veggies and Brown Rice(v), 3.99\n"
		expected += "      -- Steamed vegetables over brown rice\n"
		expected += "   Pasta(v), 3.89\n"
		expected += "      -- Spaghetti with Marinara Sauce, and a slice of sourdough bread\n"
		expected += "\nDESSERT MENU,  Dessert of course!\n"
		expected += delimiter
		expected += "   Apple Pie(v), 1.59\n"
		expected += "      -- Apple pie with a flaky crust, topped with vanilla icecream\n"
		expected += "   Cheesecake(v), 1.99\n"
		expected += "      -- Creamy New York cheesecake, with a chocolate graham crust\n"
		expected += "   Sorbet(v), 1.89\n"
		expected += "      -- A scoop of raspberry and a scoop of lime\n"
		expected += "\nCAFE MENU,  Dinner\n"
		expected += delimiter
		expected += "   Veggie Burger and Air Fries(v), 3.99\n"
		expected += "      -- Veggie burger on a whole wheat bun, lettuce, tomato, and fries\n"
		expected += "   Soup of the day, 3.69\n"
		expected += "      -- A cup of the soup of the day, with a side salad\n"
		expected += "   Burrito(v), 4.29\n"
		expected += "      -- A large burrito, with whole pinto beans, salad, guacamole\n"

		writer := bytes.NewBufferString("")
		waitress.SetWriter(writer)
		err := waitress.PrintMenu()
		assert.NoError(t, err)
		assert.Equal(t, expected, writer.String())
	}
}

func testPrintVegetarianMenu(waitress composite.Waitress) func(t *testing.T) {
	return func(t *testing.T) {
		expected := "\nVEGETARIAN MENU\n"
		expected += "----\n"
		expected += "   K&B's Pancake Breakfast(v), 2.99\n"
		expected += "      -- Pancakes with scrambled eggs, and toast\n"
		expected += "   Blueberry Pancakes(v), 3.49\n"
		expected += "      -- Pancakes made with fresh blueberries, and blueberry syrup\n"
		expected += "   Waffles(v), 3.59\n"
		expected += "      -- Waffles, with your choice of blueberries or strawberries\n"
		expected += "   Vegetarian BLT(v), 2.99\n"
		expected += "      -- (Fakin') Bacon with lettuce & tomato on whole wheat\n"
		expected += "   Steamed Veggies and Brown Rice(v), 3.99\n"
		expected += "      -- Steamed vegetables over brown rice\n"
		expected += "   Pasta(v), 3.89\n"
		expected += "      -- Spaghetti with Marinara Sauce, and a slice of sourdough bread\n"
		expected += "   Apple Pie(v), 1.59\n"
		expected += "      -- Apple pie with a flaky crust, topped with vanilla icecream\n"
		expected += "   Cheesecake(v), 1.99\n"
		expected += "      -- Creamy New York cheesecake, with a chocolate graham crust\n"
		expected += "   Sorbet(v), 1.89\n"
		expected += "      -- A scoop of raspberry and a scoop of lime\n"
		expected += "   Veggie Burger and Air Fries(v), 3.99\n"
		expected += "      -- Veggie burger on a whole wheat bun, lettuce, tomato, and fries\n"
		expected += "   Burrito(v), 4.29\n"
		expected += "      -- A large burrito, with whole pinto beans, salad, guacamole\n"

		writer := bytes.NewBufferString("")
		waitress.SetWriter(writer)
		err := waitress.PrintVegetarianMenu()
		assert.NoError(t, err)
		assert.Equal(t, expected, writer.String())
	}
}
