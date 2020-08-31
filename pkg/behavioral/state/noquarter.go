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

func (n *noQuarterState) String() string {
	return "Machine is waiting for quarter"
}

// Создать состояние.
func newNoQuarterState(machine gumballMachine) State {
	s := &noQuarterState{}
	s.machine = machine
	s.writer = os.Stdout

	return s
}
