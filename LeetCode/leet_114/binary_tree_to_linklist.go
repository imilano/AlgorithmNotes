package leet_114

/*
	Given a binary tree, flatten it to a linked list in-place.
	For example, given the following tree:

		1
	   / \
	  2   5
	 / \   \
	3   4   6
	The flattened tree should look like:
	1
	 \
	  2
	   \
		3
		 \
		  4
		   \
			5
			 \
			  6
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var pre *TreeNode
func flatten(root *TreeNode)  {
	if root == nil {
		return
	}

	flatten(root.Right)
	flatten(root.Left)

	root.Right = pre
	root.Left = nil
	pre = root
}

func Flattern(root *TreeNode) {
	flatten(root)
}