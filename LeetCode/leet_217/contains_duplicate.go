package leet_217

import "sort"

/*
Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.

Example 1:

Input: [1,2,3,1]
Output: true
Example 2:

Input: [1,2,3,4]
Output: false
*/

//---------------------------------------
// 先对数组进行排序，然后比较数字的相邻元素
func containsDuplicate(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	sort.Ints(nums)
	for i := 0; i< n-1;i++{
		if nums[i] ==nums[i+1] {
			return true
		}
	}

	return false
}


//----------------------------------------
// 使用map
func containsDuplicate2(nums []int) bool {
	m := make(map[int]bool)

	for _,v := range nums {
		if _,ok := m[v]; ok {
			return true
		}
		m[v] = true
	}

	return false
}