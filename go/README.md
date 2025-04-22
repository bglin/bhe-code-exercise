# Prime Number Sieve


## Run

```go
go run main.go
```
## Tests

```go
go test ./...
```

## Notes

### Sieve

Package sieve returns a concrete struct with 2 exported methods

```go
type Sieve struct {
	logger *log.Logger
}

func (s Sieve) NthPrime(n int64) (int64, error) {...}


func (s Sieve) GetPrimes(n int64) []int64 {...}


```

I opted to have the package return a concrete struct instead of an
interface to avoid coupling between the producer and consumer. 

I define an interface `Siever` on the consumer side in `main.go`

```go
type Siever interface {
	NthPrime(n int64) (int64, error)
}

```

Since go uses implicit interfaces  type Sieve will automatically implement  `Siever`  since it 
defines the method NthPrime in `Siever`. Using this approach I can keep interfaces small and only define methods
that I need. From the producer side, I can add new methods and functionality without breaking existing code.



### Cache

Package cache returns is in memory cache implementation. Sieve abstracts the cache via interface

```go
type PrimesCache interface {
	Get(n int64) (int64, bool)
	Set(primes []int64)
}
```
