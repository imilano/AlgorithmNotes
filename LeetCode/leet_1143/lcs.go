package leet_1143

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 递归方法，超时
func recLCS(s1, s2 string) int {
	l1, l2 := len(s1), len(s2)
	if l1 == 0 || l2 == 0 {
		return 0
	}

	if s1[l1-1] == s2[l2-1] {
		return recLCS(s1[:l1-1], s2[:l2-1]) + 1
	} else {
		return max(recLCS(s1, s2[:l2-1]), recLCS(s1[:l1-1], s2))
	}
}

// 递归中存在很多重复计算，使用备忘录进行优化
func lcsWithDP(s1, s2 string) int {
	l1, l2 := len(s1), len(s2)

	// 使用l1+1和l2+1的原因是，第一行和第一列是base case，我们需要它的值为0
	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if s1[i-1] == s2[j-1] { // 减1以防越界
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1]) // 取二者中大者即可
			}
		}
	}

	return dp[l1][l2]
}

func longestCommonSubsequence(text1 string, text2 string) int {
	// Time limit exceeded
	// return recLCS(text1,text2)
	return lcsWithDP(text1, text2)
}
