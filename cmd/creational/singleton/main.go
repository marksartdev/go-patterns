package main

import (
	"fmt"
	"time"

	"github.com/marksartdev/go-patterns/pkg/creational/singleton"
)

func main() {
	ch := make(chan string)
	cancel := make(chan struct{})
	countGoroutines := 1000

	for i := 0; i < countGoroutines; i++ {
		chocolateBoiler := singleton.GetOnceInstance()

		go makeChocolate(chocolateBoiler, ch, cancel)
	}

	check(countGoroutines, ch, cancel)
}

func check(countGoroutines int, ch chan string, cancel chan struct{}) {
ChanLoop:
	for {
		select {
		case msg := <-ch:
			if msg != "" {
				fmt.Println(msg)
			}
		case <-cancel:
			countGoroutines--
			if countGoroutines == 0 {
				break ChanLoop
			}
		}
	}
}

func makeChocolate(chocolateBoiler singleton.ChocolateBoiler, ch chan string, cancel chan struct{}) {
	for i := 0; i < 3; i++ {
		ch <- chocolateBoiler.Fill()
		// nolint:gomnd // Example
		time.Sleep(500 * time.Millisecond)

		ch <- chocolateBoiler.Boil()
		// nolint:gomnd // Example
		time.Sleep(500 * time.Millisecond)

		ch <- chocolateBoiler.Drain()
		// nolint:gomnd // Example
		time.Sleep(500 * time.Millisecond)
	}

	cancel <- struct{}{}
}
