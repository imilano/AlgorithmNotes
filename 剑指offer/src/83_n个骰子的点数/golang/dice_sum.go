package golang

import "math"

// 把 n 个骰子扔在地上，求点数和为 s 的概率。
//定义状态方程 dp[i][j] = dp[i-1][j-k]，其中dp[i][j]表示掷i个色子后，会出现和为j的次数。第i次的结果跟i-1次的结果相关，第i次掷完之后，出现
//j和为j的次数等于dp[i-1][j-1] + dp[i-1][j-2] + ... + dp[i-1][j-6]。

//初始状态，dp[1][1] = dp[1][2] = ... = dp[1][6] = 1
func diceSum(n int) []int{
	var res  []int
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int,n * 6)
	}

	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}

	for i := 2; i <= n;i++ {
		for j := i; j <= 6*i;j++ {
			for k := 1; k <= 6; k++ {
				if j - k <= 0 {
					break
				}

				dp[i][j] += dp[i-1][j-k]
			}
		}
	}

	all := int(math.Pow(6,float64(n)))
	for i := n; i <= 6*n;i++ {
		res = append(res, dp[n][i]/all)
	}

	return res
}
