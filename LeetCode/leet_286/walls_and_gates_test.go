package leet_286

import (
	"math"
	"reflect"
	"testing"
)

func TestWallsAndGates(t *testing.T) {
	matrix := [][]int{
		[]int{math.MaxInt32, -1, 0, math.MaxInt32},
		[]int{math.MaxInt32, math.MaxInt32, math.MaxInt32, -1},
		[]int{math.MaxInt32, -1, math.MaxInt32, -1},
		[]int{0, -1, math.MaxInt32, math.MaxInt32},
	}

	ans := [][]int{
		{3, -1, 0, 1},
		{2, 2, 1, -1},
		{1, -1, 2, -1},
		{0, -1, 3, 4},
	}

	wallsAndGates2(matrix)
	if !reflect.DeepEqual(matrix, ans) {
		t.Errorf("test fail: want %v, get %v", ans, matrix)
	}
}
