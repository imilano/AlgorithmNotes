package leet_259

import (
	"sort"
)

/*
Given an array of n integers nums and a target, find the number of index triplets i, j, k with 0 <= i < j < k < n that satisfy the condition nums[i] + nums[j] + nums[k] < target.

Example:

Input: nums = [-2,0,1,3], and target = 2
Output: 2
Explanation: Because there are two triplets which sums are less than 2:
             [-2,0,1]
             [-2,0,3]
Follow up: Could you solve it in O(n2) runtime?
*/

// 穷举法
func threeSumSmaller(nums []int, target int) int {
	var res int
	n := len(nums)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] < target {
					res++
				}
			}
		}
	}

	return res
}

func threeSumSmaller2(nums []int, target int) int {
	var res int
	n := len(nums)
	if n < 3 {
		return 0
	}
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		l, r := i+1, n-1
		for l < r {
			if nums[i]+nums[l]+nums[r] < target {
				res += r - l // 如果i+l的值加上r的值小于target，那么加上一个比r更小的值（注意数组已排序），也能小于target
				l++
			} else {
				r--
			}
		}
	}

	return res
}
