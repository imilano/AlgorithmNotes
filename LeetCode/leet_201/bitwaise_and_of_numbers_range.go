package leet_201

import "math"

/*
Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.

Example 1:

Input: [5,7]
Output: 4
Example 2:

Input: [0,1]
Output: 0

*/


func rangeBitwiseAnd(m int, n int) int {
	res := m
	for i := m+1; i<=n;i++{
		res &= i
	}

	return res
}


// 通过观察可以得到， 最后的结果是该数字范围内所有数的左边的共同部分，从而可以通过移位操作来实现
func rangeBitwiseAndConcise(m int,n int) int {
	d := math.MaxInt32
	for m & d != n & d {
		d <<= 1
	}

	return d & m
}