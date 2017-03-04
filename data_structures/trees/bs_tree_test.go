package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBSTree(t *testing.T) {
	tree := NewBSTree(1, "tree root")

	if assert.NotNil(t, tree.root) {
		assert.Equal(t, 1, tree.root.value)
		assert.Equal(t, "tree root", tree.root.data)
	}
}

func TestInsert(t *testing.T) {
	tree := NewBSTree(5, "tree root")

	tree.Insert(3, "1st level left")
	if assert.NotNil(t, tree.root.left) {
		assert.Equal(t, 3, tree.root.left.value)
		assert.Equal(t, tree.root, tree.root.left.parent)
		assert.Equal(t, "1st level left", tree.root.left.data)
	}

	tree.Insert(7, "1st level right")
	if assert.NotNil(t, tree.root.right) {
		assert.Equal(t, 7, tree.root.right.value)
		assert.Equal(t, tree.root, tree.root.right.parent)
		assert.Equal(t, "1st level right", tree.root.right.data)
	}

	tree.Insert(9, "2st level most right")
	if assert.NotNil(t, tree.root.right.right) {
		assert.Equal(t, 9, tree.root.right.right.value)
		assert.Equal(t, tree.root.right, tree.root.right.right.parent)
		assert.Equal(t, "2st level most right", tree.root.right.right.data)
	}

	tree.Insert(8, "2st level most right left child")
	if assert.NotNil(t, tree.root.right.right.left) {
		assert.Equal(t, 8, tree.root.right.right.left.value)
		assert.Equal(t, tree.root.right.right, tree.root.right.right.left.parent)
		assert.Equal(t, "2st level most right left child", tree.root.right.right.left.data)
	}
}

func TestSearch(t *testing.T) {
	tree := NewBSTree(5, "tree root")
	tree.Insert(3, "1st level left")
	tree.Insert(4, "1st level left")
	tree.Insert(7, "1st level right")
	exists, node := tree.Search(5)
	assert.Equal(t, true, exists)
	assert.Equal(t, tree.root, node)

	exists, node = tree.Search(7)
	assert.Equal(t, true, exists)
	assert.Equal(t, tree.root.right, node)

	exists, node = tree.Search(3)
	assert.Equal(t, true, exists)
	assert.Equal(t, tree.root.left, node)

	exists, node = tree.Search(30)
	assert.Equal(t, false, exists)
	assert.Nil(t, node)
}
