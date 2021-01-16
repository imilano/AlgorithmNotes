package leet_124

import "math"

/*
	Given a non-empty binary tree, find the maximum path sum.
	For this problem, a path is defined as any node sequence from some starting node to any node in the tree along
	the parent-child connections. The path must contain at least one node and does not need to go through the root.
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func traverse(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	// goes all the way down to the left
	left := max(traverse(root.Left, maxSum), 0)
	right := max(traverse(root.Right, maxSum), 0)

	// update the maxSum on every node
	*maxSum = max(*maxSum, left+right+root.Val)

	// return to its upper path
	return max(left, right) + root.Val
}

func maxPathSum(root *TreeNode) int {
	var sum int

	if root == nil {
		return 0
	}

	// initialize sum to negative min to allow return negative number
	sum = math.MinInt64
	traverse(root, &sum)

	return sum
}
