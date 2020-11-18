package leet_300

/*
	Given an unsorted array of integers, find the length of longest increasing subsequence.
*/

/*
	注意区分子串和子序列。子串需要连续，而子序列不需要连续

Solution 1： Dynamic Programming
	时间复杂度O(n^2)，空间复杂度O(n)
	假设dp[i]表示以字符s[i]结尾的子序列的长度；那么对于非空的数组，都有dp[i]等于1，这是初始条件。那么如何计算dp[i]呢？比较明显的是，因为我们需要的递增子序列，
	而该子序列并不连续，所以当求下标j的dp值时，必然涉及到对原数组中j位置以前的元素都进行遍历，以找出比s[j]小的元素。这难免会导致O(N^2)的时间复杂度。
*/

// Dynamic Programming
func getMax(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func lengthOfLIS(nums []int) int {
	if nums == nil {
		return 0
	}

	l := len(nums)
	dp := make([]int, l)
	var res int

	// dp数组的初始长度应该为1，因为非空数组的LIS长度至少为1
	for i := 0; i < l; i++ {
		dp[i] = 1
	}

	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			// 找到以比i位置小的元素结尾的子串，然后把i位置的数字接到该子串上去
			if nums[i] > nums[j] {
				dp[i] = getMax(dp[i], dp[j]+1)
			}
		}
	}

	for i := 0; i < l; i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}


