package leet_171

import "math"

/*
Given a column title as appear in an Excel sheet, return its corresponding column number.

For example:

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...
*/

// Too many type conversion lower the speed
func titleToNumber(s string) int {
	var res int
	length := len(s)
	index := 0
	for index < length {
		res += int(s[index]-'A'+1)*int(math.Pow(26,float64(length-1-index)))
		index++
	}

	return res
}

func titleToNumber2(s string)int {
	var res int
	length := len(s)
	tmp := 1
	for i := length; i >= 1;i-- {
		res += int(s[i-1]-'A'+1) * tmp
		tmp *= 26
	}

	return res
}
