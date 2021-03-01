package leet_257

import (
	"strconv"
)

/*
	Given a binary tree, return all root-to-leaf paths.
		Note: A leaf is a node with no children.
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, path string, ans *[]string) {
	if root.Left == nil && root.Right == nil {
		path = path + strconv.Itoa(root.Val)
		*ans = append(*ans, path)
		return
	}

	if root.Left != nil {
		helper(root.Left, path+strconv.Itoa(root.Val)+"->", ans)
	}

	if root.Right != nil {
		helper(root.Right, path+strconv.Itoa(root.Val)+"->", ans)
	}
}

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}

	helper(root, "", &res)
	return res
}
