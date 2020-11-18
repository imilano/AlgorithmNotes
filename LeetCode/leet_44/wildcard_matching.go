package leet_44

/*
Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:

    '?' Matches any single character.
    '*' Matches any sequence of characters (including the empty sequence).

The matching should cover the entire input string (not partial).
 */

// -----------------------------------------------------------------------------------------------------------
// wrong answer, we need to flag *
// Original wrong recursive method
func helper(s string, p string, sp int, pp int ) bool {
	if sp == len(s) && pp == len(p) || pp == len(p)-1 && p[pp] == '*' {
		return true
	}

	if (pp < len(p) && p[pp] == '?') || (sp < len(s) && pp < len(p) && s[sp] == p[pp]) {
		return helper(s,p,sp+1,pp+1)
	}

	if pp < len(p) && p[pp] == '*' {
		return helper(s,p,sp,pp+1) || helper(s,p,sp+1,pp+1) || helper(s,p,sp+1,pp)
	}

	return false
}

func isMatchOriginal(s string, p string) bool {
	return helper(s,p,0,0)
}

//------------------------------------------------------------------------
// Greedy Algorithm
func isMatchGreedy(s string, p string) bool {
	sp,pp,match := 0,0,0
	starIdx := -1

	for sp < len(s){
		// ? matches at least one character, we need to move forward both pointer
		if  pp < len(p) && (p[pp] == '?' || p[pp] == s[sp]) {
			sp++
			pp++
			continue
		}

		// //如果不匹配但pp指向'*'，那么记录此时pp的位置以构建回路，同时记录sp的位置, 以标记sp此时可以后移却停留在此一次，同时pp后移
		if pp < len(p) && p[pp] == '*' {
			starIdx = pp
			match = sp
			pp++
		} else if starIdx != -1 {
		//仍然不匹配，但是sp有路可走，且sp已经停在那一次了，那么sp要后移，连同pp停留的位置也要更新，
		//pp直接到回路'*'的后一个位置。此时pp也可以取starIdx，但运行速度会变慢
			pp = starIdx + 1
			match++
			sp = match
		} else {	//仍然不匹配，sp与pp均已无路可走，返回false
			return false
		}
	}

	for pp  < len(p) && p[pp] == '*' {  // //sp扫描完成后要看pp能不能够到达终点，即pp可以沿着'*'行程的通路一直向下
		pp++
	}

	return pp == len(p)  //sp与ss同时到达终点完成匹配
}

// ------------------------------------------------------------------------------------
// Dynamic Programming
// Time complexity: O(m*n)
// Space complexity: O(m*n)
// 令dp[i][j]表示字符串s的前i个字符和模式串的前j的个字符是否匹配。则存在下面三种状态转移方式，其中情况一和情况二可以合并：
//		- 情况一：s[i] == p[j]
//			此时dp[i][j] = dp[i-1][j-1]
//		- 情况二：p[i] == '?'
//			 此时对s没有任何的要求，dp[i][j] = dp[i-1][j-1]
//		 - 情况三： p[i] == '*'
//			 此时存在两种情况：
//			 	- 使用*， 那么dp[i][j] = dp[i-1][j]
//			 	- 不使用*，那么dp[i][j] = dp[i][j-1]
//
//接下来可以看一下初试状态。由于 dp[i][j] 对应着s的前i个字符和模式p的前j个字符，因此所有的 dp[0][j] 和 dp[i][0] 都是边界条件：
//		- dp[0][0] = ture
//			s和p均为空时，匹配成功
//		- dp[i][0] = false
//			p 为空，无法匹配任何字符串
//		- dp[0][j] 需要分情况讨论
//			- dp[0][j] == true
//				当前仅当p的前j-1个字符都是*时
//			- dp[0][j] = false
//				上述不成立，则为假
//golang中，bool数组默认初始化为false。

func isMatchDP(s string, p string) bool {
	m,n := len(s),len(p)
	dp := make([][]bool,m+1)
	for i := range dp {
		dp[i] = make([]bool,n+1)
	}

	// default all the element to false
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] =  true
		} else {
			break
		}
	}


	for i := 1; i <= m; i++ {
		for j := 1; j<= n;j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}

			if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}
