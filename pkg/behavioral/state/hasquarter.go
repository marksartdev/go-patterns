package state

import (
	"math/rand"
	"os"
)

const coefficient = 5

// Состояние "Есть монетка".
type hasQuarterState struct {
	baseState
	rand *rand.Rand
}

// Вернуть монетку.
func (h *hasQuarterState) ejectQuarter() {
	h.write("Quarter returned")
	h.machine.setState(h.machine.getNoQuarterState())
}

// Дернуть за рычаг.
func (h *hasQuarterState) turnCrank() bool {
	h.write("You terned...")

	winner := h.rand.Intn(coefficient)
	if winner == 0 && h.machine.GetCount() > 1 {
		h.machine.setState(h.machine.getWinnerState())
	} else {
		h.machine.setState(h.machine.getSoldState())
	}

	return true
}

// Создать состояние.
func newHasQuarterState(machine gumballMachine, seed int64) State {
	s := &hasQuarterState{}
	s.machine = machine
	s.writer = os.Stdout

	source := rand.NewSource(seed)
	// nolint:gosec // It's not important algorithm for security.
	s.rand = rand.New(source)

	return s
}
