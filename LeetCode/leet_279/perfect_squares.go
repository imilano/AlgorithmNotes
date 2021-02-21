package leet_279

import "math"

/*
Given an integer n, return the least number of perfect square numbers that sum to n.

A perfect square is an integer that is the square of an integer; in other words, it is the product of some integer with itself. For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.



Example 1:

Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.
Example 2:

Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.


Constraints:

1 <= n <= 104
*/

//四平方和定理：任意一个正整数都可以表示成四个以内整数的平方和。也就是说，答案只能是1、2、3、4中的一个。而如果一个数含有因子4，那么可以该数的
//平方数等于该数除以4的平方数；如果一个数除以8余7的话，那么一定是由四个平方数组成。上述两个性质可以用于化简，化简完成之后，可以将一个数拆为两个
//平方数之和。
func numSquares(n int) int {
	for n % 4 == 0 {
		n /= 4
	}

	if n %8 == 7 {
		return 4
	}
	for a := 0; a*a <= n;a++{
		b := int(math.Sqrt(float64(n-a*a)))
		if a*a + b*b == n {
			if a > 0 && b > 0 { // 如果两个数都是正整数，那么说明拆出了两个正整数，返回2
				return 2
			}else  {  // 上述不成立的话，说明只要一个数是正整数，那么返回1
				return 1
			}
		}
	}

	return 3  // 既不是1也不是2和4，那么只能是3
}


func numSquares2(n int) int {
	dp := []int{0}
	for len(dp) <= n {
		m := len(dp)
		val := math.MaxInt32
		for i := 1; i*i <= m; i++ {
			val = min(val,dp[m-i*i]+1)  // 举例多多领会
		}

		dp = append(dp,val)
	}

	return dp[len(dp)-1]
}

func min(a,b int) int {
	if a < b {
		return a
	}

	return b
}