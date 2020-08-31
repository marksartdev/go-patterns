package proxy

import (
	"io"

	"github.com/marksartdev/go-patterns/pkg/behavioral/state"
)

// GumballMachine Интерфейс автомата с жвачкой.
type GumballMachine interface {
	state.GumballMachine
	GetLocation() string
}

// Автомат с жвачкой.
type gumballMachine struct {
	machine  state.GumballMachine
	location string
}

// InsertQuarter Бросить монетку.
func (g gumballMachine) InsertQuarter() {
	g.machine.InsertQuarter()
}

// EjectQuarter Вернуть монетку.
func (g gumballMachine) EjectQuarter() {
	g.machine.EjectQuarter()
}

// TurnCrank Дернуть за рычаг.
func (g gumballMachine) TurnCrank() {
	g.machine.TurnCrank()
}

// GetCount Получить количество оставшихся шариков.
func (g gumballMachine) GetCount() int {
	return g.machine.GetCount()
}

// GetState Получить текущее состояние.
func (g gumballMachine) GetState() state.State {
	return g.machine.GetState()
}

// SetWriter Установить writer.
func (g gumballMachine) SetWriter(writer io.Writer) {
	g.machine.SetWriter(writer)
}

// GetLocation Получить местоположение автомата.
func (g gumballMachine) GetLocation() string {
	return g.location
}

// NewGumballMachine Создать автомат с жвачкой.
func NewGumballMachine(location string, count int, seed int64) GumballMachine {
	machine := gumballMachine{}
	machine.machine = state.NewGumballMachine(count, seed)
	machine.location = location

	return machine
}
