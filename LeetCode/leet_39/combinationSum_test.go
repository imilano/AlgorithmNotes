package leet_39

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	nums := []int{2,3,6,7}
	target := 7
	res := combinationSum(nums,target)
	fmt.Println(res)
}
