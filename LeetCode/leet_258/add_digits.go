package leet_258

/*
Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.

Example:

Input: 38
Output: 2
Explanation: The process is like: 3 + 8 = 11, 1 + 1 = 2.
             Since 2 has only one digit, return it.
*/

func addDigits(num int) int {
	for num/10 != 0 {
		var t int
		for num != 0 {
			t += num % 10
			num /= 10
		}

		num = t
	}

	return num
}

//从1到20进行一一列举，可以看出数字的出现都是有规律可循的，每九个一循环，对9取余即可，但是对于9的倍数不适用，故而采取(n-1)%9+1的形式。但是
//这样又引入了n=0的特例，额外做一个判断即可
func addDigits2(num int) int {
	if num == 0 {
		return 0
	} else {
		return (num-1)%9 + 1
	}
}
