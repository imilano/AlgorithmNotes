package leet_152

import "math"

/*
Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

Example 1:

Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/

//---------------------------------------------
// Original wrong solution
func maxProductOriginal(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
	}
	for i := 0; i < length; i++ {
		dp[i][i] = nums[i]
	}

	max := math.MinInt32
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if j-1 >= 0 && j != length-1 {
				dp[i][j] = dp[i][j-1] * nums[j]
			}
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}

	return max
}

//------------------------------------------------
// dynamic solution
//  使用两个数组f(i)和g(i)，f(i)代表0到i（包括i）范围内的最大子数组乘积，g(i)代表0到i（包括i）范围内的最小子数组乘积。为什么还要记录最小子数组呢？
// 可以看到，数组中的负数和0增加了题目的复杂度。当遍历到当前数字时，如果当前数字是负数，那么此时的最大值就可能是之前的最小负值乘以当前的负值。也就是说，
// 最大值可能从三个地方产生，nums[i],f[i-1]*nums[i],g[i-1]*nums[i]
func maxProduct(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	f, g := make([]int, length), make([]int, length)
	f[0] = nums[0]
	g[0] = nums[0]

	//you can not just initialize res to 0
	//var res int
	res := nums[0]
	for i := 1; i < length; i++ {
		f[i] = max(max(f[i-1]*nums[i], g[i-1]*nums[i]), nums[i])
		g[i] = min(min(f[i-1]*nums[i], g[i-1]*nums[i]), nums[i])

		res = max(f[i], res)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

//----------------------------------
// more concise way
func maxProductConcise(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	curMin, curMax, res := nums[0], nums[0], nums[0]
	for i := 1; i < length; i++ {
		// if negative, swap min and max
		if nums[i] < 0 {
			curMin, curMax = curMax, curMin
		}

		curMax = max(curMax*nums[i], nums[i])
		curMin = min(curMin*nums[i], nums[i])
		res = max(curMax, res)
	}

	return res
}

//--------------------------------
// iterate two time
// 一次正向遍历，一次反向遍历；正向遍历用于建立累乘数组，每次用最大值更新res，需要注意的是，当遇到0时，需要把prod设为1，方便下一次的累乘计算。
// 那为什么还需要反向遍历呢？主要是如果负数个数为奇数的话，正向遍历可能会确实最大值，比如[-1,-2,-3]的累积分别是[-1,2,-6]从而得出res=2.这是不对的。
// 所以当负数个数为奇数时，出现在最前或最后的奇数可能会被是组成最大积的机会，遍历两次就不会出现漏乘的情况。
func maxProduct2Times(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	res, prod := nums[0], 1
	for i := 0; i < length; i++ {
		prod = nums[i] * prod
		res = max(res, prod)
		if nums[i] == 0 {
			prod = 1
		}
	}

	prod = 1
	for i := length - 1; i >= 0; i-- {
		prod = nums[i] * prod
		res = max(res, prod)
		if nums[i] == 0 {
			prod = 1
		}
	}

	return res
}
