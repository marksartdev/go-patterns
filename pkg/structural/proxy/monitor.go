// Package proxy Паттерн "Заместитель".
package proxy

import (
	"context"
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
	machine GumballMachineRemoteClient
	writer  io.Writer
}

// Report Распечатать отчет.
func (g *gumballMonitor) Report() {
	ctx := context.Background()
	location, err := g.machine.GetLocation(ctx, &Nothing{})
	g.errHandle(err)
	count, err := g.machine.GetCount(ctx, &Nothing{})
	g.errHandle(err)
	state, err := g.machine.GetState(ctx, &Nothing{})
	g.errHandle(err)

	g.write(fmt.Sprintf("\nGumball Machine: %s", location.GetData()))
	g.write(fmt.Sprintf("Current inventory: %d", count.GetData()))
	g.write(fmt.Sprintf("Current state: %s\n", state.GetData()))
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

// Обработать ошибку.
func (g *gumballMonitor) errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// NewGumballMonitor Создать монитор.
func NewGumballMonitor(machine GumballMachineRemoteClient) GumballMonitor {
	return &gumballMonitor{machine, os.Stdout}
}
