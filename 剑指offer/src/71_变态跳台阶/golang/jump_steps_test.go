package golang

import (
	"fmt"
	"testing"
)

func  TestJumpSteps(t *testing.T) {
	nums := []int{3}
	ans := []int{4}
	for i := range nums {
		r := jumpFloorII(nums[i])
		if r != ans[i] {
			fmt.Printf("index %d, want %d, get %d\n",i,ans[i],r)
		}
	}
}
