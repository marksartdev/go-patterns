package main

import (
	"fmt"
	"time"

	"github.com/marksartdev/go-patterns/pkg/behavioral/state"
)

func main() {
	// nolint:gomnd // Example
	gumballMachine := state.NewGumballMachine(5, time.Now().Unix())

	for i := 0; i < 5; i++ {
		fmt.Println(gumballMachine)
		gumballMachine.InsertQuarter()
		gumballMachine.TurnCrank()
	}
}
