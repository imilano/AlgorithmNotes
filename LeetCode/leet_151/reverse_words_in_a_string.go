package leet_151

/*
Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.



Example 1:

Input: s = "the sky is blue"
Output: "blue is sky the"
*/

func reverseWords(s string) string {
	var res string
	if len(s) == 0 {
		return res
	}

	var words []string
	for len(s) != 0 {
		if s[0] == ' ' {
			s = s[1:]
			continue
		} else {
			i := 0
			for i < len(s) && s[i] != ' ' {
				i++
			}
			words = append(words, s[0:i])
			s = s[i:]
		}
	}

	length := len(words)
	for length != 0 {
		res += words[length-1] + " "
		length--
	}
	if len(res) > 1 {
		res = res[:len(res)-1]
	}
	return res
}
