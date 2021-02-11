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

func TestCappedByMaxSizeRemovingOldestFirst(t *testing.T) {
	// given
	cache := New(MaxSize(2))
	cache.Put("first", 1)
	cache.Put("second", 2)
	cache.Put("third", 3)

	// when
	_, firstOk := cache.Get("first")
	secondVal, secondOk := cache.Get("second")
	thirdVal, thirdOk := cache.Get("third")

	// then
	assert.False(t, firstOk)
	assert.True(t, secondOk)
	assert.Equal(t, 2, secondVal)
	assert.True(t, thirdOk)
	assert.Equal(t, 3, thirdVal)
}
