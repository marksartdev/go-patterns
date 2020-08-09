package templatemethod

import (
	"bufio"
	"fmt"
	"strings"
)

// Чай.
type tea struct {
	*caffeineBeverage
}

// Заваривает чай.
func (t tea) steepTeaBag() {
	t.write("Steeping the tea")
}

// Добавляет лимон.
func (t tea) addLemon() {
	t.write("Adding Lemon")
}

// Перехватчик, решающий, добавлять ли добавки в напиток.
func (t tea) customerWantsCondiments() bool {
	_, err := fmt.Fprint(t.writer, "Would you like lemon with your tea (y/n)? ")
	t.errorHandler(err)

	buffer := bufio.NewReader(t.reader)
	answer, err := buffer.ReadString('\n')
	t.errorHandler(err)

	if answer == "" {
		answer = "no"
	}

	return strings.ToLower(string(answer[0])) == "y"
}

// NewTea Создает чай.
func NewTea() CaffeineBeverage {
	t := tea{
		caffeineBeverage: newCaffeineBeverage(),
	}
	t.brew = t.steepTeaBag
	t.addCondiments = t.addLemon
	t.condimentsHook = t.customerWantsCondiments

	return t
}

// Кофе.
type coffee struct {
	*caffeineBeverage
}

// Заваривает кофе.
func (c coffee) brewCoffeeGrinds() {
	c.write("Dripping Coffee through filter")
}

// Добавляет сахар и молоко.
func (c coffee) addSugarAndMilk() {
	c.write("Adding Sugar and Milk")
}

// Перехватчик, решающий, добавлять ли добавки в напиток.
func (c coffee) customerWantsCondiments() bool {
	_, err := fmt.Fprint(c.writer, "Would you like milk and sugar with your coffee (y/n)? ")
	c.errorHandler(err)

	buffer := bufio.NewReader(c.reader)
	answer, err := buffer.ReadString('\n')
	c.errorHandler(err)

	if answer == "" {
		answer = "no"
	}

	return strings.ToLower(string(answer[0])) == "y"
}

// NewCoffee Создает кофе.
func NewCoffee() CaffeineBeverage {
	c := coffee{
		caffeineBeverage: newCaffeineBeverage(),
	}
	c.brew = c.brewCoffeeGrinds
	c.addCondiments = c.addSugarAndMilk
	c.condimentsHook = c.customerWantsCondiments

	return c
}
