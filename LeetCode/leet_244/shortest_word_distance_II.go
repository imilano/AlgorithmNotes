package leet_244

import "math"

/*
Design a class which receives a list of words in the constructor, and implements a method that takes two words word1 and word2 and return the shortest distance between these two words in the list. Your method will be called repeatedly many times with different parameters.

Example:
Assume that words = ["practice", "makes", "perfect", "coding", "makes"].

Input: word1 = “coding”, word2 = “practice”
Output: 3
Input: word1 = "makes", word2 = "coding"
Output: 1
Note:
You may assume that word1 does not equal to word2, and word1 and word2 are both in the list.
*/

//使用map
type wordDistance struct {
	m map[string][]int
}

func (w *wordDistance) New(words []string) {
	for i, v := range words {
		w.m[v] = append(w.m[v], i)
	}
}

func (w *wordDistance) shortest(w1, w2 string) int {
	res := math.MaxInt64
	for i := 0; i < len(w.m[w1]); i++ {
		for j := 0; j < len(w.m[w2]); j++ {
			res = min(res, abs(w.m[w1][i]-w.m[w2][j]))
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
