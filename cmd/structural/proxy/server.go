package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/marksartdev/go-patterns/pkg/structural/proxy"
)

func startServer(port, location string, count int) {
	reader := bufio.NewReader(os.Stdin)
	wg := &sync.WaitGroup{}

	address := fmt.Sprintf(":%s", port)
	machine := proxy.NewGumballMachine(location, count, time.Now().Unix())

	wg.Add(1)

	go func() {
		err := proxy.StartService(wg, machine, address)
		if err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()

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
		}
	}
}
