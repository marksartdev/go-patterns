package main

import (
	"fmt"

	"github.com/marksartdev/go-patterns/pkg/behavioral/iterator"
)

func main() {
	pancakeHouseMenu := iterator.NewPancakeHouseMenu()
	breakfastItems := pancakeHouseMenu.GetMenuItems()

	dinerMenu := iterator.NewDinerMenu()
	lunchItems := dinerMenu.GetMenuItems()

	fmt.Println()

	for i := 0; i < breakfastItems.Size(); i++ {
		menuItem := breakfastItems.Get(i).(iterator.MenuItem)
		fmt.Printf("%s %.2f\n", menuItem.GetName(), menuItem.GetPrice())
		fmt.Println(menuItem.GetDescription())
		fmt.Println()
	}

	fmt.Println()

	for i := 0; i < len(lunchItems); i++ {
		menuItem := lunchItems[i]
		if menuItem == nil {
			break
		}

		fmt.Printf("%s %.2f\n", menuItem.GetName(), menuItem.GetPrice())
		fmt.Println(menuItem.GetDescription())
		fmt.Println()
	}

	fmt.Println()
}
