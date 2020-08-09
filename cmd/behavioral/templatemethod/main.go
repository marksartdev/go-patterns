package main

import (
	"fmt"

	"github.com/Mark-Sart/go-patterns/pkg/behavioral/templatemethod"
)

func main() {
	tea := templatemethod.NewTea()
	coffee := templatemethod.NewCoffee()

	fmt.Println("\nMaking tea...")
	tea.PrepareRecipe()

	fmt.Println("\nMaking coffee...")
	coffee.PrepareRecipe()

	fmt.Println()
}
