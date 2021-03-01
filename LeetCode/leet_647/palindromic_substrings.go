package leet_647

/*
Given a string, your task is to count how many palindromic substrings in this string.

The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.

Example 1:

Input: "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".


Example 2:

Input: "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
*/

//--------------------------------------------
//
func countSubstrings(s string) int {
	var count int
	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if isPalindrome(s, j, i) {
				count++
			}
		}
	}

	return count
}

func isPalindrome(s string, start int, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
	}

	return true
}

//---------------------------------------------
func countSubStrings2(s string) int {
	var count int
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if s[i] == s[j] && (j-i <= 2 || dp[j+1][i-1]) {
				count++
			}
		}
	}

	return count
}
