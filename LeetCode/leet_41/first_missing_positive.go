package leet_41

/*
	Given an unsorted integer array nums, find the smallest missing positive integer.
	Follow up: Could you implement an algorithm that runs in O(n) time and uses constant extra space.?
*/

// scan the array once, then put every element to its right place
// at last, the first place where its number is not right, return the place + 1
func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func firstMissingPositive(nums []int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		// can not replace while for if
		for nums[i] > 0 && nums[i] <= length && nums[nums[i]-1] != nums[i] {
			swap(nums, i, nums[i]-1)
		}
	}

	for i := 0; i < length; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return length + 1
}
