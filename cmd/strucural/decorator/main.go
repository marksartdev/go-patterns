package main

import (
	"fmt"
	"math/rand"

	"github.com/Mark-Sart/go-patterns/pkg/strucural/decorator"
)

func main() {
	beverages := []func() decorator.Beverage{
		decorator.NewHouseBlend,
		decorator.NewDarkRoast,
		decorator.NewDecaf,
		decorator.NewEspresso,
	}

	condiments := []func(decorator.Beverage) decorator.Beverage{
		decorator.NewMilkDecorator,
		decorator.NewMochaDecorator,
		decorator.NewSoyDecorator,
		decorator.NewWhipDecorator,
	}

	for i := 0; i < 10; i++ {
		beverage := beverages[rand.Intn(len(beverages))]()
		beverage = condiments[rand.Intn(len(condiments))](beverage)
		beverage = condiments[rand.Intn(len(condiments))](beverage)

		fmt.Printf("%s: %.2f\n", beverage.GetDescription(), beverage.Cost())
	}
}
