package leet_172

/*
Given an integer n, return the number of trailing zeroes in n!.

Follow up: Could you write a solution that works in logarithmic time complexity?



Example 1:

Input: n = 3
Output: 0
Explanation: 3! = 6, no trailing zero.
*/

// 实际只需要计算一共有几个5，其个数就是0的个数

func trailingZeroesConcise(n int)int {
	var res int
	for i := 1; i <= n;i++ {
		num := i
		if num % 5 != 0 {
			continue
		}

		for num % 5 == 0 {
			res += 1
			num /= 5
		}
	}

	return res
}

func trailingZeroes(n int) int {
	var res int
	if n < 5 {
		return res
	}
	for i := 5; i <= n;i++ {
		res += countFive(i)
	}

	return res
}

func countFive(num int) int {
	if num % 5 != 0 {
		return 0
	}

	var count int
	for num % 5 == 0 {
		count+= 1
		num /= 5
	}

	return count
}
