package leet_9

import "strconv"

/*
	Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
	Follow up: Could you solve it without converting the integer to a string?
*/

/*
最简单的，当然是转换为string之后进行检测了。第二个方法的话，就是为位操作，获得它对应的回文数，看二者是否相等。
*/

// ----------------------------------------
// Original, not concise
func reverse(x int) int {
	var res []int

	for x != 0 {
		res = append(res, x%10)
		x /= 10
	}

	var r int
	i := 0
	for i < len(res) {
		r = r*10 + res[i]
		i++
	}

	return r
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	res := reverse(x)
	if res != x {
		return false
	}

	return true
}

//-----------------------------------
func isPalindrome2(x int) bool {
	s := strconv.Itoa(x)
	return judge(s)
}

func judge(s string) bool {
	start, end := 0, len(s)-1
	for start <= end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}

//------------------------------------------
// more concise one
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	y := x
	res := 0
	for y != 0 {
		res = res*10 + y%10
		y /= 10
	}

	return res == x
}
