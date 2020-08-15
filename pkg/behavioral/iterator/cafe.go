package iterator

import "github.com/marksartdev/go-patterns/pkg/common"

// Меню кафе.
type cafeMenu struct {
	menuItems map[string]MenuItem
}

// Добавляет блюдо в меню.
func (c *cafeMenu) addItem(name, description string, vegetarian bool, price float64) {
	item := newMenuItem(name, description, vegetarian, price)
	c.menuItems[item.GetName()] = item
}

// CreateIterator Создает итератор для меню.
func (c *cafeMenu) CreateIterator() common.Iterator {
	menuList := common.NewArrayList()

	for _, item := range c.menuItems {
		menuList.Add(item)
	}

	return menuList.Iterator()
}

// NewCafeMenu Создает меню кафе.
func NewCafeMenu() Menu {
	menu := &cafeMenu{}
	menu.menuItems = make(map[string]MenuItem)
	menu.addItem(
		"Veggie Burger and Air Fries",
		"Veggie burger ob a whole wheat bun, lettuce, tomato, and fries",
		true,
		3.99,
	)
	menu.addItem("Soup of the day", "A cup of the soup of the day, with a side salad", false, 3.69)
	menu.addItem("Burrito", "A large burrito, with whole pinto beans, salad, guacamole", true, 4.29)

	return menu
}
