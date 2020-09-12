package main

import (
	"fmt"

	"github.com/marksartdev/go-patterns/pkg/composite/composite"
)

func main() {
	factory := composite.CountingDuckFactory{}

	fmt.Println("\nDuck Simulator")
	startSimulator(factory)
	fmt.Printf("\nThe ducks quacked %d times\n", composite.GetQuacks())
}

func startSimulator(factory composite.AbstractDuckFactory) {
	mallardDuck := factory.CreateMallardDuck()
	redheadDuck := factory.CreateRedHeatDuck()
	duckCall := factory.CreateDuckCall()
	rubberDuck := factory.CreateRubberDuck()
	gooseDuck := composite.NewGooseAdapter(composite.NewGoose())

	simulate(mallardDuck)
	simulate(redheadDuck)
	simulate(duckCall)
	simulate(rubberDuck)
	simulate(gooseDuck)
}

func simulate(duck composite.Quackable) {
	duck.Quack()
}
