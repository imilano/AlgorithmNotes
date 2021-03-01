package leet_153

import "math"

/*
Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

[4,5,6,7,0,1,2] if it was rotated 4 times.
[0,1,2,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums, return the minimum element of this array.
*/

//---------------------------
// 直接遍历
func findMin(nums []int) int {
	res := math.MaxInt16
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] < res {
			res = nums[i]
		}
	}

	return res
}

// ---------------------------
// 二分
// 如果旋转点在右半边，那么[mid]一定大于[right]，此时在有半边进行查找即可，否则就在左半边进行查找。
// TODO #二分 如何写出正确的二分
func findMinBinary(nums []int) int {
	length := len(nums)
	left, right := 0, length-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			//right = mid -1
			right = mid // TODO figure this out, why mid rather than mid -1
		}
	}
	return nums[left]
}

// divide and conquer
func findMinDC(nums []int) int {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, start int, end int) int {
	if nums[start] <= nums[end] {
		return nums[start]
	}

	mid := (start + end) / 2
	return min(helper(nums, start, mid), helper(nums, mid+1, end))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
