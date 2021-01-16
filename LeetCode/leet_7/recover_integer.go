package leet_7

import "math"

/*
	Given a 32-bit signed integer, reverse digits of an integer.
	Note:
	Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
	For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

func reverseConcise(x int) int {
	var res int
	var neg bool
	if x < 0 {
		neg = true
		x = -x
	}

	for x != 0 {
		res *= 10
		res = res + x%10
		x /= 10
	}

	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	if neg {
		res = -res
	}

	return res
}

// original
func get_digits(x int) []int {
	var res []int
	num := x
	if x < 0 {
		num = -num
	}

	for num != 0 {
		n := num % 10
		num /= 10

		res = append([]int{n}, res...)
	}

	return res
}

func reverse(x int) int {
	var res int
	var negative bool
	if x >= -9 && x <= 9 {
		return x
	}

	if x < 0 {
		negative = true
	}

	digits := get_digits(x)
	length := len(digits)

	var overflow bool
	var old int

	for length != 0 && !overflow {
		old = res
		res = res*10 + digits[length-1]
		if (old > 0 && res < 0) || (old < 0 && res > 0) || res > math.MaxInt32 || res < math.MinInt32 {
			overflow = true
		}
		length--
	}

	if overflow {
		return 0
	}

	if negative {
		return -res
	}

	return res
}
