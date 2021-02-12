package leet_239

/*
You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.



Example 1:

Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Example 2:

Input: nums = [1], k = 1
Output: [1]
Example 3:

Input: nums = [1,-1], k = 1
Output: [1,-1]
Example 4:

Input: nums = [9,11], k = 2
Output: [11]
Example 5:

Input: nums = [4,-2], k = 2
Output: [4]
*/

//WRONG
// TODO
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	var curMax int

	n := len(nums)
	if n < k {
		res = append(res, findMax(nums))
	} else { // if n > k
		curMax = findMax(nums[:k])
		res = append(res, curMax)

		for i := k; i < n; i++ {
			if curMax < nums[i] {
				curMax = nums[i]
			}

			res = append(res, curMax)
		}
	}

	return res
}

func findMax(nums []int) int {
	t := nums[0]
	for _, v := range nums {
		if v > t {
			t = v
		}
	}

	return t
}
