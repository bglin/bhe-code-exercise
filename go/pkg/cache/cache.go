package cache

type MemoryCache struct {
	primes []int64
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{primes: make([]int64, 0)}
}

func (c *MemoryCache) Get(n int64) (int64, bool) {
	if n < int64(len(c.primes)) {
		return c.primes[n], true
	}
	return 0, false
}

func (c *MemoryCache) Set(primes []int64) {
	if int64(len(primes)) > int64(len(c.primes)) {
		c.primes = primes
	}
}

func (c *MemoryCache) Size() int64 {
	return int64(len(c.primes))
}

func (c *MemoryCache) All() []int64 {
	return c.primes
}
