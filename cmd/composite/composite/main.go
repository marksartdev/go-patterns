package main

import (
	"fmt"

	"github.com/marksartdev/go-patterns/pkg/composite/composite"
)

func main() {
	mallardDuck := composite.NewQuackCounter(composite.NewMallardDuck())
	redheadDuck := composite.NewQuackCounter(composite.NewRedHeadDuck())
	duckCall := composite.NewQuackCounter(composite.NewDuckCall())
	rubberDuck := composite.NewQuackCounter(composite.NewRubberDuck())
	gooseDuck := composite.NewGooseAdapter(composite.NewGoose())

	fmt.Println("\nDuck Simulator")

	simulate(mallardDuck)
	simulate(redheadDuck)
	simulate(duckCall)
	simulate(rubberDuck)
	simulate(gooseDuck)

	fmt.Printf("\nThe ducks quacked %d times\n", mallardDuck.GetQuacks())
}

func simulate(duck composite.Quackable) {
	duck.Quack()
}
