package iterator

import "github.com/marksartdev/go-patterns/pkg/common"

// PancakeHouseMenu Интерфейс меню блинной.
type PancakeHouseMenu interface {
	GetMenuItems() common.ArrayList
}

// Меню блинной.
type pancakeHouseMenu struct {
	menuItems common.ArrayList
}

// Добавляет блюдо в меню.
func (p pancakeHouseMenu) addItem(name, description string, vegetarian bool, price float64) {
	item := newMenuItem(name, description, vegetarian, price)
	p.menuItems.Add(item)
}

// GetMenuItems Возвращает список блюд.
func (p pancakeHouseMenu) GetMenuItems() common.ArrayList {
	return p.menuItems
}

// NewPancakeHouseMenu Создает меню блинно.
func NewPancakeHouseMenu() PancakeHouseMenu {
	menu := pancakeHouseMenu{}
	menu.menuItems = common.NewArrayList()
	menu.addItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs, and toast", true, 2.99)
	menu.addItem("Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", false, 2.99)
	menu.addItem("Blueberry Pancakes", "Pancakes made with fresh blueberries", true, 3.49)
	menu.addItem("Waffles", "Waffles with your choice of blueberries or strawberries", true, 3.59)

	return menu
}
