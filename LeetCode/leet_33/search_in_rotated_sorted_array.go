package leet_33

/*
	You are given an integer array nums sorted in ascending order, and an integer target.
	Suppose that nums is rotated at some pivot unknown to you beforehand (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
	If target is found in the array return its index, otherwise, return -1.
*/

//func searchOriginal(nums []int, target int) int {
//	var pivot int
//	if len(nums) == 0 {
//		return -1
//	}
//
//	for i := 0 ; i< len(nums)-1;i++ {
//		if nums[i] > nums[i+1] {
//			pivot = i
//			break
//		}
//	}
//
//	if target > nums[0] &&target < nums[pivot]
//}

// 采用二分查找
// 对于所有可以列出来的排序，我们取其中间数mid，如果mid小于最右边的数，则右半段是有序的，如果mid大于右半段的数，那么左半段是有序的（可以通过观察得到）。
// 那么，只需要在有序半段内，用首尾两个数组来判断目标值是否在这个区域内，即可确定保留哪半边
// TODO check it twice
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		// 右侧有序
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

	return -1
}
