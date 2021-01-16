package leet_90

import "sort"

/*
	Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).
	Note: The solution set must not contain duplicate subsets.
*/

func backtrace(res *[][]int, nums []int, chosen []int, start int) {
	tmp := make([]int, len(chosen))
	copy(tmp, chosen)
	*res = append(*res, tmp)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		chosen = append(chosen, nums[i])
		backtrace(res, nums, chosen, i+1)
		chosen = chosen[:len(chosen)-1]
	}
}

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}

	sort.Ints(nums)
	backtrace(&res, nums, nil, 0)
	return res
}
