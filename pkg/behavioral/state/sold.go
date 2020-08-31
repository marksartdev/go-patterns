package state

import "os"

// Состояние "Шарик продан".
type soldState struct {
	baseState
}

// Выдать шарик.
func (s *soldState) dispense() {
	s.machine.releaseBall()

	if s.machine.GetCount() > 0 {
		s.machine.setState(s.machine.getNoQuarterState())
	} else {
		s.write("Oops, out of gumballs!")
		s.machine.setState(s.machine.getSoldOutState())
	}
}

// Создать состояние.
func newSoldState(machine gumballMachine) State {
	s := &soldState{}
	s.machine = machine
	s.writer = os.Stdout

	return s
}
