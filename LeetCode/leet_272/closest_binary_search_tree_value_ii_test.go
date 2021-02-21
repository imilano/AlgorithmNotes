package leet_272

import (
	"fmt"
	"testing"
)

func TestClosestKValue(t *testing.T) {
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

	fmt.Println(closestKValues(tree, 3.714286, 2))
}
