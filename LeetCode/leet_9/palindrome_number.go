package leet_9


/*
	Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
	Follow up: Could you solve it without converting the integer to a string?
 */
func reverse(x int) int {
	var res []int

	for x != 0{
		res= append(res,x%10)
		x /= 10
	}

	var r int
	i := 0
	for i < len(res) {
		r = r * 10 + res[i]
		i++
	}

	return r
}

func isPalindrome(x int) bool {
	if x < 0  || (x%10 == 0 && x!= 0){
		return false
	}

	res := reverse(x)
	if res != x {
		return false
	}

	return true
}

// more concise one
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	y := x
	res := 0
	for y != 0{
		res = res *10 + y%10
		y /= 10
	}

	return res == x
}
