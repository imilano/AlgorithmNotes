package leet_quicksort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{1, 4, 2, 3, 5, 2, 6, 7, 2, 76}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
