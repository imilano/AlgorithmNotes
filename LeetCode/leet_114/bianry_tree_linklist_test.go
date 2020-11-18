package leet_114

import (
	"fmt"
	"testing"
)

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%v ",root.Val)
	traverse(root.Left)
	traverse(root.Right)
}

func TestFlattern(t *testing.T) {
	node := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}

	Flattern(node)
	traverse(node)
}
