package leet_290

import "testing"

func TestWordPattern(t *testing.T) {
	patterns := []string {
		"aaaa",
		"abba",
		"abba",
		"abba",
	}
	strs := []string{
		"dog cat cat dog",
		"dog dog dog dog",
		"dog cat cat dog",
		"dog cat cat fish",
	}
	ans := []bool{false,false,true,false}

	for i := range patterns {
		r := wordPattern(patterns[i],strs[i])
		if r != ans[i] {
			t.Errorf("test fail: index %d, want %v, get %v",i,ans[i],r)
		}
	}

}
