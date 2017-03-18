package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	l := NewTrie()

	l.Insert("key")
	l.Insert("jojo")
	l.Insert("koz")

	keys1 := l.GetChildren()
	assert.Contains(t, keys1, "k")
	assert.Contains(t, keys1, "j")
	assert.Equal(t, 2, len(keys1))

	keys2 := l.GetChildrenNodes()[byte('k')].GetChildren()
	assert.Contains(t, keys2, "e")
	assert.Contains(t, keys2, "o")
	assert.Equal(t, 2, len(keys1))

	assert.Equal(t, true, l.WordExist("key"))
	assert.Equal(t, true, l.WordExist("jojo"))
	assert.Equal(t, false, l.WordExist("jojoddgf"))
	assert.Equal(t, false, l.WordExist("gf"))
	assert.Equal(t, true, l.PrefixExist("joj"))
	assert.Equal(t, false, l.PrefixExist("jof"))
}
