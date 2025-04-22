package sieve

import (
	"log"
	"math"
)

type Sieve struct {
	logger      *log.Logger
	primesCache []int64
}

func NewSieve(log *log.Logger) Sieve {
	primesCache := make([]int64, 0)
	return Sieve{logger: log, primesCache: primesCache}
}

func (s *Sieve) estimateLimit(n int64) int64 {

	if n < 10 {
		s.logger.Print(" n is small < 10. hardcoding limit to 20")
		return int64(20)
	}

	// prime number theorem
	floatN := float64(n)
	return int64(floatN * math.Log(floatN) * 1.2)

}

// NthPrime returns the Nth prime number where  n >= 0
func (s *Sieve) NthPrime(n int64) (int64, error) {

	if n < 0 {
		s.logger.Printf("%v", ErrorInvalidInput)

		return 0, ErrorInvalidInput
	}

	if int64(len(s.primesCache)) > n {
		s.logger.Printf("fetching %dth prime from cache", n)

	}
	limit := s.estimateLimit(n)

	for {
		primes := s.GetPrimes(limit)

		if int64(len(primes)) > n {
			return primes[n], nil
		}
		//	if we are hitting this it means our limit wasn't high enough try again
		limit *= 2
	}

}

func (s *Sieve) GetPrimes(n int64) []int64 {

	// keep track of primes with slice of bools
	isPrime := make([]bool, n+1)

	for i := range isPrime {
		isPrime[i] = true
	}

	isPrime[0] = false
	isPrime[1] = false

	// iterate through n = 2 to the square root of n and mark multiples of n as non prime
	for num := int64(2); num*num <= n; num++ {
		if isPrime[num] {
			for i := num * num; i <= n; i += num {
				isPrime[i] = false
			}
		}
	}
	result := make([]int64, 0)

	for i, val := range isPrime {
		if val {
			result = append(result, int64(i))
		}
	}
	s.primesCache = result
	return result
}
