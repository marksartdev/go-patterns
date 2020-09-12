package main

import (
	"fmt"

	"github.com/marksartdev/go-patterns/pkg/composite/composite"
)

func main() {
	mallardDuck := composite.NewMallardDuck()
	redheadDuck := composite.NewRedHeadDuck()
	duckCall := composite.NewDuckCall()
	rubberDuck := composite.NewRubberDuck()
	gooseDuck := composite.NewGooseAdapter(composite.NewGoose())

	fmt.Println("\nDuck Simulator")

	simulate(mallardDuck)
	simulate(redheadDuck)
	simulate(duckCall)
	simulate(rubberDuck)
	simulate(gooseDuck)
}

func simulate(duck composite.Quackable) {
	duck.Quack()
}
