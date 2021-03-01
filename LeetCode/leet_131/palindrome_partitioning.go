package leet_131

/*
Given a string s, partition s such that every substring of the partition is a palindrome.

Return all possible palindrome partitioning of s.

Example:

Input: "aab"
Output:
[
  ["aa","b"],
  ["a","a","b"]
]
*/

// same question: 139,647

//--------------------------------------------------
// #backtracing
func partition(s string) [][]string {
	var res [][]string
	var cur []string
	helper(s, 0, &cur, &res)
	return res
}

func helper(s string, start int, cur *[]string, res *[][]string) {
	if start >= len(s) {
		tmp := make([]string, len(*cur))
		copy(tmp, *cur)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(s); i++ {
		if !isPalindrome(s, start, i) {
			continue
		}

		*cur = append(*cur, s[start:i+1])
		helper(s, i+1, cur, res)
		*cur = (*cur)[:len(*cur)-1]
	}
}

func isPalindrome(s string, start int, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}

//----------------------------------------------------------------
// 上面一个算法的优化，通过事先计算s的回文串，从而减少函数调用的开销
func partition2(s string) [][]string {
	var res [][]string
	var cur []string
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	// calculate palindrome of s
	// #trick, need to learn
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if s[j] == s[i] && (i-j <= 2 || dp[j+1][i-1]) {
				dp[j][i] = true
			}
		}
	}

	helper2(s, dp, 0, &res, &cur)
	return res
}

func helper2(s string, dp [][]bool, start int, res *[][]string, cur *[]string) {
	if start >= len(s) {
		tmp := make([]string, len(*cur))
		copy(tmp, *cur)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(s); i++ {
		if !dp[start][i] {
			continue
		}

		*cur = append(*cur, s[start:i+1])
		helper2(s, dp, i+1, res, cur)
		*cur = (*cur)[:len(*cur)-1]
	}
}
