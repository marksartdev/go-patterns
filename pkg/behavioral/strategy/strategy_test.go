package strategy

import "testing"

type testDuck struct {
	display      string
	swim         string
	performQuack string
	performFly   string
}

var errString = "%s - пришел некорректный ответ (ожидалось %q, получено %q)"

func TestNewMallardDuck(t *testing.T) {
	duck := NewMallardDuck()

	expected := &testDuck{
		display:      "I'm a mallard duck",
		swim:         "All ducks float, even decoys!",
		performQuack: "Quack-quack-quack-quack-quack",
		performFly:   "I'm flying!!",
	}

	compareResults(expected, duck, t, 5)
}

func TestNewRedheadDuck(t *testing.T) {
	duck := NewRedheadDuck()

	expected := &testDuck{
		display:      "I'm a redhead duck",
		swim:         "All ducks float, even decoys!",
		performQuack: "Quack-quack-quack",
		performFly:   "I'm flying!!",
	}

	compareResults(expected, duck, t, 3)
}

func TestNewRubberDuck(t *testing.T) {
	duck := NewRubberDuck()

	expected := &testDuck{
		display:      "I'm a rubber duck",
		swim:         "All ducks float, even decoys!",
		performQuack: "Squeak-squeak",
		performFly:   "I can't fly!!",
	}

	compareResults(expected, duck, t, 2)
}

func TestNewDecoyDuck(t *testing.T) {
	duck := NewDecoyDuck()

	expected := &testDuck{
		display:      "I'm a decoy duck",
		swim:         "All ducks float, even decoys!",
		performQuack: "<< Silence >>",
		performFly:   "I can't fly!!",
	}

	compareResults(expected, duck, t, 10)
}

func TestNewModelDuck(t *testing.T) {
	duck := NewModelDuck()

	expected := &testDuck{
		display:      "I'm a model duck",
		swim:         "All ducks float, even decoys!",
		performQuack: "Quack-quack-quack-quack",
		performFly:   "I can't fly!!",
	}

	compareResults(expected, duck, t, 4)
}

func TestSetQuackBehavior_MuteQuack(t *testing.T) {
	duck := NewMallardDuck()
	compareQuackBehaviorResult("Quack-quack-quack", duck.PerformQuack(3), t)

	duck.SetQuackBehavior(new(MuteQuack))
	compareQuackBehaviorResult("<< Silence >>", duck.PerformQuack(3), t)
}

func TestSetQuackBehavior_Quack(t *testing.T) {
	duck := NewDecoyDuck()
	compareQuackBehaviorResult("<< Silence >>", duck.PerformQuack(3), t)

	duck.SetQuackBehavior(new(Quack))
	compareQuackBehaviorResult("Quack-quack", duck.PerformQuack(2), t)
}

func TestSetQuackBehavior_Squeak(t *testing.T) {
	duck := NewModelDuck()
	compareQuackBehaviorResult("Quack-quack-quack", duck.PerformQuack(3), t)

	duck.SetQuackBehavior(new(Squeak))
	compareQuackBehaviorResult("Squeak-squeak", duck.PerformQuack(2), t)
}

func TestSetFlyBehavior_FlyNoWay(t *testing.T) {
	duck := NewMallardDuck()
	compareFlyBehaviorResult("I'm flying!!", duck.PerformFly(), t)

	duck.SetFlyBehavior(new(FlyNoWay))
	compareFlyBehaviorResult("I can't fly!!", duck.PerformFly(), t)
}

func TestSetFlyBehavior_FlyWithWings(t *testing.T) {
	duck := NewModelDuck()
	compareFlyBehaviorResult("I can't fly!!", duck.PerformFly(), t)

	duck.SetFlyBehavior(new(FlyWithWings))
	compareFlyBehaviorResult("I'm flying!!", duck.PerformFly(), t)
}

func TestSetFlyBehavior_FlyRocketPowered(t *testing.T) {
	duck := NewModelDuck()
	compareFlyBehaviorResult("I can't fly!!", duck.PerformFly(), t)

	duck.SetFlyBehavior(new(FlyRocketPowered))
	compareFlyBehaviorResult("I'm flying with a rocket!!", duck.PerformFly(), t)
}

func compareResults(expected *testDuck, duck Duck, t *testing.T, quackCount int) {
	result := testDuck{
		display:      duck.Display(),
		swim:         duck.Swim(),
		performQuack: duck.PerformQuack(quackCount),
		performFly:   duck.PerformFly(),
	}

	if expected.display != result.display {
		t.Errorf(errString, "Display()", expected.display, result.display)
	}

	if expected.swim != result.swim {
		t.Errorf(errString, "Swim()", expected.swim, result.swim)
	}

	compareQuackBehaviorResult(expected.performQuack, result.performQuack, t)
	compareFlyBehaviorResult(expected.performFly, result.performFly, t)
}

func compareQuackBehaviorResult(expected string, result string, t *testing.T) {
	if expected != result {
		t.Errorf(errString, "PerformQuack()", expected, result)
	}
}

func compareFlyBehaviorResult(expected string, result string, t *testing.T) {
	if expected != result {
		t.Errorf(errString, "PerformFly()", expected, result)
	}
}
