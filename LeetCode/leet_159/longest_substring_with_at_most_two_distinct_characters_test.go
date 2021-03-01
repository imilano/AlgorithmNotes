package leet_159

import (
	"fmt"
	"testing"
)

func TestLength(t *testing.T) {
	s := []string{
		"eceba",
		"ccaabbb",
	}

	for _, v := range s {
		fmt.Printf("%s: %d\n", v, lengthOfLongestSubstringTwoDistinct(v))
	}
}
