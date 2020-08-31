package state

import "os"

// Состояние "Нет шариков".
type soldOutState struct {
	baseState
}

func (s *soldOutState) String() string {
	return "Machine is sold out"
}

// Создать состояние.
func newSoldOutState(machine gumballMachine) State {
	s := &soldOutState{}
	s.machine = machine
	s.writer = os.Stdout

	return s
}
