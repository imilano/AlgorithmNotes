package leet_280

import (
	"sort"
)

/*
Given an unsorted array nums, reorder it in-place such that nums[0] <= nums[1] >= nums[2] <= nums[3]....

Example:

Input: nums = [3,5,2,1,6,4]
Output: One possible answer is [3,5,1,6,2,4]
*/

// 先将数组排序，然后第二个和第三个交换，第四个和第五个交换，依次持续
func wiggleSort(nums []int) {
	sort.Ints(nums)
	n := len(nums)
	if n <= 2 {
		return
	}

	i := 2
	for i < n {
		nums[i],nums[i-1]=nums[i-1],nums[i]
		i+=2
	}

	return
}
