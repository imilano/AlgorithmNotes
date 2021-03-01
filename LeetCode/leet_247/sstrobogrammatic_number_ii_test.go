package leet_247

import (
	"fmt"
	"testing"
)

func TestFindStrobogrammaticNumber(t *testing.T) {
	ns := []int{
		0, 1, 2, 3, 4,
	}

	for _, v := range ns {
		fmt.Println(v, ": ", findStrobogrammatic(v))
	}
}
