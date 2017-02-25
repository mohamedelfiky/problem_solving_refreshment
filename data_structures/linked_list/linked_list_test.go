package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	l := new(LinkedList)

	l.Push(5)
	l.Push(10)
	l.Push(15)
	assert.Equal(t, l.head.value, 5, "they should be equal")
	if assert.NotNil(t, l.head.next) {
		assert.Equal(t, l.head.next.value, 10, "they should be equal")

		if assert.NotNil(t, l.head.next.next) {
			assert.Equal(t, l.head.next.next.value, 15, "they should be equal")
		}
	}
}

func TestInsert(t *testing.T) {
	l := new(LinkedList)

	n1 := l.Push(5)
	n2 := l.Insert(n1, 10)
	l.Insert(n2, 15)
	assert.Equal(t, l.head.value, 5, "they should be equal")
	if assert.NotNil(t, l.head.next) {
		assert.Equal(t, l.head.next.value, 10, "they should be equal")

		if assert.NotNil(t, l.head.next.next) {
			assert.Equal(t, l.head.next.next.value, 15, "they should be equal")
		}
	}
}
