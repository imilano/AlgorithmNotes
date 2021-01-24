package leet_quicksort

func quickSort(nums []int, left, right int) {
	if left < right {
		pivot := partition(nums, left, right)
		quickSort(nums, left, pivot-1)
		quickSort(nums, pivot+1, right)
	}
}

// 交换法
func partition2(nums []int, left, right int) int {
	l, r := left, right
	pivot := nums[left]

	for l < r {
		for nums[r] >= pivot && l < r { // 不能缺少等于号
			r--
		}

		for nums[l] <= pivot && l < r { // 不能缺少等于号
			l++
		}

		nums[l], nums[r] = nums[r], nums[l]
	}

	// 为什么是left而不是right，因为循环直到left== right 的时候才结束，而此时left左边的元素都比pivot小，而right右边的元素都比pivot大,
	// 此时我们将pivot与left交换，交换后仍然能够保证left左边的元素都比left（也就是pivot）小
	nums[l], nums[left] = nums[left], nums[l]
	return l
}

// 填充法
func partition(nums []int, left, right int) int {
	l, r := left, right
	pivot := nums[left]

	for l < r {
		for l < r && nums[r] >= pivot { // 不能缺少等于号
			r--
		}
		if l < r {
			nums[l] = nums[r]
			l++
		}

		for l < r && nums[l] <= pivot { // 不能缺少等于号
			l++
		}
		if l < r {
			nums[l] = nums[r]
			r--
		}
	}

	nums[l] = pivot
	return l
}
