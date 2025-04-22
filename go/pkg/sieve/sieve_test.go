package sieve_test

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	sv "ssse-exercise-sieve/pkg/sieve"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNthPrime(t *testing.T) {
	logger := log.Default()
	sieve := sv.NewSieve(logger)

	fmt.Println("starting tests")

	tests := []struct {
		name     string // Added name field for the test case
		index    int64
		expected int64
		err      error
	}{
		{"Negative index", -1, 0, sv.ErrorInvalidInput},
		{"Test case 0", 0, 2, nil},
		{"Test case 19", 19, 71, nil},
		{"Test case 99", 99, 541, nil},
		{"Test case 500", 500, 3581, nil},
		{"Test case 986", 986, 7793, nil},
		{"Test case 2000", 2000, 17393, nil},
		{"Test case 1000000", 1000000, 15485867, nil},
		{"Test case 10000000", 10000000, 179424691, nil},
		{"Test case 100000000", 100000000, 2038074751, nil}, // not required, just a fun challenge
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := sieve.NthPrime(test.index)
			if !errors.Is(err, test.err) {
				t.Errorf("NthPrime(%d) returned error: %v, expected: %v", test.index, err, test.err)
			}
			assert.Equal(t, test.expected, result)
		})
	}
}

func FuzzNthPrime(f *testing.F) {
	logger := log.Default()
	sieve := sv.NewSieve(logger)

	f.Fuzz(func(t *testing.T, n int64) {
		result, err := sieve.NthPrime(n)
		if err != nil {
			t.Errorf("NthPrime(%d) returned error: %v, expected nil", n, err)
		}
		if !big.NewInt(result).ProbablyPrime(0) {
			t.Errorf("the sieve produced a non-prime number at index %d", n)
		}
	})
}

func TestCachedPrimes(t *testing.T) {
	logger := log.Default()

	sieve := sv.NewSieve(logger)

	firstCall := func() {
		defer duration(track("Call to Nth prime without cached value"))
		sieve.NthPrime(10000000)

	}

	secondCall := func() {
		defer duration(track("Call to Nth prime with cached value"))
		sieve.NthPrime(10000000)
	}

	firstCall()
	secondCall()
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
