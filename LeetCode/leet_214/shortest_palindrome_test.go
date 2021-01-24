package leet_214

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	strs := []string{"abc", "aa", "abca"}
	for _, s := range strs {
		fmt.Println(s, " : ", shortestPalindrome2(s))
	}
}
