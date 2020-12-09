package leet_202

/*
Write an algorithm to determine if a number n is "happy".

A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

Return True if n is a happy number, and False if not.

Example:

Input: 19
Output: true
Explanation:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
*/

func isHappy(n int) bool {
	if n <= 0 {
		return false
	}

	var t int
	for t <= 100 {
		digits := getDigits(n)

		n = 0
		for i := range digits {
			n += digits[i] * digits[i]
		}

		if n == 1 {
			return true
		}
		t++
	}

	return false
}

func getDigits(n int) []int {
	var res []int
	for n != 0 {
		res = append(res,n % 10)
		n /= 10
	}

	return res
}


// 非快乐数在循环的最后总会出现4(以11为例)，并且会一直循环出现。根据这个特点，可以简化代码
func isHappy2(n int) bool {
	if n <= 0 {
		return  false
	}

	for n != 1 && n != 4 {
		var sum int
		for n != 0 {
			sum += (n%10)*(n%10)
			n/= 10
		}

		n = sum
	}

	return n == 1
}


// 因为非快乐数最后都会循环出现，故而也可以利用map来记录数字出现的次数，如果出现的次数大于1，并且该数字不是1，那么就说明不是快乐数，否则就是快乐数
func isHappy3(n int)bool {
	m := make(map[int]int)
	for n != 1 {
		var sum int
		for n != 0 {
			sum += (n % 10) * ( n %10)
			n /= 10
		}

		// 如果出现过
		if m[n] == 1 {
			break
		}
		n = sum
		m[n] = 1
	}

	return n == 1
}
