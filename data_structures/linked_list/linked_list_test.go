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
	assert.Equal(t, l.Head.Value, 5, "they should be equal")
	if assert.NotNil(t, l.Head.next) {
		assert.Equal(t, l.Head.next.Value, 10, "they should be equal")

		if assert.NotNil(t, l.Head.next.next) {
			assert.Equal(t, l.Head.next.next.Value, 15, "they should be equal")
		}
	}
}

func TestIsEmpty(t *testing.T) {
	l := new(LinkedList)

	assert.Equal(t, l.IsEmpty(), true)
	l.Push(10)
	assert.Equal(t, l.IsEmpty(), false)
}

func TestNext(t *testing.T) {
	l := new(LinkedList)

	i1 := l.Push(10)
	i2 := l.Push(11)
	assert.Equal(t, i1.Next(), i2)
}

func TestPushFirst(t *testing.T) {
	l := new(LinkedList)

	l.Push(10)
	l.Push(15)
	l.PushFirst(5)
	assert.Equal(t, l.Head.Value, 5, "they should be equal")
	if assert.NotNil(t, l.Head.next) {
		assert.Equal(t, 10, l.Head.next.Value, "they should be equal")

		if assert.NotNil(t, l.Head.next.next) {
			assert.Equal(t, l.Head.next.next.Value, 15, "they should be equal")
		}
	}
}

func TestSearch(t *testing.T) {
	l := new(LinkedList)
	assert.Nil(t, l.Search(10))

	l.Push(10)
	l.Push(15)
	l.Push(5)
	assert.NotNil(t, l.Search(15))
	assert.Nil(t, l.Search(11))
}

func TestInsert(t *testing.T) {
	l := new(LinkedList)

	n1 := l.Push(5)
	n2 := l.Insert(n1, 10)
	l.Insert(n2, 15)
	assert.Equal(t, l.Head.Value, 5, "they should be equal")
	if assert.NotNil(t, l.Head.next) {
		assert.Equal(t, l.Head.next.Value, 10, "they should be equal")

		if assert.NotNil(t, l.Head.next.next) {
			assert.Equal(t, l.Head.next.next.Value, 15, "they should be equal")
		}
	}
}

func TestRemove(t *testing.T) {
	l := new(LinkedList)

	n1 := l.Push(1)
	l.Push(5)
	n2 := l.Push(10)
	l.Push(15)
	l.Remove(n2)
	assert.Equal(t, l.Head.Value, 1, "they should be equal")
	if assert.NotNil(t, l.Head.next) {
		assert.Equal(t, l.Head.next.Value, 5, "they should be equal")
		if assert.NotNil(t, l.Head.next.next) {
			assert.Equal(t, l.Head.next.next.Value, 15, "they should be equal")
		}
	}
	l.Remove(n1)
	assert.Equal(t, 5, l.Head.Value)
}
