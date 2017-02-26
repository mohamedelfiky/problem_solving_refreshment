package doubly_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	l := new(DoublyLinkedList)

	a1 := l.Push(5)
	a2 := l.Push(10)
	a3 := l.Push(15)
	assert.Equal(t, 5, l.head.value, "they should be equal")
	assert.Equal(t, a3, l.tail, "they should be equal")
	if assert.NotNil(t, l.head.next) {
		assert.Equal(t, l.head.next.value, 10, "they should be equal")
		assert.Equal(t, l.head.next.prev, a1, "they should be equal")
		assert.Equal(t, l.head.next.next, a3, "they should be equal")

		if assert.NotNil(t, l.head.next.next) {
			assert.Equal(t, l.head.next.next.value, 15, "they should be equal")
			assert.Equal(t, l.head.next.next.prev, a2, "they should be equal")
		}
	}
}

func TestInsert(t *testing.T) {
	l := new(DoublyLinkedList)

	n1 := l.Push(5)
	n3 := l.Push(6)
	n2 := l.Insert(n1, 10)
	l.Insert(n3, 15)
	assert.Equal(t, 5, l.head.value, "they should be equal")
	assert.Equal(t, n2, l.head.next, "they should be equal")
	assert.Equal(t, n3, l.tail.prev, "they should be equal")
	if assert.NotNil(t, l.head.next) {
		assert.Equal(t, 10, l.head.next.value, "they should be equal")
		assert.Equal(t, n1, l.head.next.prev, "they should be equal")

		if assert.NotNil(t, l.head.next.next) {
			assert.Equal(t, l.head.next.next.value, 6, "they should be equal")
			if assert.NotNil(t, l.head.next.next.next) {
				assert.Equal(t, 15, l.head.next.next.next.value, "they should be equal")
			}
		}
	}
}

func TestPushFirst(t *testing.T) {
	l := new(DoublyLinkedList)

	a2 := l.Push(10)
	a3 := l.Push(15)
	a1 := l.PushFirst(5)
	assert.Equal(t, 5, l.head.value, "they should be equal")
	assert.Equal(t, a3, l.tail, "they should be equal")
	if assert.NotNil(t, l.head.next) {
		assert.Equal(t, l.head.next.value, 10, "they should be equal")
		assert.Equal(t, l.head.next.prev, a1, "they should be equal")
		assert.Equal(t, l.head.next.next, a3, "they should be equal")

		if assert.NotNil(t, l.head.next.next) {
			assert.Equal(t, l.head.next.next.value, 15, "they should be equal")
			assert.Equal(t, l.head.next.next.prev, a2, "they should be equal")
		}
	}
}

func TestSearch(t *testing.T) {
	l := new(DoublyLinkedList)
	assert.Nil(t, l.Search(10))

	l.Push(10)
	l.Push(15)
	l.Push(5)
	assert.NotNil(t, l.Search(15))
	assert.Nil(t, l.Search(11))
}

func TestRemove(t *testing.T) {
	l := new(DoublyLinkedList)

	l.Push(1)
	l.Push(5)
	n2 := l.Push(10)
	n3 := l.Push(15)
	l.Remove(n2)
	l.Remove(n3)
	assert.Equal(t, l.head.value, 1, "they should be equal")
	if assert.NotNil(t, l.head.next) {
		assert.Equal(t, l.head.next.value, 5, "they should be equal")
	}
}
