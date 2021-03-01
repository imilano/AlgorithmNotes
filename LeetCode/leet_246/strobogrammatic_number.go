package leet_246

/*
A strobogrammatic number is a number that looks the same when rotated 180 degrees (looked at upside down).

Write a function to determine if a number is strobogrammatic. The number is represented as a string.

Example 1:

Input:  "69"
Output: true
Example 2:

Input:  "88"
Output: true
Example 3:

Input:  "962"
Output: false
*/

func isStrobogrammatic(num string) bool {
	n := len(num)
	left, right := 0, n-1
	for left < right {
		l, r := num[left], num[right]
		switch l {
		case '2', '3', '4', '5', '7': // 只要出现这几个数字，无论怎么翻转都不会相等。
			return false
		case '0', '1', '8': // 这几个数字，如果出现，必须要相等
			if r != l {
				return false
			}
		case '6', '9': // 这两个数字特殊
			if r != '6' && r != '9' {
				return false
			}
		}
		left++
		right--
	}

	return true
}
