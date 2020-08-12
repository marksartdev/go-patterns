package iterator

import "github.com/marksartdev/go-patterns/pkg/common"

// Меню блинной.
type pancakeHouseMenu struct {
	menuItems common.ArrayList
}

// Добавляет блюдо в меню.
func (p pancakeHouseMenu) addItem(name, description string, vegetarian bool, price float64) {
	item := newMenuItem(name, description, vegetarian, price)
	p.menuItems.Add(item)
}

// CreateIterator Создает итератор для меню.
func (p pancakeHouseMenu) CreateIterator() Iterator {
	return NewPancakeHouseMenuIterator(p.menuItems)
}

// NewPancakeHouseMenu Создает меню блинно.
func NewPancakeHouseMenu() Menu {
	menu := pancakeHouseMenu{}
	menu.menuItems = common.NewArrayList()
	menu.addItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs, and toast", true, 2.99)
	menu.addItem("Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", false, 2.99)
	menu.addItem("Blueberry Pancakes", "Pancakes made with fresh blueberries", true, 3.49)
	menu.addItem("Waffles", "Waffles with your choice of blueberries or strawberries", true, 3.59)

	return menu
}
