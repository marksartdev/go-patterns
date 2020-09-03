package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/marksartdev/go-patterns/pkg/structural/proxy"
)

func startServer(port, location string, count int) {
	reader := bufio.NewReader(os.Stdin)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	addr := fmt.Sprintf(":%s", port)
	machine := proxy.NewGumballMachine(location, count, time.Now().Unix())

	wg.Add(1)

	err := proxy.StartService(ctx, wg, machine, addr)
	if err != nil {
		log.Fatal(err)
	}

LOOP:
	for {
		fmt.Print("-> ")

		text, _ := reader.ReadString('\n')
		switch text {
		case "insert\n":
			machine.InsertQuarter()
		case "eject\n":
			machine.EjectQuarter()
		case "turn\n":
			machine.TurnCrank()
		case "exit\n":
			cancel()

			break LOOP
		}
	}

	wg.Wait()
}
