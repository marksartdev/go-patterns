package main

import (
	"fmt"
	"go-patterns/pkg/behavioral/strategy"
	"math/rand"
)

func main() {
	var quackBehavior []strategy.QuackBehavior
	quackBehavior = append(quackBehavior, new(strategy.MuteQuack))
	quackBehavior = append(quackBehavior, new(strategy.Quack))
	quackBehavior = append(quackBehavior, new(strategy.Squeak))

	var flyBehavior []strategy.FlyBehavior
	flyBehavior = append(flyBehavior, new(strategy.FlyNoWay))
	flyBehavior = append(flyBehavior, new(strategy.FlyWithWings))
	flyBehavior = append(flyBehavior, new(strategy.FlyRocketPowered))

	var ducks []strategy.Duck
	ducks = append(ducks, strategy.NewMallardDuck())
	ducks = append(ducks, strategy.NewRedheadDuck())
	ducks = append(ducks, strategy.NewRubberDuck())
	ducks = append(ducks, strategy.NewDecoyDuck())
	ducks = append(ducks, strategy.NewModelDuck())

	for num, duck := range ducks {
		quackCount := rand.Intn(5) + 1

		fmt.Printf("Duck %d\n", num)
		fmt.Printf("Display: %s\n", duck.Display())
		fmt.Printf("Swim: %s\n", duck.Swim())
		fmt.Printf("Quack: %s\n", duck.PerformQuack(quackCount))
		fmt.Printf("Fly: %s\n", duck.PerformFly())

		duck.SetQuackBehavior(quackBehavior[rand.Intn(len(quackBehavior))])
		duck.SetFlyBehavior(flyBehavior[rand.Intn(len(flyBehavior))])

		fmt.Printf("New quack:  %s\n", duck.PerformQuack(quackCount))
		fmt.Printf("New fly:  %s\n", duck.PerformFly())

		fmt.Println()
	}
}
