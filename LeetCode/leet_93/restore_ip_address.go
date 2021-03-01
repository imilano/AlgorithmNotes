package leet_93

import "strconv"

/*
Given a string s containing only digits, return all possible valid IP addresses that can be obtained from s.
You can return them in any order.
A valid IP address consists of exactly four integers, each integer is between 0 and 255, separated by single dots and cannot have
leading zeros. For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses and "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses.
*/

//--------------------------------------------
// Iterative
func restoreIpAddresses(s string) []string {
	var res []string
	length := len(s)
	if length == 0 {
		return res
	}

	for i := 1; i < 4 && i < length-2; i++ {
		for j := i + 1; j < i+4 && j < length-1; j++ {
			for k := j + 1; k < j+4 && k < length; k++ {
				s1, s2, s3, s4 := s[0:i], s[i:j], s[j:k], s[k:length]
				if isValid(s1) && isValid(s2) && isValid(s3) && isValid(s4) {
					res = append(res, s1+"."+s2+"."+s3+"."+s4)
				}
			}
		}
	}

	return res
}

func isValid(s string) bool {
	if len(s) > 3 || len(s) == 0 || (s[0] == '0' && len(s) > 1) {
		return false
	}

	p, _ := strconv.Atoi(s)
	if p > 255 {
		return false
	}

	return true
}

//-----------------------------------------------------
func restoreIpAddresses2(s string) []string {
	var res []string
	length := len(s)
	if length == 0 {
		return res
	}

	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			for c := 1; c <= 3; c++ {
				for d := 1; d <= 3; d++ {
					if a+b+c+d == length {
						if s[0:a] <= "255" && s[a:a+b] <= "255" && s[a+b:a+b+c] <= "255" && s[a+b+c:] <= "255" {
							ss := s[0:a] + "." + s[a:a+b] + "." + s[a+b:a+b+c] + "." + s[a+b+c:]
							if len(ss) == length+3 {
								res = append(res, ss)
							}
						}
					}
				}
			}
		}
	}

	return res
}

//----------------------------------------------------
// DFS
