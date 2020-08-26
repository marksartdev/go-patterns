package state

import "os"

// Состояние "Есть монетка".
type hasQuarterState struct {
	baseState
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
	h.machine.setState(h.machine.getSoldState())
}

// Выдать шарик.
func (h *hasQuarterState) dispense() {
	h.write("No gumball dispensed")
}

func (h *hasQuarterState) String() string {
	return "Machine is processing"
}

// Создать состояние.
func newHasQuarterState(machine GumballMachine) state {
	s := &hasQuarterState{}
	s.machine = machine
	s.writer = os.Stdout

	return s
}
