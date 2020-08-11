package singleton_test

import (
	"testing"

	"github.com/marksartdev/go-patterns/pkg/creational/singleton"
	"github.com/stretchr/testify/assert"
)

func BenchmarkGetFullInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = singleton.GetFullInstance()
	}
}

func BenchmarkGetConditionalInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = singleton.GetConditionalInstance()
	}
}

func BenchmarkGetAtomicInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = singleton.GetAtomicInstance()
	}
}

func BenchmarkGetOnceInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = singleton.GetOnceInstance()
	}
}

func TestGetFullInstance(t *testing.T) {
	boilerInstance1 := singleton.GetFullInstance()
	boilerInstance2 := singleton.GetFullInstance()

	assertInstances(t, boilerInstance1, boilerInstance2)
}

func TestGetConditionalInstance(t *testing.T) {
	boilerInstance1 := singleton.GetConditionalInstance()
	boilerInstance2 := singleton.GetConditionalInstance()

	assertInstances(t, boilerInstance1, boilerInstance2)
}

func TestGetAtomicInstance(t *testing.T) {
	boilerInstance1 := singleton.GetAtomicInstance()
	boilerInstance2 := singleton.GetAtomicInstance()

	assertInstances(t, boilerInstance1, boilerInstance2)
}

func TestGetOnceInstance(t *testing.T) {
	boilerInstance1 := singleton.GetOnceInstance()
	boilerInstance2 := singleton.GetOnceInstance()

	assertInstances(t, boilerInstance1, boilerInstance2)
}

func assertInstances(t *testing.T, boilerInstance1, boilerInstance2 singleton.ChocolateBoiler) {
	assert.Equal(t, boilerInstance1, boilerInstance2)

	assert.Equal(t, "Filling ...", boilerInstance1.Fill())
	assert.Equal(t, "", boilerInstance2.Fill())
	assert.Equal(t, "Boiling ...", boilerInstance1.Boil())
	assert.Equal(t, "", boilerInstance2.Boil())
	assert.Equal(t, "Draining ...", boilerInstance1.Drain())
	assert.Equal(t, "", boilerInstance2.Drain())
}
