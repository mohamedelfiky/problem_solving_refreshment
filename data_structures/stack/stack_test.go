package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := new(Stack)
	assert.Nil(t, s.Pop())
	s.Push(1)

	assert.Equal(t, 1, s.top.value)
	assert.Equal(t, 1, s.size)
	s.Push(5)
	assert.Equal(t, 1, s.top.next.value)
	assert.Equal(t, 5, s.top.value)
	assert.Equal(t, 5, s.Pop())
	assert.Equal(t, 1, s.Pop())
}
