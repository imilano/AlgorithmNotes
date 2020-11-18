package leet_27

/*
	Given an array nums and a value val, remove all instances of that value in-place and return the new length.
	Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

	The order of elements can be changed. It doesn't matter what you leave beyond the new length.
	Clarification:
		Confused why the returned value is an integer but your answer is an array?
		Note that the input array is passed in by reference, which means a modification to the input array will be known to the caller as well.
 */

// two pointer
// time complexity: O(n)
// space complexity: O(1)
func removeElement(nums []int, val int) int {
	left,right := 0,0
	for right <len(nums) {
		if val != nums[right] {
			nums[left] = nums[right]
			left++
		}

		right++

	}

	return left
}

// swap the element to the last position, and decrease the length of array
func removeElementOnSwap(nums []int, val int) int {
	left := 0
	n := len(nums)
	for left < n{
		if nums[left] == val {
			nums[left] = nums[n-1]
			n--
		} else {
			left++
		}
	}

	return n
}
