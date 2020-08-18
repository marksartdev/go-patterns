package iterator

import (
	"testing"
	"time"

	"github.com/marksartdev/go-patterns/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestCafeMenu(t *testing.T) {
	menu := NewDinerMenu().(*dinerMenu)
	expected := getExpected()

	for i, item := range expected {
		assert.Equal(t, item, menu.menuItems[i])
	}

	menu.addItem("Fakin' meal", "Fakin' meal for test", true, 0.99)

	for i, item := range expected {
		assert.Equal(t, item, menu.menuItems[i])
	}
}

func TestDinerMenuIterator(t *testing.T) {
	menu := NewDinerMenu().(*dinerMenu)
	dinerIterator := newDinerMenuIterator(menu.menuItems[:])

	assert.EqualError(t, common.IllegalStateError{}, dinerIterator.Remove().Error())

	remove := false

	for dinerIterator.HasNext() {
		if remove {
			assert.NoError(t, dinerIterator.Remove())
		} else {
			dinerIterator.Next()
		}

		remove = !remove
	}

	dinerIterator = newDinerMenuIterator(menu.menuItems[:])

	expected := getExpected()
	for i := range expected {
		if i%2 == 1 {
			item := expected[i]
			assert.Equal(t, item, dinerIterator.Next())
		}
	}
}

func TestAlternatingDinerMenuIterator(t *testing.T) {
	menu := NewDinerMenu().(*dinerMenu)
	dinerIterator := newAlternatingDinerMenuIterator(menu.menuItems[:])

	expected := getExpected()

	weekday := int(time.Now().Weekday())
	start := weekday % 2

	for i := start; i < len(expected); i += 2 {
		assert.True(t, dinerIterator.HasNext())
		assert.Equal(t, expected[i], dinerIterator.Next())
	}

	assert.EqualError(t, common.UnsupportedOperationError{}, dinerIterator.Remove().Error())
}

func getExpected() []menuItem {
	return []menuItem{
		{"Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99},
		{"BLT", "Bacon with lettuce & tomato on whole wheat", false, 2.99},
		{"Soup of the day", "Soup of the day, with a side of potato salad", false, 3.29},
		{"Hotdog", "A hot dog, with sauerkraut, relish, onions, topped with cheese", false, 3.05},
		{"Steamed Veggies and Brown Rice", "Steamed vegetables over brown rice", true, 3.99},
		{"Pasta", "Spaghetti with Marinara Sauce, and a slice of sourdough bread", true, 3.89},
	}
}
