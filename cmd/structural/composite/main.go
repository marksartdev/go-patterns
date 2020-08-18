package main

import (
	"fmt"
	"log"

	"github.com/marksartdev/go-patterns/pkg/structural/composite"
)

func main() {
	pancakeHouseMenu := composite.NewMenu("PANCAKE HOUSE MENU", "Breakfast")
	dinerMenu := composite.NewMenu("DINER MENU", "Lunch")
	cafeMenu := composite.NewMenu("CAFE MENU", "Dinner")
	dessertMenu := composite.NewMenu("DESSERT MENU", "Dessert of course!")

	allMenus := composite.NewMenu("ALL MENUS", "All menus combined")
	errorHandle(allMenus.Add(pancakeHouseMenu))
	errorHandle(allMenus.Add(dinerMenu))
	errorHandle(allMenus.Add(cafeMenu))

	fillPancakeHouseMenu(pancakeHouseMenu)
	fillDinerMenu(dinerMenu, dessertMenu)
	fillDessertMenu(dessertMenu)
	fillCafeMenu(cafeMenu)

	waitress := composite.NewWaitress(allMenus)
	errorHandle(waitress.PrintMenu())

	fmt.Println()
}

func fillPancakeHouseMenu(pancakeHouseMenu composite.MenuComponent) {
	menuItem := composite.NewMenuItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs, and toast", true, 2.99)
	errorHandle(pancakeHouseMenu.Add(menuItem))
	menuItem = composite.NewMenuItem("Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", false, 2.99)
	errorHandle(pancakeHouseMenu.Add(menuItem))
	menuItem = composite.NewMenuItem("Blueberry Pancakes", "Pancakes made with fresh blueberries", true, 3.49)
	errorHandle(pancakeHouseMenu.Add(menuItem))
	menuItem = composite.NewMenuItem("Waffles", "Waffles, with your choice of blueberries or strawberries", true, 3.59)
	errorHandle(pancakeHouseMenu.Add(menuItem))
}

func fillDinerMenu(dinerMenu, dessertMenu composite.MenuComponent) {
	menuItem := composite.NewMenuItem("Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99)
	errorHandle(dinerMenu.Add(menuItem))
	menuItem = composite.NewMenuItem("BLT", "Bacon with lettuce & tomato on whole wheat", false, 2.99)
	errorHandle(dinerMenu.Add(menuItem))
	menuItem = composite.NewMenuItem("Soup of the day", "Soup of the day, with a side of potato salad", false, 3.29)
	errorHandle(dinerMenu.Add(menuItem))
	menuItem = composite.NewMenuItem(
		"Hotdog",
		"A hot dog, with sauerkraut, relish, onions, topped with cheese",
		false,
		3.05,
	)
	errorHandle(dinerMenu.Add(menuItem))
	menuItem = composite.NewMenuItem("Steamed Veggies and Brown Rice", "Steamed vegetables over brown rice", true, 3.99)
	errorHandle(dinerMenu.Add(menuItem))
	menuItem = composite.NewMenuItem("Pasta", "Spaghetti with Marinara Sauce, and a slice of sourdough bread", true, 3.89)
	errorHandle(dinerMenu.Add(menuItem))

	errorHandle(dinerMenu.Add(dessertMenu))
}

func fillDessertMenu(dessertMenu composite.MenuComponent) {
	menuItem := composite.NewMenuItem(
		"Apple Pie",
		"Apple pie with a flaky crust, topped with vanilla icecream",
		true,
		1.59,
	)
	errorHandle(dessertMenu.Add(menuItem))
	menuItem = composite.NewMenuItem("Cheesecake", "Creamy New York cheesecake, with a chocolate graham crust", true, 1.99)
	errorHandle(dessertMenu.Add(menuItem))
	menuItem = composite.NewMenuItem("Sorbet", "A scoop of raspberry and a scoop of lime", true, 1.89)
	errorHandle(dessertMenu.Add(menuItem))
}

func fillCafeMenu(cafeMenu composite.MenuComponent) {
	menuItem := composite.NewMenuItem("Soup of the day", "A cup of the soup of the day, with a side salad", false, 3.69)
	errorHandle(cafeMenu.Add(menuItem))
	menuItem = composite.NewMenuItem("Burrito", "A large burrito, with whole pinto beans, salad, guacamole", true, 4.29)
	errorHandle(cafeMenu.Add(menuItem))
	menuItem = composite.NewMenuItem(
		"Veggie Burger and Air Fries",
		"Veggie burger on a whole wheat bun, lettuce, tomato, and fries",
		true,
		3.99,
	)
	errorHandle(cafeMenu.Add(menuItem))
}

func errorHandle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
