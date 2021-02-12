package leet_246

import (
	"fmt"
	"testing"
)

func TestIsStrobogrammaticNumber(t *testing.T) {
	nums := []string{
		"69",
		"88",
		"962",
		"96",
		"609",
		"33",
	}

	for _, v := range nums {
		fmt.Println(v, ": ", isStrobogrammatic(v))
	}
}
