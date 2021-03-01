package leet_115

/*
 Given two strings s and t, return the number of distinct subsequences of s which equals t.
A string's subsequence is a new string formed from the original string by deleting some (can be none) of the characters
without disturbing the relative positions of the remaining characters. (i.e., "ACE" is a subsequence of "ABCDE" while "AEC" is not).

It's guaranteed the answer fits on a 32-bit signed integer.
*/

//状态定义：定义dp[i][j]表示s中的前j个字符有多少种表示t中前i个字符的方法
//基准情况：
//		- 当t为空字符串时，也就是i为0时间，空字符串总是s的子串，所以dp[0][j] = 1
//		- 当s为空字符时，除了dp[0][0]=1之外（空字符串是空字符串的子串）， dp[i][0]=0 （空字符串不会有子串）
//转态转移：
//		- 当s中的第j个字符（下标j-1）和t中的第i个字符（下标i-1）相等时，根据是否使用s中的第j个字符，有两种可能：
//			- 使用s[j-1]（也就是使用s的第j个字符），那么dp[i][j] = dp[i-1][j-1]（也就是说，如果我要使用s中的第j个字符来匹配t中的第i个字符，
//			  那么我当前的匹配数量，应该等于使用s中的j-1个字符来匹配t中的i-1个字符的数量。）
//			- 不使用s[j-1](也就是不使用s的第j个字符)，意味着我只使用s中的前j-1个字符来匹配t中的前i个字符，故而dp[i][j] = dp[i][j-1]（注意我的数组长度是n+1，不是n）
//			- 综上所述，我当前的状态应该是上面两种状态的总和，故而: dp[i][j] = do[i-1][j-1] + dp[i][j-1]
//		- 当s中的第j个字符不等于t中的第i个字符时，dp[i][j] = dp[i][j-1]， 也就是说，当前的状态等于使用s中前j-1个字符去匹配t中前i个字符的数量。

func numDistinct(s string, t string) int {
	lens, lent := len(s), len(t)
	dp := make([][]int, lent+1)
	for i := range dp {
		dp[i] = make([]int, lens+1)
	}

	// base case
	for i := 0; i <= lens; i++ {
		dp[0][i] = 1
	}

	for i := 1; i <= lent; i++ {
		for j := 1; j <= lens; j++ {
			if s[j-1] == t[i-1] {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[lent][lens]
}
