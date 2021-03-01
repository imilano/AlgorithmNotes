package leet_112

/*
	Path Sum
	Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.
	Note: A leaf is a node with no children.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursive version
func hasPathSumRec(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

// With DFS
type Node struct {
	val  int
	node *TreeNode
}

func hasPathWithDFS(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	var que []Node
	var node Node
	que = append(que, Node{
		val:  root.Val,
		node: root,
	})

	for len(que) != 0 {
		node = que[0]
		que = que[1:]

		if node.node.Left == nil && node.node.Right == nil && node.val == sum {
			return true
		}

		if node.node.Right != nil {
			que = append(que, Node{
				val:  node.val + node.node.Right.Val,
				node: node.node.Right,
			})
		}

		if node.node.Left != nil {
			que = append(que, Node{
				val:  node.val + node.node.Left.Val,
				node: node.node.Left,
			})
		}
	}

	return false
}

func hasPathSum(root *TreeNode, sum int) bool {

	// recursive version
	// return hasPathSumRec(root,sum)

	// DFS with Queue
	return hasPathWithDFS(root, sum)
}
