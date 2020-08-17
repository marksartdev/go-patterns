package main

import (
	"fmt"

	"github.com/marksartdev/go-patterns/pkg/behavioral/iterator"
	"github.com/marksartdev/go-patterns/pkg/common"
)

func main() {
	menus := common.NewArrayList()
	menus.Add(iterator.NewPancakeHouseMenu())
	menus.Add(iterator.NewDinerMenu())
	menus.Add(iterator.NewCafeMenu())

	waitress := iterator.NewWaitress(menus)

	fmt.Println()
	waitress.PrintMenu()
	fmt.Println()
}
