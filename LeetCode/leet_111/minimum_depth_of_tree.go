package leet_111

/*
	Minimum depth of binary tree.
		Given a binary tree, find its minimum depth.
		The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
		Note: A leaf is a node with no children.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func depth(root *TreeNode) int {
	// A leaf node must have no child
	if root == nil {
		return 0
	}

	// If left node is null, then the height is the height of its right subtree plus 1
	if root.Left == nil {
		return depth(root.Right) + 1
	}

	// if right node is null, then the height is the height of its left subtree plug 1
	if root.Right == nil {
		return depth(root.Left) + 1
	}

	return min(depth(root.Left), depth(root.Right)) + 1
}

// TODO Too many if statement, which leads to lower performance
func minDepthUseBFS(root *TreeNode) int {
	var dep, cnt int
	var que []*TreeNode

	if root == nil {
		return 0
	}
	que = append(que, root)
	dep = 1
	for len(que) != 0 {
		cnt = len(que)
		for i := 0; i < cnt; i++ {
			node := que[0]
			que = que[1:]

			if node.Left == nil && node.Right == nil { // A leaf node must have no child. If we find a leaf node, then we find the min depth.
				return dep
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
			if node.Left != nil {
				que = append(que, node.Left)
			}
		}
		dep++
	}

	return dep
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// BFS bersion
	// return minDepthUseBFS(root)

	// recursive way
	return depth(root)
}
