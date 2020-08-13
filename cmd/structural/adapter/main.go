package main

import (
	"fmt"
	"time"

	"github.com/marksartdev/go-patterns/pkg/common"
	"github.com/marksartdev/go-patterns/pkg/structural/adapter"
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

	elements := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var (
		iterator     common.Iterator     = &adapter.SimpleIterator{Elements: elements}
		iterator2    common.Iterator     = &adapter.SimpleIterator{Elements: elements}
		enumeration  adapter.Enumeration = &adapter.SimpleEnumeration{Elements: elements}
		enumeration2 adapter.Enumeration = &adapter.SimpleEnumeration{Elements: elements}
	)

	fmt.Println("\nTest iterator")
	testIterator(iterator)

	fmt.Println("\nTest enumeration")
	testEnumeration(enumeration)

	enumerationIterator := adapter.NewEnumerationIterator(enumeration2)
	iteratorEnumeration := adapter.NewIteratorEnumeration(iterator2)

	fmt.Println("\nTest enumeration as iterator")
	testIterator(enumerationIterator)

	fmt.Println("\nTest iterator as enumeration")
	testEnumeration(iteratorEnumeration)

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

func testIterator(iterator common.Iterator) {
	for iterator.HasNext() {
		fmt.Printf("%d\t", iterator.Next())
	}

	fmt.Println()
}

func testEnumeration(enumeration adapter.Enumeration) {
	for enumeration.HasMoreElements() {
		fmt.Printf("%d\t", enumeration.NextElement())
	}

	fmt.Println()
}
