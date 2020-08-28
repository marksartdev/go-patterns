// Package composite Паттерн "Компоновщик".
package composite

// CreateWaitress Создает официантку и меню.
func CreateWaitress() (Waitress, error) {
	pancakeHouseMenu := newMenu("PANCAKE HOUSE MENU", "Breakfast")
	dinerMenu := newMenu("DINER MENU", "Lunch")
	cafeMenu := newMenu("CAFE MENU", "Dinner")
	dessertMenu := newMenu("DESSERT MENU", "Dessert of course!")

	allMenus := newMenu("ALL MENUS", "All menus combined")
	if err := allMenus.add(pancakeHouseMenu); err != nil {
		return nil, err
	}

	if err := allMenus.add(dinerMenu); err != nil {
		return nil, err
	}

	if err := allMenus.add(cafeMenu); err != nil {
		return nil, err
	}

	if err := fillPancakeHouseMenu(pancakeHouseMenu); err != nil {
		return nil, err
	}

	if err := fillDinerMenu(dinerMenu, dessertMenu); err != nil {
		return nil, err
	}

	if err := fillDessertMenu(dessertMenu); err != nil {
		return nil, err
	}

	if err := fillCafeMenu(cafeMenu); err != nil {
		return nil, err
	}

	return newWaitress(allMenus), nil
}

// Заполняет меню блинной.
func fillPancakeHouseMenu(menu menuComponent) error {
	item := newMenuItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs, and toast", true, 2.99)
	if err := menu.add(item); err != nil {
		return err
	}

	item = newMenuItem("Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", false, 2.99)
	if err := menu.add(item); err != nil {
		return err
	}

	item = newMenuItem("Blueberry Pancakes", "Pancakes made with fresh blueberries, and blueberry syrup", true, 3.49)
	if err := menu.add(item); err != nil {
		return err
	}

	item = newMenuItem("Waffles", "Waffles, with your choice of blueberries or strawberries", true, 3.59)
	if err := menu.add(item); err != nil {
		return err
	}

	return nil
}

// Заполняет меню закусочной.
func fillDinerMenu(menu, dessertMenu menuComponent) error {
	item := newMenuItem("Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99)
	if err := menu.add(item); err != nil {
		return err
	}

	item = newMenuItem("BLT", "Bacon with lettuce & tomato on whole wheat", false, 2.99)
	if err := menu.add(item); err != nil {
		return err
	}

	item = newMenuItem("Soup of the day", "A bowl of the soup of the day, with a side of potato salad", false, 3.29)
	if err := menu.add(item); err != nil {
		return err
	}

	item = newMenuItem("Hotdog", "A hot dog, with sauerkraut, relish, onions, topped with cheese", false, 3.05)
	if err := menu.add(item); err != nil {
		return err
	}

	item = newMenuItem("Steamed Veggies and Brown Rice", "Steamed vegetables over brown rice", true, 3.99)
	if err := menu.add(item); err != nil {
		return err
	}

	item = newMenuItem("Pasta", "Spaghetti with Marinara Sauce, and a slice of sourdough bread", true, 3.89)
	if err := menu.add(item); err != nil {
		return err
	}

	if err := menu.add(dessertMenu); err != nil {
		return err
	}

	return nil
}

// Заполняет десертное меню.
func fillDessertMenu(menu menuComponent) error {
	item := newMenuItem("Apple Pie", "Apple pie with a flaky crust, topped with vanilla icecream", true, 1.59)
	if err := menu.add(item); err != nil {
		return err
	}

	item = newMenuItem("Cheesecake", "Creamy New York cheesecake, with a chocolate graham crust", true, 1.99)
	if err := menu.add(item); err != nil {
		return err
	}

	item = newMenuItem("Sorbet", "A scoop of raspberry and a scoop of lime", true, 1.89)
	if err := menu.add(item); err != nil {
		return err
	}

	return nil
}

// Заполняет меню кафе.
func fillCafeMenu(menu menuComponent) error {
	item := newMenuItem(
		"Veggie Burger and Air Fries",
		"Veggie burger on a whole wheat bun, lettuce, tomato, and fries",
		true,
		3.99,
	)

	if err := menu.add(item); err != nil {
		return err
	}

	item = newMenuItem("Soup of the day", "A cup of the soup of the day, with a side salad", false, 3.69)
	if err := menu.add(item); err != nil {
		return err
	}

	item = newMenuItem("Burrito", "A large burrito, with whole pinto beans, salad, guacamole", true, 4.29)
	if err := menu.add(item); err != nil {
		return err
	}

	return nil
}
