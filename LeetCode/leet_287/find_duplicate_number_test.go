package leet_287

import "testing"

func TestFindDuplicate(t *testing.T) {
	nums := [][]int{
		{1, 3, 4, 2, 2},
		{3, 1, 3, 4, 2},
	}
	ans := []int{2, 3}

	for i := range nums {
		if ans[i] != findDuplicate3(nums[i]) {
			t.Errorf("test fail: index %d, want %d, get %d", i, ans[i], findDuplicate(nums[i]))
		}
	}
}
