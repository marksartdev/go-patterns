package main

import (
	"fmt"

	"github.com/marksartdev/go-patterns/pkg/behavioral/state"
)

func main() {
	// nolint:gomnd
	gumballMachine := state.NewGumballMachine(5)

	fmt.Println(gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	fmt.Println(gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.EjectQuarter()
	gumballMachine.TurnCrank()

	fmt.Println(gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.EjectQuarter()

	fmt.Println(gumballMachine)

	gumballMachine.InsertQuarter()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()

	fmt.Println(gumballMachine)
}
