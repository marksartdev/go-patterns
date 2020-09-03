package main

import (
	"fmt"
	"time"

	"github.com/marksartdev/go-patterns/pkg/structural/proxy"
)

func startClient(ports []string) {
	monitors := make([]proxy.GumballMonitor, 0, len(ports))

	for _, port := range ports {
		address := fmt.Sprintf(":%s", port)

		machine, err := proxy.NewGumballMachineStub(address)
		if err != nil {
			continue
		}

		monitors = append(monitors, proxy.NewGumballMonitor(machine))
	}

	for {
		report(monitors)

		time.Sleep(time.Second)

		// nolint:gomnd // Fix count of report lines.
		clean(5 * len(monitors))
	}
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
