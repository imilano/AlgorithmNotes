package leet_198

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 递归解法，超时
func dp(nums []int, start int) int {
	if start >= len(nums) {
		return 0
	}

	res := max(
		dp(nums, start+1),             // 不抢，去下一家
		dp(nums, start+2)+nums[start], // 抢，去下下家
	)

	return res
}

// 递归解法中有重叠子问题，使用备忘录
func dpWithMemo(nums, memo []int, start int) int {
	if start >= len(nums) {
		return 0
	}

	if memo[start] != -1 {
		return memo[start]
	}

	res := max(
		dp(nums, start+1),             // 不抢，去下一家
		dp(nums, start+2)+nums[start], // 抢，去下下家
	)

	memo[start] = res
	return res
}

// 自底向上
func dpFromBottom2Top(nums []int) int {
	length := len(nums)

	// base condition: dp[n] = 0
	// dp[n]=m，表示从第n间房子开始抢劫，可以得到的最大收益是m
	dp := make([]int, length+2)

	for i := length - 1; i >= 0; i-- {
		dp[i] = max(dp[i+1], nums[i]+dp[i+2])
	}

	return dp[0]
}

// dp[i]之和它的后两个状态相关
func dpFromBottom2Up2(nums []int) int {
	length := len(nums)

	N1, N2, N := 0, 0, 0
	for i := length - 1; i >= 0; i-- {
		N = max(N1, N2+nums[i])
		N2, N1 = N1, N
	}

	return N
}

func rob(nums []int) int {
	// recursive solution
	// return dp(nums,0)

	// memo
	//memo := make([]int,len(nums))
	//for i,_ := range memo {
	//	memo[i] = -1
	//}
	//
	//return dpWithMemo(nums,memo,0)

	// From bottom up
	// return dpFromBottom2Top(nums)

	// bottom2up optomized
	return dpFromBottom2Up2(nums)
}
