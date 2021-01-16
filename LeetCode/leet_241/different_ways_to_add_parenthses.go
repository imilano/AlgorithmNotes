package leet_241

import (
	"strconv"
)

/*
Given a string of numbers and operators, return all possible results from computing all the different possible ways to group numbers and operators. The valid operators are +, - and *.

Example 1:

Input: "2-1-1"
Output: [0, 2]
Explanation:
((2-1)-1) = 0
(2-(1-1)) = 2
Example 2:

Input: "2*3-4*5"
Output: [-34, -14, -10, -10, 10]
Explanation:
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10
*/

//--------------------------------
// Divide and conquer
func diffWaysToCompute(input string) []int {
	var res []int
	for i := 0; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			left := diffWaysToCompute(input[:i]) // 相当于从操作符开始进行划分，left和right分别代表左右子串中采取不同的划分方式时可以得到的值。从而在计算本次的值时，只需要在左右子串的不同值中选择一种来与当前运算符进行组合即可。
			right := diffWaysToCompute(input[i+1:])

			for j := 0; j < len(left); j++ {
				for k := 0; k < len(right); k++ {
					if input[i] == '+' {
						res = append(res, left[j]+right[k])
					} else if input[i] == '-' {
						res = append(res, left[j]-right[k])
					} else {
						res = append(res, left[j]*right[k])
					}
				}
			}
		}
	}

	if len(res) == 0 {
		num, _ := strconv.Atoi(input)
		res = append(res, num)
	}

	return res
}

//仍然是上面的方法，只是使用hashmap进行优化
var m = make(map[string][]int)

func diffWaysToCompute2(input string) []int {
	if v, ok := m[input]; ok {
		return v
	}
	var res []int
	for i := 0; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			left := diffWaysToCompute(input[:i]) // 相当于从操作符开始进行划分，left和right分别代表左右子串中采取不同的划分方式时可以得到的值。从而在计算本次的值时，只需要在左右子串的不同值中选择一种来与当前运算符进行组合即可。
			right := diffWaysToCompute(input[i+1:])

			for j := 0; j < len(left); j++ {
				for k := 0; k < len(right); k++ {
					if input[i] == '+' {
						res = append(res, left[j]+right[k])
					} else if input[i] == '-' {
						res = append(res, left[j]-right[k])
					} else {
						res = append(res, left[j]*right[k])
					}
				}
			}
		}
	}

	if len(res) == 0 {
		num, _ := strconv.Atoi(input)
		res = append(res, num)
	}
	m[input] = res
	return res
}
