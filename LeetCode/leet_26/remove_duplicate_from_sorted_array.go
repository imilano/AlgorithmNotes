package leet_26

/*
	Given a sorted array nums, remove the duplicates in-place such that each element appears only once and returns the new length.
	Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

	Clarification:
		Confused why the returned value is an integer but your answer is an array?
		Note that the input array is passed in by reference, which means a modification to the input array will be known to the caller as well.
 */

// two pointer
// use left and right, right runs faster than left, when arr[left] != arr[right], then copy right to left and increase both left and right;
// when arr[left] == arr[right], then we increase right; in the end, return i+1
// time complexity: O(n)
// space complexity: O(1)
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	left, right := 0, 1
	for right < len(nums) {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
		right++
	}

	return left +1
}

func removeDuplicates2(nums []int) int {
	var pre,cur int
	n := len(nums)
	if n == 0 {
		return 0
	}

	for cur < n {
		if nums[cur] != nums[pre] {
			pre++
			nums[pre] = nums[cur]
		}

		cur++
	}
	return pre+1
}


func removeDuplicates3(nums []int) int {
	var i int
	for _,num := range nums {
		if i < 1 || num > nums[i-1] {
			nums[i] = num
			i++
		}
	}


	return i
}