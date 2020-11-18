package leet_222

/*
	Given a complete binary tree, count the number of nodes.

	Note:
		Definition of a complete binary tree from Wikipedia:
			In a complete binary tree every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.
	Example:
		Input:
			1
		   / \
		  2   3
		 / \  /
		4  5 6
		Output: 6
*/


//  Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func left(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return 1 + left(root.Left)
}

func right(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return right(root.Right) + 1
}

func  count(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return count(root.Left) + count(root.Right) +1
}

func countNodes(root *TreeNode) int {
    // recursive way
    // return count(root)

    if root == nil {
        return 0
    }

    left,right := left(root),right(root)
    if left == right {
        return 1 << left -1
    } else {
        return count(root)
    }
}