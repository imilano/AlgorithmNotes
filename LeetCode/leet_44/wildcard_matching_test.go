package leet_44

import (
	"fmt"
	"testing"
)

func TestIsMatch(t *testing.T) {
	s := "adceb"
	p := "*a*b"
	fmt.Printf("%v", isMatch(s, p))
}
