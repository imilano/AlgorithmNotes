/*
	Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
	You may assume that each input would have exactly one solution, and you may not use the same element twice.
	You can return the answer in any order.
*/
package leet_01

// 暴力枚举
func brutalForce(nums []int, target int) []int {
	if nums == nil {
		return nil
	}

	for i := 0;i<len(nums);i++ {
		adds := target - nums[i]
		for j := i+1; j<len(nums) ;j++ {
			if nums[j] == adds {
				return []int{i,j}
			}
		}
	}
	return nil
}

// 从前向后遍历，提前占位
func Pri(nums []int, target int) []int {
	m := make(map[int]int)
	for k,v := range nums {
		_,ok := m[v]
		if !ok {
			m[target-v] = k
		} else {
			return []int{m[v],k}
		}
	}

	return nil
}

func twoSum(nums []int, target int) []int {
	return Pri(nums,target)
}
