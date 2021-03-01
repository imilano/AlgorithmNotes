package leet_263

import (
	"testing"
)

func TestIsUgly(t *testing.T) {
	nums := []int{1, 2, 3, 5, 6, 8, 14}
	ans := []bool{true, true, true, true, true, true, false}
	for i := range nums {
		r := isUgly2(nums[i])
		if ans[i] != r {
			t.Errorf("test fail: index %d, want %v, get %v", i, ans[i], r)
		}
	}
}
