package lab5

import (
	"fmt"
	"testing"
)

func TestKnapsackGreedy(t *testing.T) {
	var maxValue int
	var its items

	data := []struct {
		weights  []int
		values   []int
		capacity int
	}{
		{[]int{2, 2, 6, 5, 4}, []int{6, 3, 5, 4, 6}, 10}, // 最优解 {[2,6],[2,3],[4,6]}, maxValue = 15
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, 10}, // 最优解，顺序选择
		{[]int{10, 12, 13}, []int{10, 3, 2}, 10},
	}

	for _, d := range data {
		maxValue, its = KnapsackGreedy(d.capacity, d.weights, d.values)
		fmt.Println("MaxValue With Greedy: ", maxValue)

		maxValue, _ = KnapsackDP(d.capacity, d.weights, d.values)
		fmt.Println("MaxValue With DP: ", maxValue)

		fmt.Println(its)
		fmt.Println("---------------------------------")
	}
}
