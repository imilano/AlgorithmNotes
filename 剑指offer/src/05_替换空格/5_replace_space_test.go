package _5_替换空格

import (
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	ss := []string {
		"a b b c c ",
		"a b c d",
	}

	ans := []string {
		"a%20b%20b%20c%20c%20",
		"a%20b%20c%20d",
	}

	for i := range ss {
		r := ReplaceSpace(ss[i])
		if r != ans[i] {
			t.Errorf("test fail, index %d, want %v, get %v",i,ans[i],r)
		}
	}
}
