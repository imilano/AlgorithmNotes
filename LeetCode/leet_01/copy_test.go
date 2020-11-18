package leet_01

import (
	"fmt"
	"testing"
)

func TestCopySlice(t *testing.T) {
	a := []int{1,2,4,5}
	b := make ([]int,len(a))

	copy(b,a)

	b[0]= 55
	fmt.Println(a)
}

func TestMap(t *testing.T) {
	a := []int{1,2,2}

	m := make(map[int]int,3)

	for k,v := range a {
		m[v] = k
	}

	fmt.Println(m)
}
