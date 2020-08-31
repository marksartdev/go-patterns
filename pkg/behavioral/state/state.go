// Package state Паттерн "Состояние".
package state

import (
	"fmt"
	"io"
	"log"
)

// State Интерфейс состояния.
type State interface {
	insertQuarter()
	ejectQuarter()
	turnCrank() bool
	dispense()
	setWriter(writer io.Writer)
	write(smg string)
}

// Базовая структура состояния.
type baseState struct {
	// nolint:structcheck // machine is used in sub-structs.
	machine gumballMachine
	writer  io.Writer
}

// Бросить монетку.
func (b *baseState) insertQuarter() {
	b.operationReject("insertQuarter")
}

// Вернуть монетку.
func (b *baseState) ejectQuarter() {
	b.operationReject("ejectQuarter")
}

// Дернуть за рычаг.
func (b *baseState) turnCrank() bool {
	b.operationReject("turnCrank")

	return false
}

// Выдать шарик.
func (b *baseState) dispense() {
	b.operationReject("dispense")
}

// Установить writer.
func (b *baseState) setWriter(writer io.Writer) {
	b.writer = writer
}

// Записать в writer.
func (b *baseState) write(msg string) {
	if _, err := fmt.Fprintln(b.writer, msg); err != nil {
		log.Fatalln(err)
	}
}

// Выдать сообщение об ошибке.
func (b *baseState) operationReject(operation string) {
	b.write(fmt.Sprintf("Operation %q is rejected", operation))
}

func (b *baseState) String() string {
	return "Machine is processing"
}
