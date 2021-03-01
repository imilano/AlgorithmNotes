package leet_266

import "testing"

func TestCanPermutePalindrome(t *testing.T) {
	strs := []string{
		"code",
		"aad",
		"carerac",
	}
	ans := []bool{false, true, true}

	for i := range strs {
		r := canPermutePalindrome2(strs[i])
		if r != ans[i] {
			t.Errorf("test fail: index %v, want %v, get %v", i, ans[i], r)
		}
	}
}
