package circular_queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	q := NewCircularQueue()

	q.Push(1)
	assert.Equal(t, q.queue[0], 1, "they should be equal")
	assert.Equal(t, q.front, 0, "they should be equal")
	assert.Equal(t, q.rear, 0, "they should be equal")
	assert.Equal(t, q.size, 1, "they should be equal")

	q.Push(5)
	assert.Equal(t, q.queue[1], 5, "they should be equal")
	assert.Equal(t, q.front, 0, "they should be equal")
	assert.Equal(t, q.rear, 1, "they should be equal")
	assert.Equal(t, q.size, 2, "they should be equal")

	q.Push(6)
	assert.Equal(t, q.queue[2], 6, "they should be equal")
	assert.Equal(t, q.front, 0, "they should be equal")
	assert.Equal(t, q.rear, 2, "they should be equal")

	q.Push(7)
	assert.Equal(t, q.queue[0], 1, "they should be equal")
	assert.Equal(t, q.front, 0, "they should be equal")
	assert.Equal(t, q.rear, 2, "they should be equal")
}

func TestPop(t *testing.T) {
	q := NewCircularQueue()

	q.Push(1)
	q.Push(5)
	q.Push(6)
	q.Push(7)

	assert.Equal(t, q.Pop(), 1, "they should be equal")
	assert.Equal(t, q.Pop(), 5, "they should be equal")
	assert.Equal(t, q.Pop(), 6, "they should be equal")
	assert.Nil(t, q.Pop())

	// test circular

	q.Push(8)
	q.Push(9)
	q.Push(10)
	assert.Equal(t, q.Pop(), 8, "they should be equal")
	q.Push(11)

	assert.Equal(t, q.Pop(), 9, "they should be equal")
	assert.Equal(t, q.Pop(), 10, "they should be equal")
	assert.Equal(t, q.Pop(), 11, "they should be equal")
}
