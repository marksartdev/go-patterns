package main

import (
	"fmt"

	"github.com/Mark-Sart/go-patterns/pkg/strucural/decorator"
)

func main() {
	decaf := decorator.NewDecaf()
	fmt.Printf("%s: %.1f\n", decaf.GetDescription(), decaf.Cost())

	decaf.SetMilk()
	fmt.Printf("+ milk: %.1f\n", decaf.Cost())

	decaf.SetSoy()
	fmt.Printf("+ soy: %.1f\n", decaf.Cost())

	decaf.SetMocha()
	fmt.Printf("+ mocha: %.1f\n", decaf.Cost())

	decaf.SetWhip()
	fmt.Printf("+ whip: %.1f\n", decaf.Cost())
}
