package ch8_linear_sort

import (
	"fmt"
	"testing"
)

func TestCountingSort(t *testing.T) {
	arrs := [][]int{
		{95, 94, 91, 98, 99, 90, 99, 93, 91, 92},
		{1, 2, 3, 4, 5, 9, 3, 2, 1},
		{900, 300, 200, 400},
	}

	for _, arr := range arrs {
		max, min := findMaxAndMin(arr)
		fmt.Printf("Before: %v\n", arr)

		arr = countingSort(arr, max, min)
		fmt.Printf("After: %v\n", arr)
		fmt.Println("------------------------")
	}
}
