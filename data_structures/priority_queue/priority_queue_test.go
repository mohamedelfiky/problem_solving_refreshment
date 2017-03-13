package priority_queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	l := new(PriorityQueue)
	l.Push(10, func() { fmt.Println("Execute 10") })
	l.Push(5, func() { fmt.Println("Execute 5") })
	l.Push(15, func() { fmt.Println("Execute 15") })
	l.Push(14, func() { fmt.Println("Execute 14") })
	assert.Equal(t, l.Size(), 4)
}

func TestPop(t *testing.T) {
	l := new(PriorityQueue)

	assertPanic(t, l.Pop)
	l.Push(10, func() { fmt.Println("Execute 10") })
	l.Push(5, func() { fmt.Println("Execute 5") })
	l.Push(25, func() { fmt.Println("Execute 25") })
	l.Push(35, func() { fmt.Println("Execute 35") })
	l.Push(15, func() { fmt.Println("Execute 15") })
	l.Push(105, func() { fmt.Println("Execute 105") })
	l.Push(14, func() { fmt.Println("Execute 14") })

	assert.Equal(t, l.Pop().Order, 5)
	assert.Equal(t, l.Pop().Order, 10)
	assert.Equal(t, l.Pop().Order, 14)
}

func assertPanic(t *testing.T, f func() Item) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
