package main

import (
	"fmt"

	"github.com/marksartdev/go-patterns/pkg/composite/composite"
)

func main() {
	duckFactory := composite.CountingDuckFactory{}
	gooseFactory := composite.GooseFactory{}

	fmt.Println("\nDuck Simulator")
	startSimulator(duckFactory, gooseFactory)
	fmt.Printf("\nThe ducks quacked %d times\n", composite.GetQuacks())
}

func startSimulator(duckFactory composite.AbstractDuckFactory, gooseFactory composite.AbstractGooseFactory) {
	mallardDuck := duckFactory.CreateMallardDuck()
	redheadDuck := duckFactory.CreateRedHeatDuck()
	duckCall := duckFactory.CreateDuckCall()
	rubberDuck := duckFactory.CreateRubberDuck()
	gooseDuck := gooseFactory.CreateGoose()

	simulate(mallardDuck)
	simulate(redheadDuck)
	simulate(duckCall)
	simulate(rubberDuck)
	simulate(gooseDuck)
}

func simulate(duck composite.Quackable) {
	duck.Quack()
}
