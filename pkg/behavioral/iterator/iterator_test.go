package iterator_test

import (
	"bytes"
	"testing"

	"github.com/marksartdev/go-patterns/pkg/behavioral/iterator"
	"github.com/marksartdev/go-patterns/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestWaitress(t *testing.T) {
	writer := bytes.NewBufferString("")

	menus := common.NewArrayList()
	menus.Add(iterator.NewPancakeHouseMenu())
	menus.Add(iterator.NewDinerMenu())
	menus.Add(iterator.NewCafeMenu())

	waitress := iterator.NewWaitress(menus)
	waitress.SetWriter(writer)

	waitress.PrintMenu()

	expected := "K&B's Pancake Breakfast, 2.99 -- Pancakes with scrambled eggs, and toast\n" +
		"Regular Pancake Breakfast, 2.99 -- Pancakes with fried eggs, sausage\n" +
		"Blueberry Pancakes, 3.49 -- Pancakes made with fresh blueberries\n" +
		"Waffles, 3.59 -- Waffles, with your choice of blueberries or strawberries\n" +
		"Vegetarian BLT, 2.99 -- (Fakin') Bacon with lettuce & tomato on whole wheat\n" +
		"BLT, 2.99 -- Bacon with lettuce & tomato on whole wheat\n" +
		"Soup of the day, 3.29 -- Soup of the day, with a side of potato salad\n" +
		"Hotdog, 3.05 -- A hot dog, with sauerkraut, relish, onions, topped with cheese\n" +
		"Steamed Veggies and Brown Rice, 3.99 -- Steamed vegetables over brown rice\n" +
		"Pasta, 3.89 -- Spaghetti with Marinara Sauce, and a slice of sourdough bread\n" +
		"Soup of the day, 3.69 -- A cup of the soup of the day, with a side salad\n" +
		"Burrito, 4.29 -- A large burrito, with whole pinto beans, salad, guacamole\n" +
		"Veggie Burger and Air Fries, 3.99 -- Veggie burger on a whole wheat bun, lettuce, tomato, and fries\n"

	assert.Equal(t, expected, writer.String())
}
