package leet_67

import "strconv"

/*
Given two binary strings, return their sum (also a binary string).
The input strings are both non-empty and contains only characters 1 or 0.
Example 1:
Input: a = "11", b = "1"
Output: "100"

Example 2:
Input: a = "1010", b = "1011"
Output: "10101"
 */

func AddBinary(a string, b string) string {
	return addBinary(a,b)
}
func addBinary(a string, b string) string {
	var res string
	var carry uint8

	if len(a) > len(b) {
		b,a = a,b
	}


	i,j := len(a)-1,len(b)-1
	for i >=0 && j >=0 {
		sum := a[i] - '0' + b[j]-'0' + carry

		carry = sum / 2
		sum = sum % 2
		res = strconv.Itoa(int(sum)) + res

		i--
		j--
	}


	for j >= 0 {
		sum := b[j] - '0' + carry
		carry = sum / 2
		sum = sum % 2
		res = strconv.Itoa(int(sum)) + res
		j--
	}

	if carry == 1{
		res = "1" + res
	}
	return res
}