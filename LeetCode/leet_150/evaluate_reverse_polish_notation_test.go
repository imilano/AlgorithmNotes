package leet_150

import (
	"fmt"
	"testing"
)

func TestEvalRPM( t *testing.T) {
	s := []string{"4","13","5","/","+"}
	fmt.Print(evalRPN(s))
}
