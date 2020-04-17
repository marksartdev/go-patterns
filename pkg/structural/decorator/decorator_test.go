package decorator

import (
	"testing"

	"github.com/gonum/floats"
)

var descriptionErrString = "Не соответствует описание напитка. Ожидалось %q, получено %q."
var costErrString = "Не соответствует стоимость напитка. Ожидалось %f, получено %f."

type expectedData struct {
	description string
	cost        float64
}

func TestNewHouseBlend(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Домашняя смесь\""
	expected.cost = .89
	result := NewHouseBlend()
	result.SetSize(Small)

	checkBeverage(expected, result, t)
}

func TestNewDarkRoast(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Темная обжарка\""
	expected.cost = 1.09
	result := NewDarkRoast()
	result.SetSize(Medium)

	checkBeverage(expected, result, t)
}

func TestNewDecaf(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Без кофеина\""
	expected.cost = 1.25
	result := NewDecaf()
	result.SetSize(Large)

	checkBeverage(expected, result, t)
}

func TestNewEspresso(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Эспрессо\""
	expected.cost = 1.99
	result := NewEspresso()
	result.SetSize(Small)

	checkBeverage(expected, result, t)
}

func TestNewMilkDecorator(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Домашняя смесь\", молочная пена"
	expected.cost = .99 + .15
	result := NewHouseBlend()
	result = NewMilkDecorator(result)
	result.SetSize(Medium)

	checkBeverage(expected, result, t)
}

func TestNewMochaDecorator(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Темная обжарка\", шоколад"
	expected.cost = 1.19 + .30
	result := NewDarkRoast()
	result = NewMochaDecorator(result)
	result.SetSize(Large)

	checkBeverage(expected, result, t)
}

func TestNewSoyDecorator(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Без кофеина\", соя"
	expected.cost = 1.05 + .15
	result := NewDecaf()
	result = NewSoyDecorator(result)
	result.SetSize(Small)

	checkBeverage(expected, result, t)
}

func TestNewWhipDecorator(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Эспрессо\", взбитые сливки"
	expected.cost = 2.09 + .15
	result := NewEspresso()
	result = NewWhipDecorator(result)
	result.SetSize(Medium)

	checkBeverage(expected, result, t)
}

func TestAllCondiments(t *testing.T) {
	expected := new(expectedData)
	expected.description = "Кофе \"Домашняя смесь\", молочная пена, шоколад, соя, взбитые сливки, шоколад, взбитые сливки"
	expected.cost = 1.09 + .20 + .30 + .25 + .20 + .30 + .20
	result := NewHouseBlend()
	result = NewMilkDecorator(result)
	result = NewMochaDecorator(result)
	result = NewSoyDecorator(result)
	result = NewWhipDecorator(result)
	result = NewMochaDecorator(result)
	result = NewWhipDecorator(result)
	result.SetSize(Large)

	checkBeverage(expected, result, t)
}

func checkBeverage(expected *expectedData, result Beverage, t *testing.T) {
	if result.GetDescription() != expected.description {
		t.Errorf(descriptionErrString, expected.description, result.GetDescription())
	}

	if floats.Round(result.Cost(), 10) != floats.Round(expected.cost, 10) {
		t.Errorf(costErrString, expected.cost, result.Cost())
	}
}
