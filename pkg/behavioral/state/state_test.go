package state_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/marksartdev/go-patterns/pkg/behavioral/state"

	"github.com/stretchr/testify/assert"
)

func TestStates(t *testing.T) {
	_, buffer, machine := prepareTest(1)

	assert.Equal(t, "Machine is waiting for quarter", fmt.Sprint(machine.GetState()))
	assert.Equal(t, getMachineInfo(1, "Machine is waiting for quarter"), fmt.Sprint(machine))

	machine.EjectQuarter()
	assert.Equal(t, getErrorText("ejectQuarter"), buffer.String())
	assert.Equal(t, "Machine is waiting for quarter", fmt.Sprint(machine.GetState()))

	buffer.Reset()
	machine.TurnCrank()
	assert.Equal(t, getErrorText("turnCrank"), buffer.String())
	assert.Equal(t, "Machine is waiting for quarter", fmt.Sprint(machine.GetState()))

	buffer.Reset()
	machine.InsertQuarter()
	assert.Equal(t, "You inserted a quarter\n", buffer.String())
	assert.Equal(t, "Machine is processing", fmt.Sprint(machine.GetState()))
	assert.Equal(t, getMachineInfo(1, "Machine is processing"), fmt.Sprint(machine))

	buffer.Reset()
	machine.InsertQuarter()
	assert.Equal(t, getErrorText("insertQuarter"), buffer.String())
	assert.Equal(t, "Machine is processing", fmt.Sprint(machine.GetState()))

	buffer.Reset()
	machine.EjectQuarter()
	assert.Equal(t, "Quarter returned\n", buffer.String())
	assert.Equal(t, "Machine is waiting for quarter", fmt.Sprint(machine.GetState()))
}

func TestSoldState(t *testing.T) {
	i := 1000
	random, buffer, machine := prepareTest(i)

	for random.Intn(5) == 0 {
		machine.InsertQuarter()
		machine.TurnCrank()

		i -= 2
	}

	buffer.Reset()
	machine.InsertQuarter()
	machine.TurnCrank()
	i--

	expected := "You inserted a quarter\n" +
		"You turned...\n" +
		"A gumball comes rolling out the slot...\n"

	assert.Equal(t, expected, buffer.String())
	assert.Equal(t, "Machine is waiting for quarter", fmt.Sprint(machine.GetState()))
	assert.Equal(t, getMachineInfo(i, "Machine is waiting for quarter"), fmt.Sprint(machine))
}

func TestWinnerState(t *testing.T) {
	i := 1000
	random, buffer, machine := prepareTest(i)

	for random.Intn(5) != 0 {
		machine.InsertQuarter()
		machine.TurnCrank()
		i--
	}

	buffer.Reset()
	machine.InsertQuarter()
	machine.TurnCrank()

	i -= 2

	expected := "You inserted a quarter\n" +
		"You turned...\n" +
		"YOU'RE A WINNER! You get two gumballs for your quarter\n" +
		"A gumball comes rolling out the slot...\n" +
		"A gumball comes rolling out the slot...\n"

	assert.Equal(t, expected, buffer.String())
	assert.Equal(t, "Machine is waiting for quarter", fmt.Sprint(machine.GetState()))
	assert.Equal(t, getMachineInfo(i, "Machine is waiting for quarter"), fmt.Sprint(machine))
}

func TestSoldOut(t *testing.T) {
	_, buffer, machine := prepareTest(0)

	machine.InsertQuarter()
	assert.Equal(t, getErrorText("insertQuarter"), buffer.String())
	assert.Equal(t, "Machine is sold out", fmt.Sprint(machine.GetState()))

	buffer.Reset()
	machine.EjectQuarter()
	assert.Equal(t, getErrorText("ejectQuarter"), buffer.String())
	assert.Equal(t, "Machine is sold out", fmt.Sprint(machine.GetState()))

	buffer.Reset()
	machine.TurnCrank()
	assert.Equal(t, getErrorText("turnCrank"), buffer.String())
	assert.Equal(t, "Machine is sold out", fmt.Sprint(machine.GetState()))
	assert.Equal(t, getMachineInfo(0, "Machine is sold out"), fmt.Sprint(machine))
}

func TestStress(t *testing.T) {
	i := 5
	random, _, machine := prepareTest(i)

	for i > 0 {
		machine.InsertQuarter()
		machine.TurnCrank()

		i--
		if random.Intn(5) == 0 {
			i--
		}
	}

	assert.Equal(t, getMachineInfo(0, "Machine is sold out"), fmt.Sprint(machine))
}

func prepareTest(count int) (*rand.Rand, *bytes.Buffer, state.GumballMachine) {
	seed := time.Now().Unix()
	source := rand.NewSource(seed)
	// nolint:gosec // It's not important algorithm for security.
	random := rand.New(source)

	buffer := bytes.NewBufferString("")

	machine := state.NewGumballMachine(count, seed)
	machine.SetWriter(buffer)

	return random, buffer, machine
}

func getErrorText(operation string) string {
	return fmt.Sprintf("Operation %q is rejected\n", operation)
}

func getMachineInfo(count int, status string) string {
	return fmt.Sprintf("\nMighty Gumball, Inc\n"+
		"Go-enabled Standing Gumball Model #2020\n"+
		"Inventory: %d gumballs\n"+
		"%s\n", count, status)
}
