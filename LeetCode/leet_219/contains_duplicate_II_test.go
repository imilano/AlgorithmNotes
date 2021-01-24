package leet_219

import (
	"fmt"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1, 2, 3}
	fmt.Println(containsNearbyDuplicate(nums, 2))
}
