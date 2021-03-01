package leet_290

/*
Given a pattern and a string s, find if s follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.



Example 1:

Input: pattern = "abba", s = "dog cat cat dog"
Output: true
Example 2:

Input: pattern = "abba", s = "dog cat cat fish"
Output: false
Example 3:

Input: pattern = "aaaa", s = "dog cat cat dog"
Output: false
Example 4:

Input: pattern = "abba", s = "dog dog dog dog"
Output: false


Constraints:

1 <= pattern.length <= 300
pattern contains only lower-case English letters.
1 <= s.length <= 3000
s contains only lower-case English letters and spaces ' '.
s does not contain any leading or trailing spaces.
All the words in s are separated by a single space.
*/

func wordPattern(pattern string, s string) bool {
	words := getWords(s)
	m1,m2 := make(map[uint8]string),make(map[string]uint8)

	if len(words) != len(pattern) {
		return false
	}
	for i := range pattern {
		v1,ok1 := m1[pattern[i]];
		v2,ok2 := m2[words[i]]

		if ok1 && !ok2 || !ok1 && ok2 {
			return  false
		}
		if ok1 && ok2 {
			if v1 != words[i] && v2 != pattern[i] {
				return false
			}
		} else {
			m1[pattern[i]] = words[i]
			m2[words[i]] = pattern[i]
		}
	}

	return true
}

func getWords(s string) []string {
	var res []string
	var start,end int
	s += " "
	n := len(s)

	for end < n {
		for s[end] != ' ' {
			end++
		}

		res = append(res,s[start:end])
		end++
		start = end
	}

	return res
}