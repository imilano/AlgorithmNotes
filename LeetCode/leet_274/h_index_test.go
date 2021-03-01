package leet_274

import (
	"fmt"
	"testing"
)

func TestHIndex(t *testing.T) {
	citations := []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(citations))
}
