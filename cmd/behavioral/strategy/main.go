package main

import (
	"fmt"
	"math/rand"

	"github.com/marksartdev/go-patterns/pkg/behavioral/strategy"
)

const maxQuacks = 5

func main() {
	quackers := []strategy.Quacker{new(strategy.MuteQuack), new(strategy.Quack), new(strategy.Squeak)}
	flyers := []strategy.Flyer{new(strategy.FlyNoWay), new(strategy.FlyWithWings), new(strategy.FlyRocketPowered)}
	ducks := []strategy.Duck{
		strategy.NewMallardDuck(),
		strategy.NewRedheadDuck(),
		strategy.NewRubberDuck(),
		strategy.NewDecoyDuck(),
		strategy.NewModelDuck(),
	}

	for num, duck := range ducks {
		// nolint:gosec // Example
		quackCount := rand.Intn(maxQuacks) + 1

		fmt.Printf("Duck %d\n", num)
		fmt.Printf("Display: %s\n", duck.Display())
		fmt.Printf("Swim: %s\n", duck.Swim())
		fmt.Printf("Quack: %s\n", duck.PerformQuack(quackCount))
		fmt.Printf("Fly: %s\n", duck.PerformFly())

		// nolint:gosec // Example
		duck.SetQuacker(quackers[rand.Intn(len(quackers))])
		// nolint:gosec // Example
		duck.SetFlyer(flyers[rand.Intn(len(flyers))])

		fmt.Printf("New quack:  %s\n", duck.PerformQuack(quackCount))
		fmt.Printf("New fly:  %s\n", duck.PerformFly())

		fmt.Println()
	}
}
