package main

import (
	"fmt"

	"github.com/marksartdev/go-patterns/pkg/composite/composite"
)

func main() {
	duckFactory := composite.CountingDuckFactory{}
	gooseFactory := composite.GooseFactory{}

	startSimulator(duckFactory, gooseFactory)
}

func startSimulator(duckFactory composite.AbstractDuckFactory, gooseFactory composite.AbstractGooseFactory) {
	redheadDuck := duckFactory.CreateRedHeatDuck()
	duckCall := duckFactory.CreateDuckCall()
	rubberDuck := duckFactory.CreateRubberDuck()
	gooseDuck := gooseFactory.CreateGoose()

	fmt.Println("\nDuck Simulator: With Composite - Flocks")

	flockOfDucks := composite.NewFlock()
	flockOfDucks.Add(redheadDuck)
	flockOfDucks.Add(duckCall)
	flockOfDucks.Add(rubberDuck)
	flockOfDucks.Add(gooseDuck)

	mallardOne := duckFactory.CreateMallardDuck()
	mallardTwo := duckFactory.CreateMallardDuck()
	mallardThree := duckFactory.CreateMallardDuck()
	mallardFour := duckFactory.CreateMallardDuck()

	flockOfMallards := composite.NewFlock()
	flockOfMallards.Add(mallardOne)
	flockOfMallards.Add(mallardTwo)
	flockOfMallards.Add(mallardThree)
	flockOfMallards.Add(mallardFour)

	flockOfDucks.Add(flockOfMallards)

	fmt.Println("\nDuck Simulator: Whole Flock Simulation")
	simulate(flockOfDucks)

	fmt.Println("\nDuck Simulator: Mallard Flock Simulation")
	simulate(flockOfMallards)

	fmt.Printf("\nThe ducks quacked %d times\n", composite.GetQuacks())
}

func simulate(duck composite.Quackable) {
	duck.Quack()
}
