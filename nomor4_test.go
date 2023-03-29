package jastis

import (
	"fmt"
	"math/big"
	"testing"
)

func multiplesOf9(value int) bool {
	if value%9 == 0 {
		return true
	}

	return false
}

func checkNumber(number []int) {
	count := 1
	for _, value := range number {
		prime := big.NewInt(int64(value))
		if prime.ProbablyPrime(0) {
			fmt.Printf("Bilangan Prima: %d\n", value)
		} else if multiplesOf9(value) {
			fmt.Printf("Kelipatan 9 ke-%d: %d\n", count, value)
			count++
		} else {
			fmt.Println(value)
		}
	}
}

func TestIsPrime(t *testing.T) {
	var number []int
	for i := 1; i <= 100; i++ {
		number = append(number, i)
	}
	checkNumber(number)
}
