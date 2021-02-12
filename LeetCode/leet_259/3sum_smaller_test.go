package leet_259

import "testing"

func TestThreeSumSmaller(t *testing.T) {
	nums := [][]int{
		[]int{-2, 0, 1, 3},
	}
	targets := []int{2}
	ans := []int{2}
	for i := range ans {
		res := threeSumSmaller(nums[i], targets[i])
		if res != ans[i] {
			t.Errorf("test fail: want %d, get %d", ans[i], res)
		}
	}

}
