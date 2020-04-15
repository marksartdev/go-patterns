package decorator

import "testing"

var descriptionErrString = "Не соответствует описание напитка. Ожидалось %q, получено %q."
var costErrString = "Не соответствует стоимость напитка. Ожидалось %f, получено %f."

func TestNewHouseBlend(t *testing.T) {
	expected := new(beverage)
	expected.description = "Кофе \"Домашняя смесь\""
	expected.cost = .89
	result := NewHouseBlend()

	checkBeverage(expected, result, t)
}

func TestNewDarkRoast(t *testing.T) {
	expected := new(beverage)
	expected.description = "Кофе \"Темная обжарка\""
	expected.cost = .99
	result := NewDarkRoast()

	checkBeverage(expected, result, t)
}

func TestNewDecaf(t *testing.T) {
	expected := new(beverage)
	expected.description = "Кофе \"Без кофеина\""
	expected.cost = 1.05
	result := NewDecaf()

	checkBeverage(expected, result, t)
}

func TestNewEspresso(t *testing.T) {
	expected := new(beverage)
	expected.description = "Кофе \"Эспрессо\""
	expected.cost = 1.99
	result := NewEspresso()

	checkBeverage(expected, result, t)
}

func checkBeverage(expected Beverage, result Beverage, t *testing.T) {
	if result.GetDescription() != expected.GetDescription() {
		t.Errorf(descriptionErrString, expected.GetDescription(), result.GetDescription())
	}

	if result.Cost() != expected.Cost() {
		t.Errorf(costErrString, expected.Cost(), result.Cost())
	}
}
