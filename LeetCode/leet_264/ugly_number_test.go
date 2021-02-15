package leet_264

import "testing"

func TestNthUglyNumber(t *testing.T) {
	nums :=  []int{1,2,3,4,5,6,8,9,10,12}
	for i := range nums {
		r := nthUglyNumber(i+1)
		if r != nums[i] {
			t.Errorf("test fail: index %d, want %d, get %d",i,nums[i],r)
		}
	}
}
