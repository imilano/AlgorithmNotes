package leet_243

import "math"

/*
Given a list of words and two words word1 and word2, return the shortest distance between these two words in the list.

Example:
Assume that words = ["practice", "makes", "perfect", "coding", "makes"].

Input: word1 = “coding”, word2 = “practice”
Output: 3
Input: word1 = "makes", word2 = "coding"
Output: 1
Note:
You may assume that word1 does not equal to word2, and word1 and word2 are both in the list.
*/

// 最简单你的想法，扫描数组，分别得到两个字符串的下标，然后对下标进行计算即可。
func shortestDistance(words []string, word1, word2 string) int {
	res := math.MaxInt64
	if len(words) < 2 {
		return 0
	}

	var c1, c2 []int
	for i := range words {
		if words[i] == word1 {
			c1 = append(c1, i)
		}
		if words[i] == word2 {
			c2 = append(c2, i)
		}
	}

	for i := 0; i < len(c1); i++ {
		for j := 0; j < len(c2); j++ {
			res = min(res, abs(c1[i]-c2[j]))
		}
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a

}

//------------------------
// 实际上可能只需要一次遍历
func shortestDistance2(words []string, word1, word2 string) int {
	res := math.MaxInt64
	idx1, idx2 := -1, -1

	for i := 0; i < len(word2); i++ {
		if words[i] == word1 {
			idx1 = i
		}
		if words[i] == word2 {
			idx2 = i
		}
		if idx1 != -1 && idx2 != -1 {
			res = min(res, abs(idx1-idx2))
		}
	}

	return res
}
