package leet_81

/*
	Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
	(i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).

	You are given a target value to search. If found in the array return true, otherwise return false.
*/

// passed solution
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		// 及时返回
		if nums[mid] == target {
			return true
		}
		if nums[left] < nums[mid] {
			if target >= nums[left] && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[right] {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// 以防出现[3,1,2,3,3,3,3]这样的情况。在这样的情况下，该数并不等于我们
			// 所需要的target，故而可以直接跳过，缩小搜寻范围，从而继续使用二分查找
			if nums[mid] == nums[left] {
				left++
			}

			if nums[mid] == nums[right] {
				right--
			}
		}
	}

	if nums[left] == target {
		return true
	}
	return false
}

// TODO can't figure out the solution, redo this problem.
func searchWrong(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}

		// 右侧有序
		if nums[mid] == nums[left] && nums[mid] == nums[right] {
			left++
			right--
		}
		if nums[mid] < nums[right] {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1 // why subtract 1
			}
		} else {
			// 左侧有序
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return false
}
