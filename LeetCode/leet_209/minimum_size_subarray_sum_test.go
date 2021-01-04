package leet_209

import (
	"fmt"
	"testing"
)

func TestMinSubarrayLen(t *testing.T) {
	nums := []int {2,3,1,2,4,3}
	s := 7
	fmt.Println(minSubArrayLen(s,nums))
}
