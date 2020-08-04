package main

import (
	"fmt"

	"github.com/Mark-Sart/go-patterns/pkg/structural/facade"
)

func main() {
	homeTheater := facade.NewHomeTheater()

	fmt.Println()
	fmt.Println(homeTheater.WatchMovie("The Big Bang Theory"))
	fmt.Println()
	fmt.Println()
	fmt.Println(homeTheater.EndMovie())
	fmt.Println()
}
