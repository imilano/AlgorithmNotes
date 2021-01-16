package leet_8

import "math"

/*
	Implement atoi which converts a string to an integer.
	The function first discards as many whitespace characters as necessary until the first non-whitespace character is found.
	Then, starting from this character takes an optional initial plus or minus sign followed by as many numerical digits as possible,
	and interprets them as a numerical value.
	The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the
	behavior of this function.
	If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str
	is empty or it contains only whitespace characters, no conversion is performed.
	If no valid conversion could be performed, a zero value is returned.

	Note:
    	Only the space character ' ' is considered a whitespace character.
    	Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
		If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.

*/

func MyAtoI(s string) int {
	return myAtoi(s)
}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	var space, res int
	var negative, overflow bool
	length := len(s)
	for space < length && s[space] == ' ' {
		space++
	}

	if space < length && s[space] == '-' {
		negative = true
		space++
	}

	if !negative && space < length && s[space] == '+' {
		space++
	}

	if space >= length {
		return 0
	}

	if s[space] >= 'a' && s[space] <= 'z' || s[space] == '+' || s[space] == '-' {
		return 0
	}

	for space < length && s[space] >= '0' && s[space] <= '9' {
		res = res*10 + int(s[space]-'0')
		space++

		if res > math.MaxInt32 || res < math.MinInt32 {
			overflow = true
			break
		}
	}

	if overflow && negative {
		return math.MinInt32
	} else if overflow && !negative {
		return math.MaxInt32
	}

	if negative {
		return -res
	}

	return res
}
