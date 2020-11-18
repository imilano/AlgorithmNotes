package leet_91

import (
	"fmt"
	"testing"
)

func TestNumDecodings(t *testing.T) {
	s := "226"

	fmt.Println(string(s[1]) + string(s[2]))
	numDecodings(s)
}
