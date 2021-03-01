package leet_239

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	nums := [][]int{
		[]int{1, 3, -1, -3, 5, 3, 6, 7},
		[]int{1},
		[]int{1, -1},
		[]int{9, 11},
		[]int{4, 2},
		{7,2,4},
	}
	ks := []int{3, 1, 1, 2, 2,2}
	ans := [][]int{
		{3,3,5,5,6,7},
		{1},
		{1,-1},
		{11},
		{4},
		{7,4},
	}

	for i := range nums {
		r := maxSlidingWindowWithHeap(nums[i], ks[i])
		if !reflect.DeepEqual(r, ans[i]) {
			t.Errorf("test fail: index %d, want %+v, get %+v", i, ans[i], r)
		}
	}
}
