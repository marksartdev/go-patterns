package strategy_test

import (
	"testing"

	"github.com/marksartdev/go-patterns/pkg/behavioral/strategy"
)

const errString = "Некорректный результат. Ожидалось %q, получено %q."

type testDuck struct {
	display      string
	swim         string
	performQuack string
	performFly   string
}

func TestNewMallardDuck(t *testing.T) {
	duck := strategy.NewMallardDuck()

	expected := &testDuck{
		display:      "I'm a mallard duck",
		swim:         "All ducks float, even decoys!",
		performQuack: "Quack-quack-quack-quack-quack",
		performFly:   "I'm flying!!",
	}

	compareResults(expected, duck, t, 5)
}

func TestNewRedheadDuck(t *testing.T) {
	duck := strategy.NewRedheadDuck()

	expected := &testDuck{
		display:      "I'm a redhead duck",
		swim:         "All ducks float, even decoys!",
		performQuack: "Quack-quack-quack",
		performFly:   "I'm flying!!",
	}

	compareResults(expected, duck, t, 3)
}

func TestNewRubberDuck(t *testing.T) {
	duck := strategy.NewRubberDuck()

	expected := &testDuck{
		display:      "I'm a rubber duck",
		swim:         "All ducks float, even decoys!",
		performQuack: "Squeak-squeak",
		performFly:   "I can't fly!!",
	}

	compareResults(expected, duck, t, 2)
}

func TestNewDecoyDuck(t *testing.T) {
	duck := strategy.NewDecoyDuck()

	expected := &testDuck{
		display:      "I'm a decoy duck",
		swim:         "All ducks float, even decoys!",
		performQuack: "<< Silence >>",
		performFly:   "I can't fly!!",
	}

	compareResults(expected, duck, t, 10)
}

func TestNewModelDuck(t *testing.T) {
	duck := strategy.NewModelDuck()

	expected := &testDuck{
		display:      "I'm a model duck",
		swim:         "All ducks float, even decoys!",
		performQuack: "Quack-quack-quack-quack",
		performFly:   "I can't fly!!",
	}

	compareResults(expected, duck, t, 4)
}

func TestSetQuacker_MuteQuack(t *testing.T) {
	duck := strategy.NewMallardDuck()
	compareQuackBehaviorResult("Quack-quack-quack", duck.PerformQuack(3), t)

	duck.SetQuacker(new(strategy.MuteQuack))
	compareQuackBehaviorResult("<< Silence >>", duck.PerformQuack(3), t)
}

func TestSetQuacker_Quack(t *testing.T) {
	duck := strategy.NewDecoyDuck()
	compareQuackBehaviorResult("<< Silence >>", duck.PerformQuack(3), t)

	duck.SetQuacker(new(strategy.Quack))
	compareQuackBehaviorResult("Quack-quack", duck.PerformQuack(2), t)
}

func TestSetQuacker_Squeak(t *testing.T) {
	duck := strategy.NewModelDuck()
	compareQuackBehaviorResult("Quack-quack-quack", duck.PerformQuack(3), t)

	duck.SetQuacker(new(strategy.Squeak))
	compareQuackBehaviorResult("Squeak-squeak", duck.PerformQuack(2), t)
}

func TestSetFlyer_FlyNoWay(t *testing.T) {
	duck := strategy.NewMallardDuck()
	compareFlyBehaviorResult("I'm flying!!", duck.PerformFly(), t)

	duck.SetFlyer(new(strategy.FlyNoWay))
	compareFlyBehaviorResult("I can't fly!!", duck.PerformFly(), t)
}

func TestSetFlyer_FlyWithWings(t *testing.T) {
	duck := strategy.NewModelDuck()
	compareFlyBehaviorResult("I can't fly!!", duck.PerformFly(), t)

	duck.SetFlyer(new(strategy.FlyWithWings))
	compareFlyBehaviorResult("I'm flying!!", duck.PerformFly(), t)
}

func TestSetFlyer_FlyRocketPowered(t *testing.T) {
	duck := strategy.NewModelDuck()
	compareFlyBehaviorResult("I can't fly!!", duck.PerformFly(), t)

	duck.SetFlyer(new(strategy.FlyRocketPowered))
	compareFlyBehaviorResult("I'm flying with a rocket!!", duck.PerformFly(), t)
}

func compareResults(expected *testDuck, duck strategy.Duck, t *testing.T, quackCount int) {
	result := testDuck{
		display:      duck.Display(),
		swim:         duck.Swim(),
		performQuack: duck.PerformQuack(quackCount),
		performFly:   duck.PerformFly(),
	}

	if expected.display != result.display {
		t.Errorf(errString, expected.display, result.display)
	}

	if expected.swim != result.swim {
		t.Errorf(errString, expected.swim, result.swim)
	}

	compareQuackBehaviorResult(expected.performQuack, result.performQuack, t)
	compareFlyBehaviorResult(expected.performFly, result.performFly, t)
}

func compareQuackBehaviorResult(expected, result string, t *testing.T) {
	if expected != result {
		t.Errorf(errString, expected, result)
	}
}

func compareFlyBehaviorResult(expected, result string, t *testing.T) {
	if expected != result {
		t.Errorf(errString, expected, result)
	}
}
