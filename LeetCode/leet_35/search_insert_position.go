package leet_35

/*
	Given a sorted array of distinct integers and a target value, return the index if the target is found.
	If not, return the index where it would be if it were inserted in order.
*/

// original solution
func searchInsertOriginal(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}

func searchInsert(nums []int, target int) int {
	if nums[len(nums)-1] < target {
		return len(nums)
	}

	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
