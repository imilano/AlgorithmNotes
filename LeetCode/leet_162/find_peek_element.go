package leet_162

/*
A peak element is an element that is strictly greater than its neighbors.

Given an integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.

You may imagine that nums[-1] = nums[n] = -∞.
*/

func findPeakElement(nums []int) int {
	length := len(nums)
	left, right := 0, length-1
	for left < right {
		if nums[left] < nums[right] {
			left++
		} else {
			right--
		}
	}

	return left
}
