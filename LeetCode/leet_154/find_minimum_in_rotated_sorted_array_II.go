package leet_154

/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).

Find the minimum element.

The array may contain duplicates.

Example 1:

Input: [1,3,5]
Output: 1
*/

//----------------------------
// 问题在于出现重复值的情况，解决办法：当nums[left] = nums[mid]时，将right--,或者将left++，也就是跳开一个重复值，从而可以对剩下的区间继续使用二分查找
func findMin(nums []int) int {
	length := len(nums)
	left, right := 0, length-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] { // 剪去左半边
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
			// left++
		}
	}

	return nums[left]
}

//-----------------------------------------------
// divide and conquer
func findMinDC(nums []int) int {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, start int, end int) int {
	if nums[start] < nums[end] {
		return nums[start]
	}

	if start == end {
		return nums[start]
	}

	mid := (start + end) / 2
	return min(helper(nums, start, mid), helper(nums, mid+1, end))
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
