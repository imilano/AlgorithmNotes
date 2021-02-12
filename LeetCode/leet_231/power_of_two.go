package leet_231

/*
Given an integer n, return true if it is a power of two. Otherwise, return false.

An integer n is a power of two, if there exists an integer x such that n == 2x.



Example 1:

Input: n = 1
Output: true
Explanation: 20 = 1
*/

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}

	if n%2 == 1 {
		return false
	}

	return isPowerOfTwo(n / 2)
}

// 2的幂，写成二进制都有一个特点，就是最高位都是1，其他位都是0.利用这个特点，只需要计算数字的二进制中1的个数，如果1的个数是1，那么就是幂，否则不是
func isPowerOfTwo2(n int) bool {
	cnt := 0
	for n > 0 {
		cnt += (n & 1)
		n >>= 1
	}

	return cnt == 1
}

// 如上所述，如果一个数的为2的幂，那么最高位为1，其余位都为0；那么，如果将这样的一个数减去1，那么最高位变为0，其余位变为1，将这个数与原数做与运算，如果结果是0，那么就是2的幂，否则不是
func isPowerOfTwo3(n int) bool {
	return n > 0 && ((n-1)&n) == 0 // 减1的方法对负数不适用
}
