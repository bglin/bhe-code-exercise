package main

import (
	"log"
	"ssse-exercise-sieve/pkg/sieve"
)

type Siever interface {
	NthPrime(n int64) (int64, error)
}

type App struct {
	SieveProvider Siever
	log           *log.Logger
}

func NewApp(log *log.Logger, sieveProvider Siever) *App {
	return &App{sieveProvider, log}
}

func main() {

	logger := log.Default()

	logger.Printf("intializing app instance")
	defer logger.Print("finished calculation of Nth prime number")

	sieveProvider := sieve.NewSieve(logger)

	app := NewApp(logger, sieveProvider)

	if app == nil {
		logger.Fatal("App creation failed")
	}
	logger.Print()
	n := int64(19)
	primeNum, err := app.SieveProvider.NthPrime(n)

	if err != nil {
		logger.Fatalf("%v", err)
	}

	logger.Printf("%dth prime is %v", n, primeNum)

}
