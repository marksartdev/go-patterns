package state

import "os"

// Состояние "Нет монетки".
type noQuarterState struct {
	baseState
}

// Бросить монетку.
func (n *noQuarterState) insertQuarter() {
	n.write("You inserted a quarter")
	n.machine.setState(n.machine.getHasQuarterState())
}

// Вернуть монетку.
func (n *noQuarterState) ejectQuarter() {
	n.write("You haven't inserted a quarter")
}

// Дернуть за рычаг.
func (n *noQuarterState) turnCrank() {
	n.write("You turned, but there's no quarter")
}

// Выдать шарик.
func (n *noQuarterState) dispense() {
	n.write("You need to pay first")
}

func (n *noQuarterState) String() string {
	return "Machine is waiting for quarter"
}

// Создать состояние.
func newNoQuarterState(machine GumballMachine) state {
	s := &noQuarterState{}
	s.machine = machine
	s.writer = os.Stdout

	return s
}
