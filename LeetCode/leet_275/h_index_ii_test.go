package leet_275

import (
	"fmt"
	"testing"
)

func TestHIndex(t *testing.T) {
	citations := []int{0,1,3,5,6}
	fmt.Println(hIndex(citations))
}
