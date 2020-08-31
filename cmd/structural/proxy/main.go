package main

import (
	"time"

	"github.com/marksartdev/go-patterns/pkg/structural/proxy"
)

func main() {
	machine := proxy.NewGumballMachine("Seattle", 112, time.Now().Unix())
	monitor := proxy.NewGumballMonitor(machine)
	monitor.Report()

	machine.InsertQuarter()
	machine.TurnCrank()

	monitor.Report()
}
