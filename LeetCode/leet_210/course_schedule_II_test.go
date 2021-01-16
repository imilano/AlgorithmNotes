package leet_210

import (
	"fmt"
	"testing"
)

func TestFindOrder(t *testing.T) {
	n := 4
	pre := [][]int{
		{1, 0},
		{2, 0},
		{3, 1},
		{3, 2},
	}

	fmt.Println(findOrder(n, pre))
}
