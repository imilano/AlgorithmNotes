package leet_219

/*
Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true
Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false
*/

// 采用希尔排序的思想
func containsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)

	for gap := k; gap > 0; gap-- {
		for i := gap; i < n;i++ {
			if nums[i] == nums[i-gap] {
				return true
			}
		}
	}

	return false
}
