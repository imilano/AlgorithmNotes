/*
	Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
	You may assume that each input would have exactly one solution, and you may not use the same element twice.
	You can return the answer in any order.
*/
package leet_01

// 暴力枚举
// time complexity: O(n^2)
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


// 使用map，从前向后遍历nums，对于nums中每个num，检查map中key为num的记录是否存在，如果不存在，则map[target-num]=indexOfNum，记录当前的下标值；
// 否则如果存在，则找到了一个结果，直接返回即可。
// time complexity: O(n)
// space complexity: O(n)
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

// 双指针法
// time complexity: O(nlogn)
// space complexity: O(1)
// WRONG，由于需要返回的是元素下标，而排序已经改变了元素的下标。

//func twoSum2(nums []int, target int) []int {
//	sort.Ints(nums)
//	left,right := 0,len(nums) -1
//	for left < right {
//		if nums[left] + nums[right] == target {
//			return []int{left,right}
//		}
//
//		left++
//		right--
//	}
//
//	return nil
//}
