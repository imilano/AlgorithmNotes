package leet_110

/*
	Given a binary tree, determine if it is height-balanced.
		For this problem, a height-balanced binary tree is defined as:
   			 a binary tree in which the left and right subtrees of every node differ in height by no more than 1.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(depth(root.Right), depth(root.Left)) + 1
}

func abs(x, y int) bool {
	diff := x - y

	return diff <= 1 && diff >= -1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left, right := depth(root.Left), depth(root.Right)

	return abs(left, right) && isBalanced(root.Left) && isBalanced(root.Right)
}
