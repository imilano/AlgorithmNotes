package leet_204

import (
	"fmt"
	"testing"
)

func TestCountPrime(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 10, 11}

	for _, v := range arr {
		fmt.Println(v, ": ", countPrimes(v))
	}
}
