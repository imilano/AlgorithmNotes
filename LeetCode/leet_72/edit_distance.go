package leet_72

import "math"

/*
	Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.
	You have the following 3 operations permitted on a word:
    	Insert a character
		Delete a character
    	Replace a character
*/

// 递归解法
func getMin(i, j, k int) int {
	return int(math.Min(float64(i), math.Min(float64(j), float64(k))))
}
func dp(i, j int, s1, s2 string) int {
	if i == -1 {
		return j + 1
	}

	if j == -1 {
		return i + 1
	}

	if s1[i] == s2[j] {
		return dp(i-1, j-i, s1, s2)
	} else {
		return getMin(
			dp(i, j-1, s1, s2)+1,
			dp(i-1, j, s1, s2)+1,
			dp(i-1, j-1, s1, s2))
	}
}

func minDistance(word1 string, word2 string) int {
	// 初始化为下标，故而初始两个参数为两个字符串的最后一个下标。
	return dp(len(word1)-1, len(word2)-1, word1, word2)
}

func MinDistance(word1, word2 string) int {
	return minDistance(word1, word2)
}

// 递归很容易就重复计算了子问题，此处采用DP数组

func minDistanceDP(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	dp := make([][]int, len1+1)
	for i := 0; i < len1; i++ {
		dp[i] = make([]int, len2+1)
	}

	// base condition
	for i := 0; i <= len1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] { // warning
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = getMin(
					dp[i-1][j]+1,
					dp[i][j-1]+1,
					dp[i-1][j-1]+1,
				)
			}
		}
	}

	return dp[len1][len2]
}
