package main

import (
	"fmt"
	"math/rand"

	"github.com/Mark-Sart/go-patterns/pkg/behavioral/strategy"
)

const maxQuacks = 5

func main() {
	var quackers []strategy.Quacker
	quackers = append(quackers, new(strategy.MuteQuack))
	quackers = append(quackers, new(strategy.Quack))
	quackers = append(quackers, new(strategy.Squeak))

	var flyers []strategy.Flyer
	flyers = append(flyers, new(strategy.FlyNoWay))
	flyers = append(flyers, new(strategy.FlyWithWings))
	flyers = append(flyers, new(strategy.FlyRocketPowered))

	var ducks []strategy.Duck
	ducks = append(ducks, strategy.NewMallardDuck())
	ducks = append(ducks, strategy.NewRedheadDuck())
	ducks = append(ducks, strategy.NewRubberDuck())
	ducks = append(ducks, strategy.NewDecoyDuck())
	ducks = append(ducks, strategy.NewModelDuck())

	for num, duck := range ducks {
		quackCount := rand.Intn(maxQuacks) + 1

		fmt.Printf("Duck %d\n", num)
		fmt.Printf("Display: %s\n", duck.Display())
		fmt.Printf("Swim: %s\n", duck.Swim())
		fmt.Printf("Quack: %s\n", duck.PerformQuack(quackCount))
		fmt.Printf("Fly: %s\n", duck.PerformFly())

		duck.SetQuacker(quackers[rand.Intn(len(quackers))])
		duck.SetFlyer(flyers[rand.Intn(len(flyers))])

		fmt.Printf("New quack:  %s\n", duck.PerformQuack(quackCount))
		fmt.Printf("New fly:  %s\n", duck.PerformFly())

		fmt.Println()
	}
}
