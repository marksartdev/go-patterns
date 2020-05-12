package simplefactory

import "fmt"

type pizzaTypeError struct {
	pizzaType string
}

func (p *pizzaTypeError) Error() string {
	return fmt.Sprintf("this factory can't create %s pizza", p.pizzaType)
}

func newPizzaTypeError(pizzaType string) error {
	return &pizzaTypeError{
		pizzaType,
	}
}
