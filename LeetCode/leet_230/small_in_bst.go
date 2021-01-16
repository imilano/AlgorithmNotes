package leet_230

import "lightsinger.life/algorithm/leetcode/utils"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := utils.NewStack()
	cur := root

	for !stack.Empty() || cur != nil { // Don't forget to check nil for cur and remember to use ORing ,not and.
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}

		cur = stack.Pop().(*TreeNode)
		k-- // k-- must be placed before if statement
		if k == 0 {
			break
		}
		cur = cur.Right
	}

	return cur.Val
}
