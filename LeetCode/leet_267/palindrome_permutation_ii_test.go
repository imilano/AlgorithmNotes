package leet_267

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGeneratePalindrome(t *testing.T) {
	strs :=[]string{
		"aabb","aba",
	}
	ans := [][]string {
		[]string{"abba","baab"},
		[]string{"aba"},
	}

	for i := range strs {
		r := generatePalindromes2(strs[i])
		if !reflect.DeepEqual(r,ans[i]) {
			t.Errorf("test fail: index %d, want %v, get %v",i,ans[i],r)
		}
	}
}

func TestReverseString(t *testing.T) {
	strs := []string {
		"aabb","aaabbb","aba",
		"abcba","abccba",
	}
	ans := []string {
		"bbaa","bbbaaa","aba","abcba","abccba",
	}

	for i := 0; i < len(strs);i++ {
		r := reverseString(strs[i])
		if r != ans[i] {
			t.Errorf("test fail: index %v, want %v, get %v",i,ans[i],r)
		}
	}
}
func TestSwapString(t *testing.T) {
}

func TestUint8ToString(t *testing.T) {
	fmt.Println(string('a'))
}
