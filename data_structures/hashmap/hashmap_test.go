package hashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New())
}

func TestPutGetRemove(t *testing.T) {
	h := New()
	h.Put("key1", "Key 1")
	h.Put("key2", "Key 2")
	h.Put("1key", "2Key") // same index with key1
	h.Put("key3", "Key 3")
	assert.Equal(t, h.Get("key1").(MapItem).Value, "Key 1")
	assert.Equal(t, h.Get("key2").(MapItem).Value, "Key 2")
	assert.Equal(t, h.Get("key3").(MapItem).Value, "Key 3")
	assert.Equal(t, h.Get("1key").(MapItem).Value, "2Key")
	assert.Nil(t, h.Get("11key"), "return nil if key not found")
	h.Remove("1key")
	h.Remove("key1")
	assert.Nil(t, h.Get("1key"))
	assert.Nil(t, h.Get("key1"))
}

func TestHindex(t *testing.T) {
	assert.Equal(t, 70, hindex("F"))
	assert.Equal(t, 323%150, hindex("oKFC"))
}
