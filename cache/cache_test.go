package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"pgxs.io/chassis/config"
)

type testT struct {
	A string
	B int
}

func TestRedisCacheStore_Set(t *testing.T) {
	config.LoadFromEnvFile()
	var number int
	cache1 := NewRedisCache("test", number, 0)
	cache1.Set("ab", 3)
	val, ok := cache1.Get("ab")
	assert.True(t, ok)
	assert.Equal(t, 3, val)
	tt := &testT{
		A: "test",
		B: 10,
	}
	cache2 := NewRedisCache("test2", testT{}, 5*time.Minute)
	cache2.Set("abc", tt)
	val2, ok := cache2.Get("abc")
	assert.True(t, ok)
	assert.Equal(t, "test", val2.(testT).A)

	assert.True(t, cache2.Contains("abc"))
	cache2.Delete("abc")
	assert.False(t, cache2.Contains("abc"))
}
