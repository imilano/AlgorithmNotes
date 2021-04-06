package golang

import (
	"fmt"
	"testing"
)

func TestLongestSubstringWithoutDup(t *testing.T) {
	strs := []string{
		"arabcacfr",
	}

	ans := []int{4}

	for i := range strs {
		r := LongestSubstringWithoutDup2(strs[i])
		if r != ans[i] {
			fmt.Printf("test fail, index %d, want %d, get %d",i,ans[i],r)
		}
	}
}
