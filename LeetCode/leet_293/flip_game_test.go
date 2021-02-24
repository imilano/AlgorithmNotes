package leet_293

import (
	"reflect"
	"testing"
)

func TestFlipGame(t *testing.T) {
	strs := []string{"++++"}
	ans := [][]string{
		{"--++","+--+","++--"},
	}

	for i := range strs {
		r := generatePossibleNextMoves(strs[i])
		if !reflect.DeepEqual(r,ans[i]) {
			t.Errorf("test fail")
		}
	}

}
