package leet_255

import (
	"testing"
)

func TestVerifyPreorder(t *testing.T) {
	nums := [][]int{
		[]int{5, 2, 6, 1, 3},
		[]int{5, 2, 1, 3, 6},
	}

	ans := []bool{false, true}

	for i := range nums {
		a := verifyPreorder2(nums[i])
		if a != ans[i] {
			t.Errorf("test fail: index %d, want %v, get %v", i, ans[i], a)
		}
	}
}
