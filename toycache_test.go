package toycache

import "testing"
import "github.com/stretchr/testify/assert"

func TestEmptyState(t *testing.T) {
	// given
	cache := New()

	// when
	_, ok := cache.Get("missing")

	// then
	assert.False(t, ok)
}

func TestSimplePutAndGet(t *testing.T) {
	// given
	cache := New()
	cache.Put("a", 1)

	// when
	val, ok := cache.Get("a")

	// then
	assert.True(t, ok)
	assert.Equal(t, 1, val)
}
