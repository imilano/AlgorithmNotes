package leet_91

/*
A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26

Given a non-empty string containing only digits, determine the total number of ways to decode it.

The answer is guaranteed to fit in a 32-bit integer.
 */

func numDecodings(s string) int {
	length := len(s)
	if length == 2 {
		return 2
	}

	if length == 1 {
		return 1
	}


	str := string(s[0]) + string(s[1])
	if str >= "10" && str <= "26" {
		return numDecodings(s[2:]) * 2 + numDecodings(s[1:])
	} else {
		return numDecodings(s[1:])
	}
}
