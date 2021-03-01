package leet_05

/*
	Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
*/

/*
	Solutions:
		1. 翻转字符串，然后让该字符串和原字符串进行比较，找出二者的最长公共子串。需要注意的是，回文串一定是二者的最长公共字串，但反之不成立。
			如 S = "abacdfgdcaba","abacd"是二者的最长公共子串，但是并不是回文。因此，需要对最长公共子串进行回文验证。
		2. 双指针法(中心扩散）
			时间复杂度O(n^2)，空间复杂度O(1)(只用到常数个临时变量，与字符串长度n无关)
			回文的关键在于根据特定字符为中心，不断向两边扩展；然而，对于偶数长度的回文串，其中心并不是单独一个字符，而是两个字符。故而在findPalindrome函数中，
			使用了left和right参数，从而使得一个函数可以处理奇偶回文串的情况。
		3.	动态规划
			时间复杂度O(n^2)，空间复杂度O(n^2)。
			动态规划三大特征：重叠子问题、最优子结构、状态转移方程。对于回文串而言，其三个特征均具备。
			问题的最优解包含子问题的最优解，是为最优子结构。显然，对于一个字符串，该字符串为回文串的前提是该字符串的子串也是回文串。
			如果算法反复求解相同的子问题，那么就称该问题具有重叠子问题性质。显然，回文串具有该性质。
			定义二维数组dp[i][j]，表示S[i]到S[j]是否为回文串，j可取到。则有dp[i][j]为最长回文的前提是dp[i+1][j-1]为回文且S[i]=S[j]。

			也即，dp[i][j] = (dp[i+1][j-1] && S[i] == S[j])。对于边界情况：
			    如果子串 s[i + 1..j - 1] 只有 1 个字符，即去掉两头，剩下中间部分只有1个字符（如111），显然是回文；
			    如果子串 s[i + 1..j - 1] 为空串，那么子串 s[i, j] 一定是回文子串。

			因此，在 s[i] == s[j] 成立和 j - i < 3 的前提下，直接可以下结论，dp[i][j] = true，否则才执行状态转移。初始化时，dp[i][i]必然为true。

			Links: https://leetcode-cn.com/problems/longest-palindromic-substring/solution/zhong-xin-kuo-san-dong-tai-gui-hua-by-liweiwei1419/
*/

/*
Tag: Binary Search
Tag: Dynamic Programming
Tag: Palindrome
*/

// Solution1, two pointer
func findPalindrome(s string, left, right int) string {
	l := len(s)
	for left >= 0 && right < l && s[right] == s[left] {
		left--
		right++
	}

	return s[left+1 : right]
}

func longestPalindrome(s string) string {
	var result string
	l := len(s)
	for i := 0; i < l; i++ {
		// 中心扩散法，对于每个元素，都以该元素为中心向两边扩展
		// 当字符个数为偶数时，中心字符有两个字符，故而findPalindrome应该从考虑左右两个字符的位置。
		r := findPalindrome(s, i, i)
		ss := findPalindrome(s, i, i+1)

		if len(r) > len(result) {
			result = r
		}

		if len(ss) > len(result) {
			result = ss
		}
	}

	return result
}

// Solution 2 , brute force
func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

//#TODO 对于长度为1的字符会出错
func bruteForce(s string) string {
	var res string
	l := len(s)
	//if l <= 1  {
	//	return s
	//}

	// Warning!!!!! LENGTH -1 !!!!!
	//for i := 0; i < l-1; i++ {
	//	for j := i + 1; j < l; j++ {
	// 改成下面这样的循环方式，可以处理控制符以及只有单个字符的情况
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] && isPalindrome(s, i, j) {
				if len(res) < j-i+1 {
					res = s[i : j+1]
				}
			}
		}
	}

	return res
}

// Solution 3, dynamic programming
// #TODO 对于长度为2的字符会出错。
func dynamicProgramming(s string) string {
	// 空串或长度为1，必然为回文串
	// 非空串必然会有大小至少为1的回文串
	start, maxLen, n := 0, 1, len(s)

	// TODO 优化空间复杂度
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	// 注意，在状态转移方程中，我们是从长度较短的字符串向长度较长的字符串进行转移的，因此一定要注意动态规划的循环顺序。
	// 动态规划中，注意状态方程的转移方向
	for j := 0; j < n; j++ {
		// base case
		dp[j][j] = true
		for i := 0; i < j; i++ {
			//if s[i] != s[j] {
			//	dp[i][j] = false
			//} else {
			//	if j-i < 3 {
			//		dp[i][j] = true
			//	} else {
			//		dp[i][j] = dp[i+1][j-1]
			//	}
			//}

			// 以上表达式的更加简练的写法
			dp[i][j] = s[i] == s[j] && (j-i < 2 || dp[i+1][j-1])
			if dp[i][j] && j-i+1 > maxLen {
				start = i
				maxLen = j - i + 1
			}
		}
	}
	return s[start : start+maxLen]
}
