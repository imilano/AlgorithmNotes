package leet_58

/*
Given a string s consists of upper/lower-case alphabets and empty space characters ' ',
return the length of last word (last word means the last appearing word if we loop from left to right) in the string.
If the last word does not exist, return 0.
Note: A word is defined as a maximal substring consisting of non-space characters only.
Example:
Input: "Hello World"
Output: 5
*/

func lengthOfLastWord(s string) int {
	var res int
	length := len(s)
	if length == 0 {
		return res
	}

	i := length - 1
	for i > 0 && s[i] == ' ' {
		i--
	}

	for i >= 0 && s[i] != ' ' {
		res++
		i--
	}

	return res
}
