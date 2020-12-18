package ch6_heap_sort

// Using max heap for demonstration
// Suppose index start from 1
func Insert(nums []int, x int) []int {
	res := make([]int, len(nums)+1)
	copy(res, nums)
	res[len(nums)+1] = x

	res[1], res[len(nums)+1] = res[len(nums)+1], res[1]
	MaxHeapify(res, 1)
	return res
}

func Maximum(nums []int) int {
	return nums[1]
}

// 去掉并返回nums中最大的元素
func ExtractMax(nums []int) int {
	var res int
	length := len(nums)

	if length < 1 {
		// error
	}

	res = nums[1]
	nums[1], nums[len(nums)] = nums[len(nums)], nums[1]

	nums = nums[:len(nums)]
	nums[len(nums)] = 0

	MaxHeapify(nums, 1)
	return res
}

func IncreaseKey(nums []int, i int, k int) {
	if nums[i] < k {
		// error, new key is smaller then the current key
	}

	nums[i] = k
	// adjust upward
	if i > 1 && nums[i/2] < nums[k] {
		nums[i], nums[i/2] = nums[i/2], nums[i]
		i = i / 2
	}
}
