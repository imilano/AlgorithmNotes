package leet_140

/*
Share
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, add spaces in s to construct a sentence where each word is a valid dictionary word. Return all such possible sentences.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.
Example 1:

Input:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
Output:
[
  "cats and dog",
  "cat sand dog"
]
*/

// TLE, WRONG answer
func wordBreak(s string, wordDict []string) []string {
	var res []string
	m := make(map[string]int)
	for _, v := range wordDict {
		m[v] = 1
	}

	helper(s, "", &m, 0, &res)
	return res
}

func helper(s string, cur string, m *map[string]int, start int, res *[]string) {
	if start >= len(s) {
		*res = append(*res, cur[1:])
		return
	}

	for i := start; i < len(s); i++ {
		if _, ok := (*m)[s[start:i+1]]; !ok {
			continue
		}
		cur += " " + s[start:i+1]
		helper(s, cur, m, i+1, res)
		cur = cur[:len(cur)-(i-start+2)]
	}
}

//----------------------------------------------
// 使用memo数组进行优化
func wordBreak2(s string, wordDict []string) []string {
	m := make(map[string][]string)
	return helper2(s, wordDict, &m)
}

func helper2(s string, dict []string, m *map[string][]string) []string {
	if _, ok := (*m)[s]; ok {
		return (*m)[s]
	}

	if s == "" {
		return []string{""}
	}

	var res []string
	for _, d := range dict {
		if s[0:len(d)] != d {
			continue
		}

		rem := helper2(s[len(d):], dict, m)
		for _, str := range rem {
			var t string
			if str == "" {
				t = d + "" + str
			} else {
				t = d + " " + str
			}

			res = append(res, t)
		}
	}

	(*m)[s] = res
	return (*m)[s]
}
