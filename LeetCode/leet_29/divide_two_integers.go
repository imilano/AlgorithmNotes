package leet_29

import "math"

/*
	Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.
	Return the quotient after dividing dividend by divisor.
	The integer division should truncate toward zero, which means losing its fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) = -2.

Note:
    	Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For this problem, assume that your function returns 231 − 1 when the division result overflows.
*/

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func divide(dividend int, divisor int) int {
	var res int
	if divisor == 0 || dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	var negative bool
	if (dividend < 0) == (divisor < 0) {
		negative = false
	} else {
		negative = true
	}

	dvd, dvs := int64(abs(dividend)), int64(abs(divisor))
	for dvd >= dvs {
		tmp := dvs
		mul := 1
		for tmp<<1 <= dvd {
			tmp <<= 1
			mul <<= 1
		}
		dvd -= tmp
		res += mul
	}

	if negative {
		return -res
	}

	return res
}
