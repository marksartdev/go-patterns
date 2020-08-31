package state

import "os"

type winnerState struct {
	baseState
}

// Бросить монетку.
func (w *winnerState) insertQuarter() {
	w.write("Please wait, we're already giving you a gumball")
}

// Вернуть монетку.
func (w *winnerState) ejectQuarter() {
	w.write("Sorry, you already turned th crank")
}

// Дернуть за рычаг.
func (w *winnerState) turnCrank() {
	w.write("Turning twice doesn't get you another gumball!")
}

// Выдать шарик.
func (w *winnerState) dispense() {
	w.write("YOU'RE A WINNER! You get two gumballs for your quarter")
	w.machine.releaseBall()

	if w.machine.getCount() == 0 {
		w.machine.setState(w.machine.getSoldOutState())
	} else {
		w.machine.releaseBall()

		if w.machine.getCount() > 0 {
			w.machine.setState(w.machine.getNoQuarterState())
		} else {
			w.write("Oops, out of gumballs!")
			w.machine.setState(w.machine.getSoldOutState())
		}
	}
}

func (w *winnerState) String() string {
	return "Machine is processing"
}

// Создать состояние.
func newWinnerState(machine GumballMachine) state {
	s := &winnerState{}
	s.machine = machine
	s.writer = os.Stdout

	return s
}
