// Package proxy Паттерн "Заместитель".
package proxy

import (
	"fmt"
	"io"
	"log"
	"os"
)

// GumballMonitor Интерфейс монитора.
type GumballMonitor interface {
	Report()
	SetWriter(writer io.Writer)
}

// Монитор.
type gumballMonitor struct {
	machine GumballMachine
	writer  io.Writer
}

// Report Распечатать отчет.
func (g *gumballMonitor) Report() {
	g.write(fmt.Sprintf("\nGumball Machine: %s", g.machine.GetLocation()))
	g.write(fmt.Sprintf("Current inventory: %d", g.machine.GetCount()))
	g.write(fmt.Sprintf("Current state: %s\n", g.machine.GetState()))
}

// SetWriter Установить writer.
func (g *gumballMonitor) SetWriter(writer io.Writer) {
	g.writer = writer
}

// Записать в writer.
func (g *gumballMonitor) write(msg string) {
	if _, err := fmt.Fprintln(g.writer, msg); err != nil {
		log.Fatalln(err)
	}
}

// NewGumballMonitor Создать монитор.
func NewGumballMonitor(machine GumballMachine) GumballMonitor {
	return &gumballMonitor{machine, os.Stdout}
}
