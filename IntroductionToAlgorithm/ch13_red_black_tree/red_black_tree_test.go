package ch13_red_black_tree

import (
	"fmt"
	"testing"
)

func TestRBT(t *testing.T) {
	nums := []int{1,5,6,7,8,9,10,11,12,13,14,15}
	tree := &Tree{
		Root: nil,
		size: 0,
	}

	for _,v := range nums {
		fmt.Printf("---------------Insert %d: --------------------\n",v)
		tree.Put(v,v)
		fmt.Println(tree.String())
	}


	fmt.Println("------------------------------------")
	tree.Remove(14)
	tree.Remove(9)
	tree.Remove(5)

	fmt.Println(tree.String())
}
