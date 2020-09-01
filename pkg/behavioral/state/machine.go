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
	GetCount() int
	GetState() State
	SetWriter(writer io.Writer)
}

// Внутренний интерфейс автомата с жвачкой.
type gumballMachine interface {
	GumballMachine
	releaseBall()
	setState(state State)
	getNoQuarterState() State
	getHasQuarterState() State
	getSoldState() State
	getSoldOutState() State
	getWinnerState() State
}

// Автомат с жвачкой.
type machine struct {
	noQuarterState  State
	hasQuarterState State
	soldState       State
	soldOutState    State
	winnerState     State
	state           State
	count           int
}

// InsertQuarter Бросить монетку.
func (m *machine) InsertQuarter() {
	m.state.insertQuarter()
}

// EjectQuarter Вернуть монетку.
func (m *machine) EjectQuarter() {
	m.state.ejectQuarter()
}

// TurnCrank Дернуть за рычаг.
func (m *machine) TurnCrank() {
	success := m.state.turnCrank()
	if success {
		m.state.dispense()
	}
}

// SetWriter Установить writer.
func (m *machine) SetWriter(writer io.Writer) {
	m.noQuarterState.setWriter(writer)
	m.hasQuarterState.setWriter(writer)
	m.soldState.setWriter(writer)
	m.soldOutState.setWriter(writer)
	m.winnerState.setWriter(writer)
}

// Выдать шарик.
func (m *machine) releaseBall() {
	m.state.write("A gumball comes rolling out the slot...")

	if m.count != 0 {
		m.count--
	}
}

// GetCount Получить количество оставшихся шариков.
func (m *machine) GetCount() int {
	return m.count
}

// GetState Получить текущее состояние.
func (m *machine) GetState() State {
	return m.state
}

// Установить состояние.
func (m *machine) setState(state State) {
	m.state = state
}

// Вернуть состояние "Нет монетки".
func (m *machine) getNoQuarterState() State {
	return m.noQuarterState
}

// Вернуть состояние "Есть монетка".
func (m *machine) getHasQuarterState() State {
	return m.hasQuarterState
}

// Вернуть состояние "Шарик продан".
func (m *machine) getSoldState() State {
	return m.soldState
}

// Вернуть состояние "Нет шариков".
func (m *machine) getSoldOutState() State {
	return m.soldOutState
}

// Вернуть состояние "Победа".
func (m *machine) getWinnerState() State {
	return m.winnerState
}

func (m *machine) String() string {
	msg := "\nMighty Gumball, Inc\n"
	msg += "Go-enabled Standing Gumball Model #2020\n"
	msg += fmt.Sprintf("Inventory: %d gumballs\n", m.count)
	msg += fmt.Sprintln(m.state)

	return msg
}

// NewGumballMachine Создать автомат с жвачкой.
func NewGumballMachine(numberGumballs int, seed int64) GumballMachine {
	m := &machine{}
	m.noQuarterState = newNoQuarterState(m)
	m.hasQuarterState = newHasQuarterState(m, seed)
	m.soldState = newSoldState(m)
	m.soldOutState = newSoldOutState(m)
	m.winnerState = newWinnerState(m)
	m.count = numberGumballs

	if numberGumballs > 0 {
		m.state = m.noQuarterState
	} else {
		m.state = m.soldOutState
	}

	return m
}
