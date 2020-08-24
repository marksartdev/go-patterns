package main

import (
	"fmt"
	"log"

	"github.com/marksartdev/go-patterns/pkg/structural/composite"
)

func main() {
	var (
		waitress composite.Waitress
		err      error
	)

	waitress, err = composite.CreateWaitress()
	errorHandle(err)

	errorHandle(waitress.PrintMenu())
	fmt.Println()

	errorHandle(waitress.PrintVegetarianMenu())
	fmt.Println()
}

func errorHandle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
