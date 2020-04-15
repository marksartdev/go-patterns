package main

import (
	"fmt"

	"github.com/Mark-Sart/go-patterns/pkg/strucural/decorator"
)

func main() {
	houseBlend := decorator.NewHouseBlend()
	houseBlend = decorator.NewMilkDecorator(houseBlend)

	darkRoast := decorator.NewDarkRoast()
	darkRoast = decorator.NewMilkDecorator(darkRoast)
	darkRoast = decorator.NewMochaDecorator(darkRoast)

	decaf := decorator.NewDecaf()
	decaf = decorator.NewMilkDecorator(decaf)
	decaf = decorator.NewMochaDecorator(decaf)
	decaf = decorator.NewSoyDecorator(decaf)

	espresso := decorator.NewEspresso()
	espresso = decorator.NewMilkDecorator(espresso)
	espresso = decorator.NewMochaDecorator(espresso)
	espresso = decorator.NewSoyDecorator(espresso)
	espresso = decorator.NewWhipDecorator(espresso)

	fmt.Printf("%s: %.2f\n", houseBlend.GetDescription(), houseBlend.Cost())
	fmt.Printf("%s: %.2f\n", darkRoast.GetDescription(), darkRoast.Cost())
	fmt.Printf("%s: %.2f\n", decaf.GetDescription(), decaf.Cost())
	fmt.Printf("%s: %.2f\n", espresso.GetDescription(), espresso.Cost())
}
