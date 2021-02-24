package leet_268

import "sort"

/*
Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.

Follow up: Could you implement a solution using only O(1) extra space complexity and O(n) runtime complexity?



Example 1:

Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.
*/

func missingNumber(nums []int) int {
	n := len(nums) + 1
	m := make([]bool, n)
	for i := range nums {
		m[nums[i]] = true
	}

	for i := 0; i <= n; i++ {
		if m[i] == false {
			return i
		}
	}

	return -1
}

// 利用等差数列求和公式，先求出0到n的和sum，再用sum减去nums中的数字之和，就是结果
func missingNumber2(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for i := range nums {
		sum -= nums[i]
	}

	return sum
}

// 使用位操作异或，将0到n的异或值和sum中的所有值进行异或，相同的数都被消去，最后留下不同的数
func missingNumber3(nums []int) int {
	var res int
	n := len(nums)
	for i := 1; i <= n; i++ {
		res ^= i ^ nums[i-1]
	}

	return res
}

// 使用二分查找
func missingNumber4(nums []int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
