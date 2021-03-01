package leet_289

import (
	"reflect"
	"testing"
)

func TestGameOfLife(t *testing.T){
	arrs := [][][]int{
		{
			{0,1,0},
			{0,0,1},
			{1,1,1},
			{0,0,0},
		},
	}


	ans := [][][]int{
		{
			{0,0,0},
			{1,0,1},
			{0,1,1,},
			{0,1,0},
		},
	}

	for i := range arrs {
		gameOfLife2(arrs[i])
		if !reflect.DeepEqual(arrs[i],ans[i]) {
			t.Errorf("test fail: want %v,get %v",ans[i],arrs[i])
		}
	}
}