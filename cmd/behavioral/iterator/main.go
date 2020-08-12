package main

import (
	"fmt"

	"github.com/marksartdev/go-patterns/pkg/behavioral/iterator"
)

func main() {
	pancakeHouseMenu := iterator.NewPancakeHouseMenu()
	dinerMenu := iterator.NewDinerMenu()

	waitress := iterator.NewWaitress(pancakeHouseMenu, dinerMenu)

	fmt.Println()
	waitress.PrintMenu()
	fmt.Println()
}
