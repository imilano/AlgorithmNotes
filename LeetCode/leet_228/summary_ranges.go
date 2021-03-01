package leet_228

import "strconv"

/*
You are given a sorted unique integer array nums.

Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.

Each range [a,b] in the list should be output as:

"a->b" if a != b
"a" if a == b


Example 1:

Input: nums = [0,1,2,4,5,7]
Output: ["0->2","4->5","7"]
Explanation: The ranges are:
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
*/

// 一次遍历，每次检查下一个数是不是递增的，如果是，则继续向下遍历；如果不是，则需要判断此时是一个数还是一个序列，如果是一个数，直接加入结果，
// 如果是一个序列，则需要添加起始位置。
func summaryRanges(nums []int) []string {
	var res []string
	length := len(nums)
	if length == 0 {
		return res
	}

	index := 0
	for index < length {
		n := 1
		for index+n < length && nums[index+n]-nums[index] == n {
			n++
		}

		if n == 1 {
			res = append(res, strconv.Itoa(nums[index]))
		} else {
			res = append(res, strconv.Itoa(nums[index])+"->"+strconv.Itoa(nums[index+n-1]))
		}
		index = index + n
	}

	return res
}
