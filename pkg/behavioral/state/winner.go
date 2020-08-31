package state

import "os"

type winnerState struct {
	baseState
}

// Выдать шарик.
func (w *winnerState) dispense() {
	w.write("YOU'RE A WINNER! You get two gumballs for your quarter")
	w.machine.releaseBall()

	if w.machine.GetCount() == 0 {
		w.machine.setState(w.machine.getSoldOutState())
	} else {
		w.machine.releaseBall()

		if w.machine.GetCount() > 0 {
			w.machine.setState(w.machine.getNoQuarterState())
		} else {
			w.write("Oops, out of gumballs!")
			w.machine.setState(w.machine.getSoldOutState())
		}
	}
}

// Создать состояние.
func newWinnerState(machine gumballMachine) State {
	s := &winnerState{}
	s.machine = machine
	s.writer = os.Stdout

	return s
}
