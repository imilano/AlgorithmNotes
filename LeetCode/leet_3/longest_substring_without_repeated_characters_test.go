package leet_03

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "hello,world"
	//fmt.Println(LengthOfLongestSubstring(s))
	//
	//m :=  make(map[int]int)
	//fmt.Println(m[1])
	fmt.Println(LengthOfLongestSubstringOptimized(s))
}
