package decorator_test

import (
	"testing"

	"github.com/Mark-Sart/go-patterns/pkg/structural/decorator"

	"github.com/gonum/floats"
)

type expectedData struct {
	description string
	cost        float64
}

const descriptionErrString = "Не соответствует описание напитка. Ожидалось %q, получено %q."
const costErrString = "Не соответствует стоимость напитка. Ожидалось %f, получено %f."

func TestNewHouseBlend(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Домашняя смесь\""
	expected.cost = .89
	result := decorator.NewHouseBlend()
	result.SetSize(decorator.Small)

	checkBeverage(expected, result, t)
}

func TestNewDarkRoast(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Темная обжарка\""
	expected.cost = 1.09
	result := decorator.NewDarkRoast()
	result.SetSize(decorator.Medium)

	checkBeverage(expected, result, t)
}

func TestNewDecaf(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Без кофеина\""
	expected.cost = 1.25
	result := decorator.NewDecaf()
	result.SetSize(decorator.Large)

	checkBeverage(expected, result, t)
}

func TestNewEspresso(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Эспрессо\""
	expected.cost = 1.99
	result := decorator.NewEspresso()
	result.SetSize(decorator.Small)

	checkBeverage(expected, result, t)
}

func TestNewMilkDecorator(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Домашняя смесь\", молочная пена"
	expected.cost = .99 + .15
	result := decorator.NewHouseBlend()
	result = decorator.NewMilkDecorator(result)
	result.SetSize(decorator.Medium)

	checkBeverage(expected, result, t)
}

func TestNewMochaDecorator(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Темная обжарка\", шоколад"
	expected.cost = 1.19 + .30
	result := decorator.NewDarkRoast()
	result = decorator.NewMochaDecorator(result)
	result.SetSize(decorator.Large)

	checkBeverage(expected, result, t)
}

func TestNewSoyDecorator(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Без кофеина\", соя"
	expected.cost = 1.05 + .15
	result := decorator.NewDecaf()
	result = decorator.NewSoyDecorator(result)
	result.SetSize(decorator.Small)

	checkBeverage(expected, result, t)
}

func TestNewWhipDecorator(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Эспрессо\", взбитые сливки"
	expected.cost = 2.09 + .15
	result := decorator.NewEspresso()
	result = decorator.NewWhipDecorator(result)
	result.SetSize(decorator.Medium)

	checkBeverage(expected, result, t)
}

func TestAllCondiments(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Домашняя смесь\", молочная пена, шоколад, соя, взбитые сливки, шоколад, взбитые сливки"
	expected.cost = 1.09 + .20 + .30 + .25 + .20 + .30 + .20
	result := decorator.NewHouseBlend()
	result = decorator.NewMilkDecorator(result)
	result = decorator.NewMochaDecorator(result)
	result = decorator.NewSoyDecorator(result)
	result = decorator.NewWhipDecorator(result)
	result = decorator.NewMochaDecorator(result)
	result = decorator.NewWhipDecorator(result)
	result.SetSize(decorator.Large)

	checkBeverage(expected, result, t)
}

func checkBeverage(expected *expectedData, result decorator.Beverage, t *testing.T) {
	if result.GetDescription() != expected.description {
		t.Errorf(descriptionErrString, expected.description, result.GetDescription())
	}

	if floats.Round(result.Cost(), 10) != floats.Round(expected.cost, 10) {
		t.Errorf(costErrString, expected.cost, result.Cost())
	}
}
