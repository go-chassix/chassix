package cache

import (
	"testing"
	"time"

	"c6x.io/chassix.v2/config"
	"github.com/stretchr/testify/assert"
)

type testT struct {
	A string
	B int
	Inner
}

type testT2 struct {
	A string
	B int
}
type Inner struct {
	C string
}

func TestRedisCacheStore_Set(t *testing.T) {
	config.LoadFromEnvFile()
	var number int
	cache1, err := NewRedisCacheStore("test", number, 0)
	ok := cache1.Set("ab", 3)
	assert.True(t, ok)
	ok = cache1.Set("ab", "3")
	assert.False(t, ok)
	intVal := 3
	ok = cache1.Set("ab", &intVal)
	assert.False(t, ok)
	val, ok := cache1.Get("ab")
	assert.True(t, ok)
	assert.Equal(t, 3, val)

	cache11, err := NewRedisCacheStore("test", number, 0)
	assert.NoError(t, err)
	assert.NotNil(t, cache11)
	tt := &testT{
		A: "test",
		B: 10,
	}

	tt2 := testT{A: "ab", B: 3, Inner: Inner{
		C: "a",
	}}
	t21 := &testT2{
		A: "t",
		B: 2,
	}
	cache2, err := NewRedisCacheStore("test2", &testT{}, 5*time.Minute)
	assert.NoError(t, err)
	assert.True(t, ok)
	ok2 := cache2.Set("abc", tt)
	assert.True(t, ok2)
	ok3 := cache2.Set("cde", tt2)
	assert.False(t, ok3)

	ok4 := cache2.Set("ijk", &tt2)
	assert.True(t, ok4)
	ok5 := cache2.Set("efg", t21)
	assert.False(t, ok5)

	val2, ok6 := cache2.Get("abc")
	assert.True(t, ok6)
	assert.Equal(t, "test", val2.(*testT).A)

	val3, ok61 := cache2.Get("cde")
	assert.False(t, ok61)
	assert.Nil(t, val3)

	val4, ok62 := cache2.Get("ijk")
	assert.True(t, ok62)
	assert.Equal(t, tt2.B, val4.(*testT).B)
	assert.Equal(t, tt2.A, val4.(*testT).A)

	assert.True(t, cache2.Contains("abc"))
	cache2.Delete("abc")
	assert.False(t, cache2.Contains("abc"))

	cache3, err := NewRedisCacheStore("test3", testT{}, 5*time.Minute)
	assert.NoError(t, err)
	assert.NotNil(t, cache3)
	ok7 := cache3.Set("hde", tt2)
	assert.True(t, ok7)
	val8, ok8 := cache3.Get("hde")
	assert.True(t, ok8)
	assert.NotNil(t, val8)
	val8t := val8.(testT)
	assert.Equal(t, tt2.A, val8t.A)
	assert.Equal(t, tt2.B, val8t.B)
	assert.Equal(t, tt2.C, val8t.C)

}
