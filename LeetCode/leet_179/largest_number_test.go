package leet_179

import (
	"fmt"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	a := [][]int{
		[]int{10, 2},
		[]int{3, 30, 34, 59},
	}

	for _, v := range a {
		fmt.Println(largestNumber(v))
	}

}
