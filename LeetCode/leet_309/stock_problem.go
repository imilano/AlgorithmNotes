package leet_309

// 动态规划，时间复杂度O(N),空间复杂度O(N).
// 有一天的冷冻期，意味着买入时候，只能从i-2开始状态转移，而不是i-1
func maxProfit(prices []int) int {
	length := len(prices)

	if length < 1 {
		return 0
	}

	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}

	// base condition
	// 分别对应不持有股票和持有股票两种情况
	dp[0][0], dp[0][1] = 0, -prices[0]

	var tmp int
	for i := 1; i < length; i++ {
		if i < 2 {
			tmp = 0
		} else {
			tmp = dp[i-2][0]
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], tmp-prices[i])
	}

	return dp[length-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
