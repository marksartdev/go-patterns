// Package state Паттерн "Состояние".
package state

import (
	"fmt"
	"io"
	"log"
)

// Интерфейс состояния.
type state interface {
	insertQuarter()
	ejectQuarter()
	turnCrank()
	dispense()
	setWriter(writer io.Writer)
	write(smg string)
}

// Базовая структура состояния.
type baseState struct {
	// nolint:structcheck // machine is used in sub-structs.
	machine GumballMachine
	writer  io.Writer
}

// Установить writer.
func (b *baseState) setWriter(writer io.Writer) {
	b.writer = writer
}

// Вывести на экран.
func (b *baseState) write(msg string) {
	if _, err := fmt.Fprintln(b.writer, msg); err != nil {
		log.Fatalln(err)
	}
}
