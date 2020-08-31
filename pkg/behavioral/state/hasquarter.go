package state

import (
	"math/rand"
	"os"
)

// Состояние "Есть монетка".
type hasQuarterState struct {
	baseState
	rand *rand.Rand
}

// Бросить монетку.
func (h *hasQuarterState) insertQuarter() {
	h.write("You can't insert another quarter")
}

// Вернуть монетку.
func (h *hasQuarterState) ejectQuarter() {
	h.write("Quarter returned")
	h.machine.setState(h.machine.getNoQuarterState())
}

// Дернуть за рычаг.
func (h *hasQuarterState) turnCrank() {
	h.write("You terned...")

	winner := h.rand.Intn(5)
	if winner == 0 && h.machine.getCount() > 1 {
		h.machine.setState(h.machine.getWinnerState())
	} else {
		h.machine.setState(h.machine.getSoldState())
	}
}

// Выдать шарик.
func (h *hasQuarterState) dispense() {
	h.write("No gumball dispensed")
}

func (h *hasQuarterState) String() string {
	return "Machine is processing"
}

// Создать состояние.
func newHasQuarterState(machine GumballMachine, seed int64) state {
	s := &hasQuarterState{}
	s.machine = machine
	s.writer = os.Stdout

	source := rand.NewSource(seed)
	s.rand = rand.New(source)

	return s
}
