package golang

import (
	"fmt"
	"testing"
)

func TestGetIndex(t *testing.T) {
	nums := [][]int{
		{2,2,3,5,4,4},
		{3,3,4,5,6,6},
	}

	for i:= range nums {
		fmt.Println(GetIndex(nums[i]))
	}
}

func GetIndex(array []int) int {
	res := array[0]
	n := len(array)
	for i := 1; i < n; i++ {
		res ^= array[i]
	}

	var index int
	for i := 0; i < 32;i++ {
		if res & 1 == 1 {
			index = i
			break
		}

		res >>= 1
	}

	return index
}


func TestFindNumsAppearOnce(t *testing.T) {
	nums := []int{2,4,3,6,3,2,5,5}
	fmt.Println(FindNumsAppearOnce(nums))
}
