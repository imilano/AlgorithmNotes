package leet_270

import (
	"fmt"
	"testing"
)

func TestClosestValue(t *testing.T) {
	tree := &TreeNode{
		val: 4,
		left: &TreeNode{
			val: 2,
			left: &TreeNode{
				val:   1,
				left:  nil,
				right: nil,
			},
			right: &TreeNode{
				val:   3,
				left:  nil,
				right: nil,
			},
		},
		right: &TreeNode{
			val:   5,
			left:  nil,
			right: nil,
		},
	}

	fmt.Println(closestValue(tree, 3.714286))
}
