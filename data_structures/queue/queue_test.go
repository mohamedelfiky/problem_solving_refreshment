package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	q := new(Queue)

	q.Push(5)
	if assert.NotNil(t, q.head) {
		assert.Equal(t, q.size, 1, "they should be equal")
		assert.Equal(t, q.head.value, 5, "they should be equal")
		assert.Equal(t, q.tail.value, 5, "they should be equal")
		assert.Nil(t, q.head.next)
	}
	q.Push(6)
	if assert.NotNil(t, q.head.next) {
		assert.Equal(t, q.size, 2, "they should be equal")
		assert.Equal(t, q.head.next.value, 6, "they should be equal")
		assert.Equal(t, q.tail.value, 6, "they should be equal")
		assert.Nil(t, q.head.next.next)
	}
	q.Push(7)
	if assert.NotNil(t, q.head.next.next) {
		assert.Equal(t, q.size, 3, "they should be equal")
		assert.Equal(t, q.head.next.next.value, 7, "they should be equal")
		assert.Equal(t, q.tail.value, 7, "they should be equal")
		assert.Nil(t, q.head.next.next.next)
	}
}

func TestPop(t *testing.T) {
	q := new(Queue)

	assert.Nil(t, q.Pop())

	q.Push(1)
	q.Push(2)
	q.Push(3)

	assert.Equal(t, 3, q.size, "they should be equal")
	assert.Equal(t, 1, q.Pop(), "they should be equal")
	assert.Equal(t, 2, q.head.value, "they should be equal")
	assert.Equal(t, 2, q.size, "they should be equal")
	assert.Equal(t, 2, q.Pop(), "they should be equal")
	assert.Equal(t, 3, q.Pop(), "they should be equal")
	assert.Nil(t, q.head)
	assert.Nil(t, q.tail)
}
