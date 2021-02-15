package leet_265

import (
	"testing"
)

func TestMinCostII(t *testing.T) {
	nums := [][][]int{
		[][]int{
			[]int{1,5,3},
			[]int{2,9,4},
		},
	}
	ans := []int{5}

	for i := range nums {
		cost := minCostII(nums[i])
		if cost != ans[i] {
			t.Errorf("test fail: index %d, want %d, get %d",i,ans[i],cost)
		}
	}
}
