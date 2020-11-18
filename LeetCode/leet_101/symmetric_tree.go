package leet_101

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// Recursive way
func isSymmetricRec(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if right == nil || left == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricRec(left.Left,right.Right) && isSymmetricRec(left.Right,right.Left)
}

// Iterative way, use queue
func check(p,q *TreeNode) bool {
	if p== nil && q== nil {
		return true
	}

	if p == nil || q== nil {
		return false
	}

	if q.Val != p.Val {
		return false
	}

	return true
}

func isSymmetricWithQueue (root *TreeNode) bool {
	var que []*TreeNode
	var first,second *TreeNode

	que = append(que,root,root)
	for len(que) != 0 {
		first,second = que[0],que[1]
		que = que[2:]

		if !check(first,second) {
			return false
		}

		if first != nil && second != nil {
			que = append(que,first.Left,second.Right)
			que = append(que,first.Right,second.Left)
		}
	}

	return true
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// recursive way
	return isSymmetricRec(root,root)

	// iterative using queue
	//return isSymmetricWithQueue(root)
}
