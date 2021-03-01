package leet_53

import "math"

/*
	Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
	Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

// Dynamic Programming
func maxSubArray(nums []int) int {
	var res int
	max := math.MinInt64

	for i := 0; i < len(nums); i++ {
		res += nums[i]
		if res > max {
			max = res
		}

		if res <= 0 {
			res = 0
		}
	}

	return max
}
