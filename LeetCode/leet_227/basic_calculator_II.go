package leet_227

import (
	"lightsinger.life/algorithm/leetcode/utils"
)

/*
Given a string s which represents an expression, evaluate this expression and return its value.

The integer division should truncate toward zero.



Example 1:

Input: s = "3+2*2"
Output: 7
Example 2:

Input: s = " 3/2 "
Output: 1
Example 3:

Input: s = " 3+5 / 2 "
Output: 5


Constraints:

1 <= s.length <= 3 * 105
s consists of integers and operators ('+', '-', '*', '/') separated by some number of spaces.
s represents a valid expression.
All the integers in the expression are non-negative integers in the range [0, 231 - 1].
The answer is guaranteed to fit in a 32-bit integer.
*/

//-------------------------------------------
// wrong answer
func calculate(s string) int {
	s = trimSpace(s)
	n := len(s)
	stack := utils.NewStack()
	for i := 0; i< n;i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			stack.Push(int(c-'0'))
		}

		if c == '+' || c == '-' {
			stack.Push(int(c))
		}

		if c == '*' || c == '/' {
			n1 := stack.Pop().(int)
			n2 := int(s[i+1] - '0')
			if c == '*' {
				stack.Push(n1 * n2)
			} else if c == '/' {
				stack.Push(n1/n2)
			}
			i++
		}
	}

	for stack.Len() != 1 {
		n1,op,n2 := stack.Pop().(int),stack.Pop().(int),stack.Pop().(int)
		if op == int('+') {
			stack.Push(n1+n2)
		} else if op == int('-') {
			stack.Push(n1 - n1)
		}
	}

	return stack.Top().(int)
}

func trimSpace(s string) string {
	var ss []byte
	for i := 0; i< len(s);i++ {
		if s[i] != ' ' {
			ss = append(ss,s[i])
		}
	}

	return string(ss)
}


//------------------------------------
// 使用栈保存数字，如果该数字之前的符号是加或者减，那么将当前数字压入栈中，注意如果是减号，则压入当前数字的相反数（原因后面再说）。如果当前符号是乘或者除号，
//那么从栈顶取出一个数字和当前数字进行乘和除的运算，再把结果压入栈中。完成一遍遍历后，把栈中所有数字加起来就是最终的结果。
func calculate2(s string)int {
	res,num,n := 0,0,len(s)
	var op uint8 = '+'
	stack := utils.NewStack()

	for i := 0; i< n;i++ {
		if s[i] >= '0' {
			num = num * 10 + int(s[i] - '0')
		}

		if s[i] < '0' && s[i] != ' ' || i == n-1 {
			if op == '+' {
				stack.Push(num)
			}

			if op == '-' {
				stack.Push(-num)
			}

			if op == '*' || op == '/' {
				var tmp int
				if op == '*' {
					tmp = stack.Top().(int)*num
				}
				if op == '/' {
					tmp = stack.Top().(int)/num
				}

				stack.Pop()
				stack.Push(tmp)
			}

			op = s[i]
			num = 0
		}
	}

	for !stack.Empty() {
		res += stack.Top().(int)
		stack.Pop()
	}

	return res
}