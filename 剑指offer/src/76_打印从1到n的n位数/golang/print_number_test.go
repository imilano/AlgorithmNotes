package golang

import (
	"fmt"
	"sort"
	"testing"
)

func TestPrintNumber(t *testing.T) {
	ns := []int{1,2,3}
	for i := range ns {
		r := PrintNumber(ns[i])
		sort.Strings(r)
		fmt.Println(r)
	}
}

func TestRune(t *testing.T) {
	var i rune = '0'
	for ;i <='9';i++ {
		fmt.Println(string(i))
	}
}
