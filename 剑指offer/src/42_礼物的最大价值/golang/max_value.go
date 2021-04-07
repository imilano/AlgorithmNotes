package golang

// 简单的动态规划
func MaxValue(gifts [][]int) int {
	n := len(gifts)
	dp := make([][]int,n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n;i++ {
		for j := 1; j <= n;j++ {
			dp[i][j] = gifts[i-1][j-1]
		}
	}

	for i := 1; i <= n; i++ {
		for j:= 1; j <= n;j++ {
			dp[i][j] = max(dp[i-1][j],dp[i][j-1]) + gifts[i-1][j-1]  // 状态转移方程
		}
	}

	return dp[n][n]
}


func max(a,b int)int {
	if a > b {
		return a
	}

	return b
}