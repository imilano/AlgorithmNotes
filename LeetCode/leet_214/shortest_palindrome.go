package leet_214

/*
Given a string s, you can convert it to a palindrome by adding characters in front of it. Find and return the shortest palindrome you can find by performing this transformation.



Example 1:

Input: s = "aacecaaa"
Output: "aaacecaaa"
Example 2:

Input: s = "abcd"
Output: "dcbabcd"


Constraints:

0 <= s.length <= 5 * 104
s consists of lowercase English letters only.
*/

func reverse(s string) string {
	r := []rune(s)
	n := len(r)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

//-----------------------------------------
// 直接和翻转字符串做比较
func shortestPalindrome2(s string) string {
	r := reverse(s)
	n := len(s)

	var i int
	for i = n; i >= 0; i-- {
		if s[:i] == r[n-i:] {
			break
		}
	}

	return r[:n-i] + s

}

//----------------------------------------------
// TODO use KMP here
func shortestPalindrome(s string) string {
	return ""
}
