package leet_256

import "testing"

func TestMinCost(t *testing.T) {
	nums := [][][]int{
		nil,
		[][]int{},
		[][]int{
			[]int{17, 2, 17},
			[]int{16, 16, 5},
			[]int{14, 3, 9},
		},
	}

	ans := []int{0, 0, 10}

	for i := range ans {
		res := minCost2(nums[i])
		if res != ans[i] {
			t.Errorf("test fail: want %d, get %d", ans[i], res)
		}
	}
}
