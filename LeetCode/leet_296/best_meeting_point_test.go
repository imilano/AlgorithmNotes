package leet_296

import "testing"

func TestMinTotalDistance(t *testing.T) {
	grid := [][]int{
		{1,0,0,0,1},
		{0,0,0,0,0},
		{0,0,1,0,0},
	}

	r := minTotalDistance(grid)
	if r != 6 {
		t.Errorf("test fail")
	}
}
