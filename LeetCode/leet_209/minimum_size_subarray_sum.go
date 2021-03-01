package leet_209

import "math"

/*
Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

Example:

Input: s = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: the subarray [4,3] has the minimal length under the problem constraint.
Follow up:
If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n).
*/

//
//func minSubArrayLen(s int, nums []int) int {
//	var res int
//	var cur int
//
//	n := len(nums)
//	left,right := 0,n-1
//
//	for i := 0; i < n;i++ {
//		cur += nums[i]
//		left,right = i-1,i+1
//		for left >=0 && right <= n-1 && cur < s{
//			if nums[left] > nums[right] {
//				cur += nums[left]
//				left--
//			} else {
//				cur += nums[right]
//				right++
//			}
//		}
//
//		if cur > s {
//			res = min(res, right-left)
//		}
//		cur = 0
//	}
//
//	return res
//}

func minSubArrayLen(s int, nums []int) int {
	res := math.MaxInt32
	var left, sum int

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for left <= i && sum >= s {
			res = min(res, i-left+1)
			sum -= nums[left]
			left++
		}
	}

	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func minSubArrayLen2(s int, nums []int) int {
	var left, right, sum int
	n := len(nums)
	res := n + 1

	for right < n {
		for sum < s && right < n {
			sum += nums[right]
			right++
		}

		for sum >= s {
			res = min(res, right-left)
			sum -= nums[left]
			left++
		}
	}

	if res == n+1 {
		return 0
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
