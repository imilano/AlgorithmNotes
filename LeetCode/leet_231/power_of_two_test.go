package leet_231

import (
	"fmt"
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	nums := []int{
		1, 16, 3, 4, 5, 536870912,
	}

	for _, v := range nums {
		fmt.Printf("%d: %v\n", v, isPowerOfTwo(v))
	}
}

func TestIsPowerOfTwo3(t *testing.T) {
	nums := []int{16}
	for _, v := range nums {
		fmt.Println(v, ": ", isPowerOfTwo3(v))
	}
}
