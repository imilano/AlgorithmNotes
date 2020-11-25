package leet_125

/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.
 */


// 字母数字都要考虑
func isPalindrome(s string) bool {
	var t []uint8
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			t = append(t,s[i])
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			t = append(t,s[i]+32)
		} else if s[i] >='0' && s[i] <= '9'{
			t = append(t, s[i])
		}
	}

	length := len(t)
	if length == 0 ||length == 1 {
		return true
	}

	left,right := 0,length-1
	for left <= right {
		if t[left] != t[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func abs(a,b uint8) uint8 {
	r := a -b
	if r < 0 {
		r = -r
	}

	return  r
}





