package leet_245

import "math"

/*
Given a list of words and two words word1 and word2, return the shortest distance between these two words in the list.

word1 and word2 may be the same and they represent two individual words in the list.

Example:
Assume that words = ["practice", "makes", "perfect", "coding", "makes"].

Input: word1 = “makes”, word2 = “coding”
Output: 1
Input: word1 = "makes", word2 = "makes"
Output: 3
Note:
You may assume word1 and word2 are both in the list.
*/

// 与之前不同的是，本次的题目中w1和w2可以相同，如果我们继续采用之前的map的方法，那么当w1和w2相同的时候，就会得到最短距离0，显然错误。
// 这里我们分情况讨论：还是采用map的思路，如果w1和w2不同，那么继续采用之前的思路是没有问题的；但是如果w1和w2相同的话，就用一个变量t记录前
// 一次这个单词出现的位置，最后进行比较即可.
// 对map的思路进行优化，其实只需要使用三个临时变量即可。
func shortestWordDistance(words []string, w1, w2 string) int {
	res := math.MaxInt64
	n := len(words)
	p1, p2, t := -1, -1, -1
	for i := 0; i < n; i++ {
		t = p1
		if words[i] == w1 {
			p1 = i
		}
		if words[i] == w2 {
			p2 = i
		}

		if p1 != -1 && p2 != -1 {
			if t != -1 && t != p1 && w1 == w2 {
				res = min(res, abs(p1-t))
			} else {
				res = min(res, abs(p1-p2))
			}
		}
	}

	return res
}

//----------------------------
// 优化思路，不需要用三个变量，只需要使用两个变量即可。将p1初始化为数组长度，将p2初始化为数组长度的相反数，然后当p1和p2相等时，用p1记录p2的历史值
func shortestWordDistance2(words []string, w1, w2 string) int {
	n := len(words)
	p1, p2 := n, -n
	res := math.MaxInt64

	for i, w := range words {
		if w == w1 && w1 == w2 {
			p1 = p2
		} else {
			p1 = i
		}

		if w == w2 {
			p2 = i
		}

		res = min(res, abs(p1-p2))
	}

	return res
}

//util func
func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
