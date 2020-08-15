package main

import (
	"fmt"

	"github.com/marksartdev/go-patterns/pkg/behavioral/iterator"
)

func main() {
	pancakeHouseMenu := iterator.NewPancakeHouseMenu()
	dinerMenu := iterator.NewDinerMenu()
	cafeMenu := iterator.NewCafeMenu()

	waitress := iterator.NewWaitress(pancakeHouseMenu, dinerMenu, cafeMenu)

	fmt.Println()
	waitress.PrintMenu()
	fmt.Println()
}
