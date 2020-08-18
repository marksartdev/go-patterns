package iterator

import (
	"fmt"

	"github.com/marksartdev/go-patterns/pkg/common"
)

const maxItems = 6

// Меню закусочной.
type dinerMenu struct {
	numberOfItems int
	menuItems     [maxItems]MenuItem
}

// Добавляет блюдо в меню.
func (d *dinerMenu) addItem(name, description string, vegetarian bool, price float64) {
	item := newMenuItem(name, description, vegetarian, price)

	if d.numberOfItems > maxItems {
		fmt.Println("Sorry, menu is full! Can't add item to menu")

		return
	}

	d.menuItems[d.numberOfItems] = item
	d.numberOfItems++
}

// CreateIterator Создает итератор для меню.
func (d *dinerMenu) CreateIterator() common.Iterator {
	return newAlternatingDinerMenuIterator(d.menuItems[:])
}

// NewDinerMenu Создает меню закусочной.
func NewDinerMenu() Menu {
	menu := &dinerMenu{}
	menu.addItem("Vegetarian BLT", "(Faking) Bacon with lettuce & tomato on whole wheat", true, 2.99)
	menu.addItem("BLT", "Bacon with lettuce & tomato on whole wheat", false, 2.99)
	menu.addItem("Soup of the day", "Soup ogf the day, with a side of potato salad", false, 3.29)
	menu.addItem("Hotdog", "A hot dog, with sauerkraut, relish, onions, topped with cheese", false, 3.05)

	return menu
}
