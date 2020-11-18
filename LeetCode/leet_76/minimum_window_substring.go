package leet_76

import (
	"math"
)

/*
Given two strings s and t, return the minimum window in s which will contain all the characters in t.
If there is no such window in s that covers all characters in t, return the empty string "".
Note that If there is such a window, it is guaranteed that there will always be only one unique minimum window in s.

Example 1:
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Example 2:
Input: s = "a", t = "a"
Output: "a"
 */


//------------------------------------------------------
// Sliding window
// Details here: https://www.cnblogs.com/grandyang/p/4340948.html

// time complexity：O(n)
// space complexity: O(1)
// 1. Use two pointer left and right to represent a window
// 2. Move end to find a valid window
// 3. When a valid window was found, move left to find a smaller window

// TODO how to solve this using dp
func minWindow(s string, t string) string {
	var res string
	letterCnt := make(map[uint8]int,128)
	for i := range t {
		letterCnt[t[i]]++
	}

	left,right,counter,minStart,minLen := 0,0,len(t),0,math.MaxInt32
	// move end to find a valid window
	for right < len(s) {
		// If char in s existed in t, decrease count
		if letterCnt[s[right]] > 0 {
			counter--
		}

		// Decrease letterCnt[s[end]]. If char does not exist in t, m[s[end]] will be negative.
		letterCnt[s[right]]--
		right++

		// When found valid window, then move left to find the smallest window
		for counter ==  0 {
			// update minLen and minStart
			if right-left < minLen {
				minStart = left
				minLen = right - left
			}

			// We decrease it before when moving right,
			// cause we move over the character, we need to increase count
			letterCnt[s[left]]++
			// If cnt is greater than 0 after addition, which means we pass a character that is in t, so we need to update counter
			if letterCnt[s[left]] > 0 {
				counter++
			}

			// keep moving left to narrow the window
			left++
		}
	}

	// TODO figure this out
	if minLen == math.MaxInt32 {
		res = ""
	}
	res = s[minStart:minStart+minLen]

	return res
}



// -------------------------------------------------------
// Original solution
func contain(c uint8, s string) bool {
	for i := range s {
		if c == s[i] {
			return true
		}
	}

	return false
}

func containAll(s string, t string) bool {
	m := make(map[uint8]bool)
	for i := range s{
		m[s[i]] = true
	}

	for i := range t {
		if _,ok := m[t[i]]; !ok {
			return false
		}
	}

	return true
}

func minWindowOriginal(s string, t string) string {
	var res string
	lens:= len(s)
	m := make(map[uint8]bool)
	for i := range t {
		m[t[i]] = true
	}
	dp := make([][]uint8,lens)
	for i := range dp {
		dp[i] = make([]uint8,lens)
	}

	// base condition
	for i := 0; i< lens;i++ {
		if contain(s[i],t) {
			dp[i][i] = 1
		} else {
			dp[i][i] = 0
		}
	}

	// status transfer
	for i := 0; i<lens;i++ {
		for j := i+1; j<lens;j++ {
			if _,ok := m[s[j]]; ok {
				dp[i][j] = dp[i][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	// find the minimum
	start,end := math.MinInt8,math.MaxInt8
	for i := lens-1; i>=0 ; i-- {
		for j := lens-1; j>= i;j-- {
			if int(dp[i][j]) == len(t) && containAll(s[i:j+1],t) {
				if j-i < end-start {
					start = i
					end = j
				}
			}
		}
	}

	if start == math.MinInt8 {
		return ""
	}
	res = s[start:end+1]
	return res
}