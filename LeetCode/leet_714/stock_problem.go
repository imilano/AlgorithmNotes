package leet_714

// 动态规划，时间复杂度O(N),空间复杂度O(N).
// 有手续费，意味着买入的代价增加了，只需要在买入的时候减去手续费即可
func maxProfit(prices []int, fee int) int {
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
	// 第一次购买时间即产生手续费
	dp[0][0], dp[0][1] = 0, -prices[0]-fee

	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
	}

	return dp[length-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
