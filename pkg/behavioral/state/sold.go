package state

import "os"

// Состояние "Шарик продан".
type soldState struct {
	baseState
}

// Бросить монетку.
func (s *soldState) insertQuarter() {
	s.write("Please wait, we're already giving you a gumball")
}

// Вернуть монетку.
func (s *soldState) ejectQuarter() {
	s.write("Sorry, you already turned th crank")
}

// Дернуть за рычаг.
func (s *soldState) turnCrank() {
	s.write("Turning twice doesn't get you another gumball!")
}

// Выдать шарик.
func (s *soldState) dispense() {
	s.machine.releaseBall()

	if s.machine.getCount() > 0 {
		s.machine.setState(s.machine.getNoQuarterState())
	} else {
		s.write("Oops, out of gumballs!")
		s.machine.setState(s.machine.getSoldOutState())
	}
}

func (s *soldState) String() string {
	return "Machine is processing"
}

// Создать состояние.
func newSoldState(machine GumballMachine) state {
	s := &soldState{}
	s.machine = machine
	s.writer = os.Stdout

	return s
}
