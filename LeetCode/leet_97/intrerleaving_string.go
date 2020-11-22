package leet_97

/*
Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.

An interleaving of two strings s and t is a configuration where they are divided into non-empty substrings such that:

s = s1 + s2 + ... + sn
t = t1 + t2 + ... + tm
|n - m| <= 1
The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...
Note: a + b is the concatenation of strings a and b.
 */

// Dynamic Programming
// dp[i][j] 表示s3的前i+j个字符能否由s1的前i个字符（对应下标i-1)和s2的前j个字符(对应下标j-1)组成。
// base case:
//		- 由于s3由s1和s2组合而成，所以如果s3的长度不等于s1的长度与s2的长度之和，那么返回false
//		- 当s1为空，s2为空时，s3也为空，故赋值dp[0][0] 为true
//		- 当i为0，那么dp[0][j]取决于dp[0][j-1]，只有j-1为true的情况下，do[0][j]才为true；j为空的情况同理
// 转态转移：
//		当前dp[i][j]可以从上方dp[i-1][j]转移而来，也可以从左边dp[i][j-1]转移而来，也就说，只要二者有一个为true，那么当前dp[i][j]就是true。
//		从上边转移：
//			dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i-1+j]
//		从左边转移：
//			dp[i][j] = dp[i][j-1] && s2[j-1] == s3[j-1+i]
//
//		综合，状态转移方程为: dp[i][j] =  dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i-1+j] || dp[i][j] = dp[i][j-1] && s2[j-1] == s3[j-1+i]


// Time complexity: O(m*n)
// Space complexity: O(m*n)
func isInterleave(s1 string, s2 string, s3 string) bool {
	len1,len2,len3 := len(s1),len(s2),len(s3)
	if len1 + len2 != len3 {
		return false
	}

	dp := make([][]bool, len1+1)
	for i  := range dp {
		dp[i] = make([]bool, len2+1)
	}

	// base case
	dp[0][0] = true
	for i:= 1; i<=len1;i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}

	for i := 1; i<=len2;i++ {
		dp[0][i] = dp[0][i-1] && s2[i-1] == s3[i-1]
	}

	// dynamic programming
	for i:= 1;i<=len1;i++ {
		for j := 1; j<=len2;j++ {
			dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i-1+j] || dp[i][j-1] && s2[j-1] == s3[i+j-1]
		}
	}

	return dp[len1][len2]
}
