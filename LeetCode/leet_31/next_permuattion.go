package leet_31

/*
	Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.
	If such an arrangement is not possible, it must rearrange it as the lowest possible order (i.e., sorted in ascending order).
	The replacement must be in place and use only constant extra memory.
*/

// time complexity: O(n)
// space complexity: O(1)
func nextPermutation(nums []int) {
	start := len(nums) - 2

	// find the index start, where nums[start+1] > nums[start]
	for start >= 0 && nums[start+1] <= nums[start] {
		start--
	}

	// if start exist
	if start >= 0 {
		// find the index end, where nums[end] is just bigger than nums[start]
		end := len(nums) - 1
		for end > start && nums[end] <= nums[start] {
			end--
		}

		nums = swap(nums, start, end)
	}

	nums = reverse(nums, start+1)
}

func swap(nums []int, i, j int) []int {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp

	return nums
}

func reverse(nums []int, start int) []int {
	end := len(nums) - 1
	for start < end {
		tmp := nums[start]
		nums[start] = nums[end]
		nums[end] = tmp

		start++
		end--

	}

	return nums
}
