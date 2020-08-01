package adapter_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/Mark-Sart/go-patterns/pkg/structural/adapter"
	"github.com/stretchr/testify/assert"
)

func TestDuckAdapter(t *testing.T) {
	seed := time.Now().Unix()
	source := rand.NewSource(seed)
	random := rand.New(source)

	duck := adapter.MallardDuck{}
	duckAdapter := adapter.NewDuckAdapter(duck, seed)

	assert.Equal(t, "Quack", duckAdapter.Gobble())

	for i := 0; i < 100; i++ {
		expected := ""
		if random.Intn(3) == 0 {
			expected = "I'm flying"
		}

		assert.Equal(t, expected, duckAdapter.Fly())
	}
}

func TestTurkeyAdapter(t *testing.T) {
	turkey := adapter.WildTurkey{}
	turkeyAdapter := adapter.NewTurkeyAdapter(turkey)

	assert.Equal(t, "Gobble gobble", turkeyAdapter.Quack())

	expected := "I'm flying a short distance\n"
	expected += "I'm flying a short distance\n"
	expected += "I'm flying a short distance"

	assert.Equal(t, expected, turkeyAdapter.Fly())
}

func TestIteratorEnumeration(t *testing.T) {
	elements := []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	iterator := &adapter.SimpleIterator{Elements: elements}
	iteratorEnumeration := adapter.NewIteratorEnumeration(iterator)

	assert.NoError(t, iterator.Remove())

	elements = elements[:len(elements)-1]

	for _, elem := range elements {
		assert.True(t, iteratorEnumeration.HasMoreElements())
		assert.Equal(t, elem, iteratorEnumeration.NextElement())
	}

	assert.False(t, iteratorEnumeration.HasMoreElements())
}

func TestEnumerationIterator(t *testing.T) {
	elements := []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	enumeration := &adapter.SimpleEnumeration{Elements: elements}
	enumerationIterator := adapter.NewEnumerationIterator(enumeration)

	for _, elem := range elements {
		assert.Error(t, enumerationIterator.Remove())
		assert.True(t, enumerationIterator.HasNext())
		assert.Equal(t, elem, enumerationIterator.Next())
	}

	assert.False(t, enumerationIterator.HasNext())
}
