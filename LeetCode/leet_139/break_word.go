package leet_139

/*
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.

Example 3:

Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false
*/

// 首先，为了加快查询速度，需要经dict转换为map。其次，稍稍分析可以看到，此题适用动态规划，定义dp[i][j]s中i到j之间字符构成的字符串是否出现在dict中，其中j大于等于i
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		if len(wordDict) == 0 {
			return true
		}

		return false
	}

	m := make(map[string]int)
	for i := range wordDict {
		m[wordDict[i]] = 1
	}

	// TODO
	// 遍历从而完成dp数组的初始化，然后查找dp中是否存在一条从第一行到达最后一行的全是true的路径
	// 如果存在，则为true，否则为false
	return false
}

//---------------------------------------------------------
// Recursive + memo
// 定会memo[i]表示从i到n是否出现在dict中
func wordBreak1(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for i := range wordDict {
		m[wordDict[i]] = true
	}

	// use memo to reduce time complexity
	memo := make([]int, len(s)) // 0表示默认初始化值，1表示出现在dict中，2表示未出现在dict中
	return check(s, m, memo, 0)
}

func check(s string, wordDict map[string]bool, memo []int, start int) bool {
	if start >= len(s) {
		return true
	}

	if memo[start] != 0 {
		return memo[start] == 1
	}

	for i := start + 1; i < len(s); i++ {
		if wordDict[s[i:]] == true && check(s, wordDict, memo, i) {
			memo[start] = 1
		}
	}

	memo[start] = 2
	return false
}

// -------------------------------------------------
// dynamic programming
// 使用dp[i] 表示从s从0到i是否可以拆分。对于0到i之间的字符串，我们使用另一个循环0到j，如果0到j在dict中并且j到i也在dict中，则dp[i] =true
func wordBreak2(s string, wordDcit []string) bool {
	m := make(map[string]bool)
	for i := range wordDcit {
		m[wordDcit[i]] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	// s[0...i]是否可拆分
	for i := 0; i <= len(s); i++ {
		// s[0...i]是否可以拆分为存在于dict中的[0...j] [j+1...i]
		for j := 0; j < i; j++ {
			if dp[j] && m[s[j:i]] == true {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}

//-----------------------------
// bfs
func wordBreak3(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for i := range wordDict {
		m[wordDict[i]] = true
	}

	visited := make([]bool, len(s))
	var que []int
	que = append(que, 0)

	for len(que) != 0 {
		start := que[0]
		que = que[1:]

		if !visited[start] {
			for i := start + 1; i <= len(s); i++ {
				if m[s[start:i]] == true {
					que = append(que, i)
					if i == len(s) {
						return true
					}
				}
			}
			visited[start] = true
		}
	}

	return false
}
