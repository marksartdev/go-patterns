package templatemethod_test

import (
	"bytes"
	"testing"

	"github.com/Mark-Sart/go-patterns/pkg/behavioral/templatemethod"
	"github.com/stretchr/testify/assert"
)

func TestNewTea(t *testing.T) {
	tea := templatemethod.NewTea()

	t.Run("withCondiments", testNewTeaWithCondiments(tea))
	t.Run("withoutCondiments", testNewTeaWithoutCondiments(tea))
}

func testNewTeaWithCondiments(tea templatemethod.CaffeineBeverage) func(t *testing.T) {
	return func(t *testing.T) {
		reader := bytes.NewBufferString("y\n")
		writer := bytes.NewBufferString("")

		tea.SetReader(reader)
		tea.SetWriter(writer)

		// nolint:goconst
		expected := "Boiling water\n"
		expected += "Steeping the tea\n"
		// nolint:goconst
		expected += "Pouring into cup\n"
		expected += "Would you like lemon with your tea (y/n)? Adding Lemon\n"

		tea.PrepareRecipe()

		assert.Equal(t, expected, writer.String())
	}
}

func testNewTeaWithoutCondiments(tea templatemethod.CaffeineBeverage) func(t *testing.T) {
	return func(t *testing.T) {
		reader := bytes.NewBufferString("n\n")
		writer := bytes.NewBufferString("")

		tea.SetReader(reader)
		tea.SetWriter(writer)

		expected := "Boiling water\n"
		expected += "Steeping the tea\n"
		expected += "Pouring into cup\n"
		expected += "Would you like lemon with your tea (y/n)? "

		tea.PrepareRecipe()

		assert.Equal(t, expected, writer.String())
	}
}

func TestNewCoffee(t *testing.T) {
	coffee := templatemethod.NewCoffee()

	t.Run("withCondiments", testNewCoffeeWithCondiments(coffee))
	t.Run("withoutCondiments", testNewCoffeeWithoutCondiments(coffee))
}

func testNewCoffeeWithCondiments(coffee templatemethod.CaffeineBeverage) func(t *testing.T) {
	return func(t *testing.T) {
		reader := bytes.NewBufferString("y\n")
		writer := bytes.NewBufferString("")

		coffee.SetReader(reader)
		coffee.SetWriter(writer)

		expected := "Boiling water\n"
		expected += "Dripping Coffee through filter\n"
		expected += "Pouring into cup\n"
		expected += "Would you like milk and sugar with your coffee (y/n)? Adding Sugar and Milk\n"

		coffee.PrepareRecipe()

		assert.Equal(t, expected, writer.String())
	}
}

func testNewCoffeeWithoutCondiments(coffee templatemethod.CaffeineBeverage) func(t *testing.T) {
	return func(t *testing.T) {
		reader := bytes.NewBufferString("n\n")
		writer := bytes.NewBufferString("")

		coffee.SetReader(reader)
		coffee.SetWriter(writer)

		expected := "Boiling water\n"
		expected += "Dripping Coffee through filter\n"
		expected += "Pouring into cup\n"
		expected += "Would you like milk and sugar with your coffee (y/n)? "

		coffee.PrepareRecipe()

		assert.Equal(t, expected, writer.String())
	}
}
