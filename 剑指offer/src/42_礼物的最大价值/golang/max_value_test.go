package golang

import (
	"fmt"
	"testing"
)

func TestMaxValue(t *testing.T) {
	arr := [][][]int{
		{{1,10,3,8},
		{12,2,9,6},
		{5,7,4,11},
		{3,7,16,5},
	}}

	ans := []int{53}
	for i := range arr {
		r := MaxValue(arr[i])
		if r != ans[i] {
			fmt.Printf("test fail, index %d, want %d, get %d",i,ans[i],r)
		}
	}

}
