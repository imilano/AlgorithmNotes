package leet_159

/*
Given a string s , find the length of the longest substring t  that contains at most 2 distinct characters.

Example 1:

Input: "eceba"
Output: 3
Explanation: tis "ece" which its length is 3.

Example 2:

Input: "ccaabbb"
Output: 5
Explanation: tis "aabbb" which its length is 5.
 */

//-----------------------------------------------------
// 使用map，记录一个字符出现的次数。每当map中的映射数量大于3时，就需要从map中删除字符。删除哪一个字符呢？我们需要维持一个left 指针，指向s
// 中当前被删除的字符的位置，每次删除完后，将left指针后移。依次遍历整个字符串。
func lengthOfLongestSubstringTwoDistinct(s string) int {
	var res,left int
	m := make(map[uint8]int)
	for i :=0; i< len(s);i++  {
		m[s[i]]++  // 字符计数
		for len(m) > 2 {  // 每当映射数大于2，则需要删除字符
			m[s[left]]-- // 删除当前left执行的字符
			if m[s[left]] == 0 {  // 如果当前字符的出现次数为0，那么可以直接从map中剔除该字符
				delete(m,s[left])
			}
			left++ // 下一个待删除字符
		}

		// 不断更新字符的最大值
		res = max(res, i - left +1)   // core operation
	}

	return res
}

//---------------------------------
// 使用map映射每个字符最新的坐标，思路如同上一个
func lengthOfLongestSubstringTwoDistinctMapCor(s string) int {
	var left,res int
	m := make(map[uint8]int)
	for i := 0; i < len(s) ;i++ {
		m[s[i]] = i
		for len(m) > 2 {
			if m[s[left]] == left {
				delete(m,s[left])
			}

			left++
		}

		res = max(res,i-left+1)
	}

	return res
}







// ----------------------------------
// utils
func max(a,b int) int {
	if a > b {
		return a
	}

	return  b
}
