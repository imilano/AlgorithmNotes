package leet_34

/*
	Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.
	If target is not found in the array, return [-1, -1].
	Follow up: Could you write an algorithm with O(log n) runtime complexity?
 */

// time complexity: O(n)
// space complexity: O(1)
func brutalForce(nums []int, target int) []int {
	res := []int{-1,-1}
	if nums == nil {
		return  res
	}

	// find leftmost position
	for i := 0; i< len(nums);i++ {
		if nums[i] ==  target {
			res[0] = i
			break
		}
	}

	// we didn't find any element in above loop, so res[0] is -1
	if res[0] == -1 {
		return res
	}

	// find rightmost position
	for i := len(nums)-1; i>=0;i-- {
		if nums[i] == target {
			res[1] = i
			break
		}
	}

	return res
}

// Binary search
// time complexity: O(logn)
// space complexity: O(1)
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	start,end := 0,len(nums)-1
	res := []int{-1,-1}

	// find leftmost position
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] < target {
			start = mid +1
		} else {
			end = mid
		}
	}

	if nums[start] == target {
		res[0] = start
	}

	// find rightmost position
	end = len(nums)-1
	for start < end {
		mid := start + (end-start)/2 +1 // biased to right

		if target < nums[mid] {
			end  = mid -1
		} else {
			start = mid
		}
	}

	if nums[end] == target {
		res[1] = end
	}

	return res
}
