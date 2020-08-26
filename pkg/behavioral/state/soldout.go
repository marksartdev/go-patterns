package state

import "os"

// Состояние "Нет шариков".
type soldOutState struct {
	baseState
}

// Бросить монетку.
func (s *soldOutState) insertQuarter() {
	s.write("You can't insert a quarter, the machine is sold out")
}

// Вернуть монетку.
func (s *soldOutState) ejectQuarter() {
	s.write("You can't eject, you haven't inserted a quarter yet")
}

// Дернуть за рычаг.
func (s *soldOutState) turnCrank() {
	s.write("You terned, but there are no gumballs")
}

// Выдать шарик.
func (s *soldOutState) dispense() {
	s.write("No gumball dispensed")
}

func (s *soldOutState) String() string {
	return "Machine is sold out"
}

// Создать состояние.
func newSoldOutState(machine GumballMachine) state {
	s := &soldOutState{}
	s.machine = machine
	s.writer = os.Stdout

	return s
}
