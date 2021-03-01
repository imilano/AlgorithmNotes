package leet_89

import "math"

/*
The gray code is a binary numeral system where two successive values differ in only one bit.

Given a non-negative integer n representing the total number of bits in the code, print the sequence of gray code. A gray code sequence must begin with 0.
*/

//也就是将十进制转换为格雷码，参见维基百科。

func grayCode(n int) []int {
	var res []int
	for i := 0; i < int(math.Pow(2, float64(n))); i++ {
		res = append(res, (i>>1)^i)
	}

	return res
}
