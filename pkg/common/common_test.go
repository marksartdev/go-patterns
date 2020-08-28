package common_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marksartdev/go-patterns/pkg/common"
)

func TestArrayList(t *testing.T) {
	expected := []interface{}{1, "two", -3, "four", 5.0}
	arrayList := common.NewArrayList()

	for _, item := range expected {
		arrayList.Add(item)
	}

	assert.Equal(t, len(expected), arrayList.Size())

	for i, item := range expected {
		assert.Equal(t, item, arrayList.Get(i))
	}

	expected = append(expected[:2], expected[3:]...)

	arrayList.Remove(2)

	for i, item := range expected {
		assert.Equal(t, item, arrayList.Get(i))
	}
}

func TestArrayList_Iterator(t *testing.T) {
	expected := []interface{}{1, "two", -3, "four", 5.0}
	arrayList := common.NewArrayList()

	for _, item := range expected {
		arrayList.Add(item)
	}

	iterator := arrayList.Iterator()
	assert.EqualError(t, common.IllegalStateError{}, iterator.Remove().Error())

	for _, item := range expected {
		assert.True(t, iterator.HasNext())
		assert.Equal(t, item, iterator.Next())
	}

	iterator = arrayList.Iterator()

	for i := range expected {
		if i%2 == 0 {
			_ = iterator.Next()
		} else {
			assert.NoError(t, iterator.Remove())
		}
	}

	assert.NoError(t, iterator.Remove())

	for i, item := range expected {
		if i%2 == 1 {
			assert.Equal(t, item, arrayList.Get((i-1)/2))
		}
	}
}

func TestNewNullIterator(t *testing.T) {
	iterator := common.NewNullIterator()
	assert.False(t, iterator.HasNext())
	assert.Nil(t, iterator.Next())
	assert.EqualError(t, common.UnsupportedOperationError{}, iterator.Remove().Error())
}
