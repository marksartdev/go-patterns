package state

import (
	"fmt"
	"io"
)

// GumballMachine Интерфейс автомата с жвачкой.
type GumballMachine interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	SetWriter(writer io.Writer)
	releaseBall()
	getCount() int
	setState(state state)
	getNoQuarterState() state
	getHasQuarterState() state
	getSoldState() state
	getSoldOutState() state
	getWinnerState() state
}

// Автомат с жвачкой.
type gumballMachine struct {
	noQuarterState  state
	hasQuarterState state
	soldState       state
	soldOutState    state
	winnerState     state
	state           state
	count           int
}

// InsertQuarter Бросить монетку.
func (g *gumballMachine) InsertQuarter() {
	g.state.insertQuarter()
}

// EjectQuarter Вернуть монетку.
func (g *gumballMachine) EjectQuarter() {
	g.state.ejectQuarter()
}

// TurnCrank Дернуть за рычаг.
func (g *gumballMachine) TurnCrank() {
	g.state.turnCrank()
	g.state.dispense()
}

// SetWriter Установить writer.
func (g *gumballMachine) SetWriter(writer io.Writer) {
	g.noQuarterState.setWriter(writer)
	g.hasQuarterState.setWriter(writer)
	g.soldState.setWriter(writer)
	g.soldOutState.setWriter(writer)
}

// Выдать шарик.
func (g *gumballMachine) releaseBall() {
	g.state.write("A gumball comes rolling out the slot")

	if g.count != 0 {
		g.count--
	}
}

// Получить количество оставшихся шариков.
func (g *gumballMachine) getCount() int {
	return g.count
}

// Установить состояние.
func (g *gumballMachine) setState(state state) {
	g.state = state
}

// Вернуть состояние "Нет монетки".
func (g *gumballMachine) getNoQuarterState() state {
	return g.noQuarterState
}

// Вернуть состояние "Есть монетка".
func (g *gumballMachine) getHasQuarterState() state {
	return g.hasQuarterState
}

// Вернуть состояние "Шарик продан".
func (g *gumballMachine) getSoldState() state {
	return g.soldState
}

// Вернуть состояние "Нет шариков".
func (g *gumballMachine) getSoldOutState() state {
	return g.soldOutState
}

// Вернуть состояние "Победа".
func (g *gumballMachine) getWinnerState() state {
	return g.winnerState
}

func (g *gumballMachine) String() string {
	msg := "\nMighty Gumball, Inc\n"
	msg += "Go-enabled Standing Gumball Model #2020\n"
	msg += fmt.Sprintf("Inventory: %d gumballs\n", g.count)
	msg += fmt.Sprintln(g.state)

	return msg
}

// NewGumballMachine Создать автомат с жвачкой.
func NewGumballMachine(numberGumballs int, seed int64) GumballMachine {
	machine := &gumballMachine{}
	machine.noQuarterState = newNoQuarterState(machine)
	machine.hasQuarterState = newHasQuarterState(machine, seed)
	machine.soldState = newSoldState(machine)
	machine.soldOutState = newSoldOutState(machine)
	machine.winnerState = newWinnerState(machine)
	machine.count = numberGumballs

	if numberGumballs > 0 {
		machine.state = machine.noQuarterState
	} else {
		machine.state = machine.soldOutState
	}

	return machine
}
