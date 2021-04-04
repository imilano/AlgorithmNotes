package _6_连续子数组的最大和

import "math"

// 输入一个整型数组，数组里有正数也有负数。数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。要求时间复杂度为 O(n).

func FindGreatestSumOfSubArray( array []int ) int {
	// write code here
	res := math.MinInt32
	n := len(array)
	if n <= 0 {
		return res
	}

	var cur int
	for i := 0; i < n;i++ {
		if cur < 0 {
			cur = 0
		}

		cur +=  array[i]
		res = max(res,cur)
	}
	return res
}


func max(a,b int) int {
	if a < b {
		return b
	}

	return a
}