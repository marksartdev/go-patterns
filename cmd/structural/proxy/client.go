package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/marksartdev/go-patterns/pkg/structural/proxy"
)

func startClient(ports []string) {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	ch := make(chan struct{})

	monitors := make([]proxy.GumballMonitor, 0, len(ports))

	for _, port := range ports {
		addr := fmt.Sprintf(":%s", port)

		machine, err := proxy.NewGumballMachineStub(ctx, wg, addr)
		if err != nil {
			continue
		}

		monitors = append(monitors, proxy.NewGumballMonitor(machine))

		wg.Add(1)
	}

	go readFromConsole(ch)

LOOP:
	for {
		select {
		case <-ch:
			cancel()

			break LOOP
		default:
			report(monitors)
		}

		time.Sleep(time.Second)
		// nolint:gomnd // Fix count of report lines.
		clean(5 * len(monitors))
	}

	wg.Wait()
}

func readFromConsole(ch chan struct{}) {
	reader := bufio.NewReader(os.Stdin)

	_, _ = reader.ReadString('\n')

	ch <- struct{}{}
}

func report(monitors []proxy.GumballMonitor) {
	for i := range monitors {
		monitors[i].Report()
	}
}

func clean(lines int) {
	for i := 0; i < lines; i++ {
		fmt.Print("\033[F\033[K")
	}
}
