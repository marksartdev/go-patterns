package composite_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marksartdev/go-patterns/pkg/composite/composite"
)

func TestCompositePattern(t *testing.T) {
	t.Run("WithoutCounter", testDuckEmulator(composite.DuckFactory{}, composite.GooseFactory{}))
	t.Run("WithCounter", testDuckEmulator(composite.CountingDuckFactory{}, composite.CountingGooseFactory{}))
}

func testDuckEmulator(df composite.AbstractDuckFactory, gf composite.AbstractGooseFactory) func(t *testing.T) {
	return func(t *testing.T) {
		expectedCount := 0
		if _, ok := df.(composite.CountingDuckFactory); ok {
			expectedCount += 7
		}

		if _, ok := gf.(composite.CountingGooseFactory); ok {
			expectedCount++
		}

		flock, quackBuffer := createFlock(df, gf)

		t.Run("WithoutQuackologist", testFlock(flock, quackBuffer, expectedCount))

		quackologistBuffer := bytes.NewBufferString("")
		quackologist := composite.NewQuackologist()
		quackologist.SetWriter(quackologistBuffer)
		flock.RegisterObserver(quackologist)

		t.Run("WithQuackologist", testFlock(flock, quackBuffer, expectedCount))

		expected := "quackologist: Redhead Duck just quacked\n"
		expected += "quackologist: Duck Call just quacked\n"
		expected += "quackologist: Rubber Duck just quacked\n"
		expected += "quackologist: Goose just quacked\n"
		expected += "quackologist: Mallard Duck just quacked\n"
		expected += "quackologist: Mallard Duck just quacked\n"
		expected += "quackologist: Mallard Duck just quacked\n"
		expected += "quackologist: Mallard Duck just quacked\n"

		assert.Equal(t, expected, quackologistBuffer.String())
	}
}

func testFlock(flock composite.Quackable, buffer *bytes.Buffer, expectedCount int) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("testQuacks", testQuacks(flock, buffer))
		t.Run("testQuacksCount", testQuacksCount(expectedCount))
	}
}

func testQuacks(quacker composite.Quackable, buffer *bytes.Buffer) func(t *testing.T) {
	return func(t *testing.T) {
		quacker.Quack()

		expected := "Quack\n"
		expected += "Kwak\n"
		expected += "Squeak\n"
		expected += "Honk\n"
		expected += "Quack\n"
		expected += "Quack\n"
		expected += "Quack\n"
		expected += "Quack\n"

		assert.Equal(t, expected, buffer.String())

		buffer.Reset()
	}
}

func testQuacksCount(expected int) func(t *testing.T) {
	return func(t *testing.T) {
		assert.Equal(t, expected, composite.GetQuacks())

		composite.ResetCounter()
	}
}

func createFlock(
	duckFactory composite.AbstractDuckFactory,
	gooseFactory composite.AbstractGooseFactory,
) (
	composite.Quackable,
	*bytes.Buffer,
) {
	buffer := bytes.NewBufferString("")

	redheadDuck := duckFactory.CreateRedHeatDuck()
	duckCall := duckFactory.CreateDuckCall()
	rubberDuck := duckFactory.CreateRubberDuck()
	gooseDuck := gooseFactory.CreateGoose()

	flockOfDucks := composite.NewFlock()
	flockOfDucks.Add(redheadDuck)
	flockOfDucks.Add(duckCall)
	flockOfDucks.Add(rubberDuck)
	flockOfDucks.Add(gooseDuck)

	mallardOne := duckFactory.CreateMallardDuck()
	mallardTwo := duckFactory.CreateMallardDuck()
	mallardThree := duckFactory.CreateMallardDuck()
	mallardFour := duckFactory.CreateMallardDuck()

	flockOfMallards := composite.NewFlock()
	flockOfMallards.Add(mallardOne)
	flockOfMallards.Add(mallardTwo)
	flockOfMallards.Add(mallardThree)
	flockOfMallards.Add(mallardFour)

	flockOfDucks.Add(flockOfMallards)

	flockOfDucks.SetWriter(buffer)

	return flockOfDucks, buffer
}
