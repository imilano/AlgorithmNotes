package leet_260

/*
Given an integer array nums, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once. You can return the answer in any order.

Follow up: Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?



Example 1:

Input: nums = [1,2,1,3,2,5]
Output: [3,5]
Explanation:  [5, 3] is also a valid answer.
Example 2:

Input: nums = [-1,0]
Output: [-1,0]
Example 3:

Input: nums = [0,1]
Output: [1,0]
*/

// 最简单的方法当然是使用map了
// 或者也可以直接排个序来搞
func singleNumber(nums []int) []int {
	var res []int
	if len(nums) == 0 {
		return nil
	}

	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}

	return res
}

// TODO
// 如何做到O（n)时间复杂度以及常量空间复杂度呢
