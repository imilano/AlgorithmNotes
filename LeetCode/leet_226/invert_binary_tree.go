package leet_226

/*
	Invert a binary tree.
	Example:
		Input:
			 4
		   /   \
		  2     7
		 / \   / \
		1   3 6   9
		Output:
			 4
		   /   \
		  7     2
		 / \   / \
		9   6 3   1
*/

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func revert(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left, right := revert(root.Left), revert(root.Right)
	root.Left, root.Right = right, left
	return root
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	return revert(root)
}
