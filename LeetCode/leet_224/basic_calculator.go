package leet_224

import (
	"lightsinger.life/algorithm/leetcode/utils"
)

/*
Given a string s representing an expression, implement a basic calculator to evaluate it.



Example 1:

Input: s = "1 + 1"
Output: 2
Example 2:

Input: s = " 2-1 + 2 "
Output: 3
Example 3:

Input: s = "(1+(4+5+2)-3)+(6+8)"
Output: 23


Constraints:

1 <= s.length <= 3 * 105
s consists of digits, '+', '-', '(', ')', and ' '.
s represents a valid expression.
*/

//--------------------------------------------------------------------------
// 逐次扫描
// 使用栈进行辅助计算。用变量sign表示当前符号，如果遇到加号，则sign赋为1；如果遇到减号，sign赋为-1；如果遇到左括号，将当前结果res和符号sign入栈，
//然后重置res为0，sign为1；如果遇到了右括号，结果res乘以栈顶的符号，栈顶符号出栈，结果res加上栈顶的数字，栈顶元素出栈
func calculate(s string) int {
	res, sign, n := 0, 1, len(s)
	stack := utils.NewStack()

	for i := 0; i < n; i++ {
		c := s[i]
		if c >= '0' && c <= '9' { // 读取一个完整的数字
			var num int
			for i < n && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			res += sign * num
			i-- // 回退多读的那一个字符
		}

		if c == '+' {
			sign = 1
		}

		if c == '-' {
			sign = -1
		}
		if c == '(' {
			stack.Push(res)
			stack.Push(sign)
			res = 0
			sign = 1
		}
		if c == ')' {
			preSign := stack.Pop().(int)
			preNum := stack.Pop().(int)
			res *= preSign
			res += preNum
		}
	}

	return res
}

//------------------------------------------------------------------
// 递归方式。对括号的计算可以进行递归处理
func calculate2(s string) int {
	res, sign, num, n := 0, 1, 0, len(s)
	for i := 0; i < n; i++ {
		c := s[i]
		if c >= '0' && s[i] <= '9' {
			num = num*10 + int(c-'0')
		}
		if c == '(' { // 对括号内的进行递归计算
			j := i
			cnt := 0
			for ; i < n; i++ { // 找到相匹配的右括号
				if s[i] == '(' {
					cnt++
				}
				if s[i] == ')' {
					cnt--
				}
				if cnt == 0 {
					break
				}
			}

			num = calculate2(s[j+1 : i])
		}

		if c == '+' || c == '-' || i == n-1 {
			res += sign * num
			num = 0
			if c == '+' {
				sign = 1
			} else if c == '-' {
				sign = -1
			}
		}
	}

	return res
}
