package leet_252

import "sort"

/*
	Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei),
	determine if a person could attend all meetings.
*/

func max (i,j int) int {
	if i > j {
		return i
		}
 	return j
}

func canAttend(nums [][]int) bool {
	length := len(nums)
	if nums == nil {
		return true
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	for i := 1; i < length; i++  {
		if nums[i-1][1] > nums[i][0] {
			return false
		}
	}

	return true
}