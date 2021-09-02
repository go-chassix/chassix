package cache

import (
	"testing"

	"github.com/go-chassix/chassix/v2/config"
	"github.com/stretchr/testify/assert"
)

func TestMemoryCacheStore_Set(t *testing.T) {
	config.LoadFromEnvFile()
	cache, err := NewMemoryCacheStore("test", &testT{}, 6)
	assert.NoError(t, err)
	assert.NotNil(t, cache)
	t1 := &testT{
		A: "456",
		B: 1,
	}
	ok := cache.Set("test", t1)
	assert.True(t, ok)
	val, ok := cache.Get("test")
	assert.True(t, ok)
	assert.Equal(t, "456", val.(*testT).A)
}
