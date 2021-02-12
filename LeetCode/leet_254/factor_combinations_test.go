package leet_254

import (
	"fmt"
	"testing"
)

func TestGetFactors(t *testing.T) {
	nums := []int{2, 3, 4, 8, 9, 10, 12, 20, 32, 40}
	for _, v := range nums {
		fmt.Println(v, ": ", getFactors(v))
	}
}
