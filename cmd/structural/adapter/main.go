package main

import (
	"fmt"
	"time"

	"github.com/Mark-Sart/go-patterns/pkg/structural/adapter"
)

func main() {
	var (
		duck   adapter.Duck   = adapter.MallardDuck{}
		turkey adapter.Turkey = adapter.WildTurkey{}
	)

	seed := time.Now().Unix()

	turkeyAdapter := adapter.NewTurkeyAdapter(turkey)
	duckAdapter := adapter.NewDuckAdapter(duck, seed)

	fmt.Println("\nThe Turkey says...")
	testTurkey(turkey)

	fmt.Println("\nThe Duck says...")
	testDuck(duck)

	fmt.Println("\nThe TurkeyAdapter says...")
	testDuck(turkeyAdapter)

	fmt.Println("\nThe DuckAdapter says...")
	testTurkey(duckAdapter)

	fmt.Println()
}

func testDuck(duck adapter.Duck) {
	fmt.Println(duck.Quack())
	fmt.Println(duck.Fly())
}

func testTurkey(turkey adapter.Turkey) {
	fmt.Println(turkey.Gobble())

	if result := turkey.Fly(); result != "" {
		fmt.Println(result)
	}
}
