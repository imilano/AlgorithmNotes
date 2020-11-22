package leet_91

/*
A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26

Given a non-empty string containing only digits, determine the total number of ways to decode it.

The answer is guaranteed to fit in a 32-bit integer.
 */

func numDecodingsStraight(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}


	dp := make([]int ,length+1)
	dp[0] = 1
	for i:= 1; i<= length; i++ {
		if s[i-1] == '0' {
			dp[i] = 0
		}else {
			dp[i] += dp[i-1] //  如果当前数字可以单独拆分，dp[i] 至少可以有跟 dp[i-1] 一样多的拆分情况(此时dp[i]为0)
		}

		if i > 1 && (s[i-2] == '1' || (s[i-2] == '2' && s[i-1] <= '6' )) {
			dp[i]  += dp[i-2]  // 如果当前数字可以和前一个数字组合，那么就可以构成数列，dp[i] = dp[i-1] + dp[i-2]
		}
	}

	return dp[length]
}


func numDecodings(s string)int {
	length := len(s)
	if length == 0  {
		return 0
	}

	dp := make([]int,length +1)
	// base case
	dp[0]  = 1
	dp[1] = 1
	if s[0] == '0' {
		dp[1] = 0
	}

	// 注意此处的i代表第i个，对应s中的下标为i-1
	for i := 2; i<=length;i++ {
		one := string(s[i-1:i])
		two := string(s[i-2:i])

		if one >= "1" && one <= "9" {
			dp[i] += dp[i-1]
		}

		if two >= "10" && two <= "26" {
			dp[i] += dp[i-2]
		}
	}

	return dp[length]
}

func numDecodingsOriginal(s string) int {
	length := len(s)
	if length == 2 {
		return 2
	}

	if length == 1 {
		return 1
	}


	str := string(s[0]) + string(s[1])
	if str >= "10" && str <= "26" {
		return numDecodings(s[2:]) * 2 + numDecodings(s[1:])
	} else {
		return numDecodings(s[1:])
	}
}
