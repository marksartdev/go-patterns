package main

import (
	"fmt"
	"sort"

	"github.com/marksartdev/go-patterns/pkg/behavioral/templatemethod"
)

func main() {
	tea := templatemethod.NewTea()
	coffee := templatemethod.NewCoffee()

	fmt.Println("\nMaking tea...")
	tea.PrepareRecipe()

	fmt.Println("\nMaking coffee...")
	coffee.PrepareRecipe()

	// nolint:gomnd // Example
	ducks := templatemethod.Ducks{Ducks: []templatemethod.Duck{
		{Name: "Daffy", Wright: 8},
		{Name: "Dewey", Wright: 2},
		{Name: "Howard", Wright: 7},
		{Name: "Louie", Wright: 2},
		{Name: "Donald", Wright: 10},
		{Name: "Huey", Wright: 2},
	}}

	fmt.Println("\nBefore sorting:")
	fmt.Println(ducks)

	sort.Sort(ducks)

	fmt.Println("\nAfter sorting:")
	fmt.Println(ducks)

	fmt.Println()
}
