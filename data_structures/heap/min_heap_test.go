package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	l := new(MinHeap)
	l.Push(10)
	l.Push(5)
	l.Push(15)
	l.Push(14)
	assert.Equal(t, l.Size(), 4)
}

func TestPop(t *testing.T) {
	l := new(MinHeap)

	assertPanic(t, l.Pop)
	l.Push(10)
	l.Push(5)
	l.Push(25)
	l.Push(35)
	l.Push(15)
	l.Push(105)
	l.Push(14)

	assert.Equal(t, l.Pop(), 5)
	assert.Equal(t, l.Pop(), 10)
	assert.Equal(t, l.Pop(), 14)
}

func assertPanic(t *testing.T, f func() int) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
