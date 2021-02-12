package leet_239

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	nums := [][]int{
		[]int{1, 3, -1, -3, 5, 3, 6, 7},
		[]int{1},
		[]int{1, -1},
		[]int{9, 11},
		[]int{4, 2},
	}
	ks := []int{3, 1, 1, 2, 2}

	for i := range nums {
		fmt.Println(ks[i], ": ", maxSlidingWindow(nums[i], ks[i]))
	}
}
