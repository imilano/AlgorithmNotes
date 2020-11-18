package leet_96


func numTrees(n int) int {
	dp := make([]int,n+1)
	dp[0],dp[1] =1,1

	for m := 2; m <= n;m++ {
		for i := 1;i<=m;i++ {
			dp[m] += dp[i-1] * dp[m-i]
		}
	}

	return dp[n]
}
