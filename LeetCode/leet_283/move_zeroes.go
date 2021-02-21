package leet_283

/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
*/

func moveZeroes(nums []int) {
	var count int
	var slow, fast int

	n := len(nums)
	for fast < n {
		for fast < n && nums[fast] == 0 {
			count++
			fast++
		}

		if fast < n {
			nums[slow] = nums[fast]
			slow++
			fast++
		}
	}

	for i := n - count; i < n; i++ {
		nums[i] = 0
	}

	return
}

func moveZeroes2(nums []int) {
	n := len(nums)
	for fast,slow := 0,0; fast < n;fast++{
		//if nums[fast] > 0 {
		if nums[fast] != 0 {
			nums[fast],nums[slow] = nums[slow],nums[fast]  // 交换的过程其实就是不断的把0往后移动的过程
			slow++
		}
	}
}
