package ch7_quick_sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arrs := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1, 5, 2, 3, 1, 4, 5},
		{99, -99, -100, 22},
	}

	for _, arr := range arrs {
		fmt.Printf("Before: %v\n", arr)
		QuickSort(arr, 0, len(arr)-1)

		fmt.Printf("After: %v\n", arr)
		fmt.Println("------------------------")
	}

}
