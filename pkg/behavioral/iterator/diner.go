package iterator

import "fmt"

const maxItems = 6

// DinerMenu Интерфейс меню закусочной.
type DinerMenu interface {
	GetMenuItems() [maxItems]MenuItem
}

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

// GetMenuItems Возвращает список блюд.
func (d *dinerMenu) GetMenuItems() [maxItems]MenuItem {
	return d.menuItems
}

// NewDinerMenu Создает меню закусочной.
func NewDinerMenu() DinerMenu {
	menu := &dinerMenu{}
	menu.addItem("Vegetarian BLT", "(Faking) Bacon with lettuce & tomato on whole wheat", true, 2.99)
	menu.addItem("BLT", "Bacon with lettuce & tomato on whole wheat", false, 2.99)
	menu.addItem("Soup of the day", "Soup ogf the day, with a side of potato salad", false, 3.29)
	menu.addItem("Hotdog", "A hot dog, with sauerkraut, relish, onions, topped with cheese", false, 3.05)

	return menu
}
