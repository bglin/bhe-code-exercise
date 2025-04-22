package cache_test

import (
	"github.com/stretchr/testify/assert"
	"ssse-exercise-sieve/pkg/cache"
	"testing"
)

func TestMemoryCache(t *testing.T) {
	memCache := cache.NewMemoryCache()

	tests := []struct {
		desc       string
		action     func()
		assertions func(t *testing.T)
	}{
		{
			desc:   "Empty memCache should not return value",
			action: func() {},
			assertions: func(t *testing.T) {
				_, ok := memCache.Get(0)
				assert.False(t, ok, "expected empty memCache to not return a value")
			},
		},
		{
			desc: "Set primes and validate size",
			action: func() {
				memCache.Set([]int64{2, 3, 5, 7, 11})
			},
			assertions: func(t *testing.T) {
				assert.Equal(t, int64(5), memCache.Size(), "memCache size mismatch")
			},
		},
		{
			desc:   "Check values exist in memCache",
			action: func() {},
			assertions: func(t *testing.T) {
				primes := []int64{2, 3, 5, 7, 11}
				for i, val := range primes {
					cachedVal, ok := memCache.Get(int64(i))
					assert.True(t, ok, "expected value to exist in memCache")
					assert.Equal(t, val, cachedVal)
				}
			},
		},
		{
			desc:   "All should return full slice",
			action: func() {},
			assertions: func(t *testing.T) {
				expected := []int64{2, 3, 5, 7, 11}
				assert.Equal(t, expected, memCache.All(), "All() should return the full primes slice")
			},
		},
		{
			desc: "Shorter Set call should not overwrite",
			action: func() {
				memCache.Set([]int64{2, 3})
			},
			assertions: func(t *testing.T) {
				expected := []int64{2, 3, 5, 7, 11}
				assert.Equal(t, expected, memCache.All(), "shorter Set() call should not overwrite memCache")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tt.action()
			tt.assertions(t)
		})
	}
}
