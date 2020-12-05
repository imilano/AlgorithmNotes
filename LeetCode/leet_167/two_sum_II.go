package leet_167

/*
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

Your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution and you may not use the same element twice.
*/


func twoSum(numbers []int, target int) []int {
	var res []int
	length := len(numbers)
	if length < 2 {
		return res
	}

	left,right := 0,length-1
	for left < right {
		num :=  numbers[left] + numbers[right]
		if num == target {
			res = append(res,left+1,right+1)
			return res
		}else if num < target {
			left++
		} else {
			right--
		}
	}

	return res
}
