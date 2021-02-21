package leet_280

import (
	"reflect"
	"testing"
)

func TestWiggleSort(t *testing.T) {
	nums := [][]int{
		[]int{3,5,2,1,6,4},
	}

	ans := [][]int{
		[]int{1,3,2,5,4,6},
	}

	for i := range nums {
		r := make([]int,len(nums[i]))
		copy(r,nums[i])
		wiggleSort(r)
		if !reflect.DeepEqual(r,ans[i]) {
			t.Errorf("test fail: index %d,want %v, get %v",i,ans[i],r)
		}
	}
}
