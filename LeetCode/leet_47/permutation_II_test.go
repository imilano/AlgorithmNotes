package leet_47

import (
	"fmt"
	"testing"
)

func TestPermutationUnique(t *testing.T) {
	nums := []int {1,2,2}
	a := permuteUnique(nums)
	fmt.Print(a)
}
