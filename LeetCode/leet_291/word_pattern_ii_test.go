package leet_291

import (
	"testing"
)

func TestWordPattern(t *testing.T) {
	patterns := []string{
		"abab",
		"aaaa",
		"aabb",
	}
	strs := []string{
		"redblueredblue",
		"asdasdasdasd",
		"xyzabcxyzabc",
	}
	ans := []bool {
		true,true,false,
	}
	for i := range patterns {
		r := wordPattern(patterns[i],strs[i])
		if r != ans[i] {
			t.Errorf("test fail: index %d, want %v, get %v",i,ans[i],r)
		}
	}
}
