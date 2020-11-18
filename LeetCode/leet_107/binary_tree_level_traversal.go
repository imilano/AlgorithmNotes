package leet_107
/*
	Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).
	For example:
	Given binary tree [3,9,20,null,null,15,7],
		3
	   / \
	  9  20
		/  \
	   15   7

	return its bottom-up level order traversal as:
	[
	  [15,7],
	  [9,20],
	  [3]
	]
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	var que []*TreeNode

	if root == nil {
		return res
	}

	que = append(que,root)
	for len(que) != 0 {
		var s []int
		var t [][]int

		cnt := len(que)
		for i := 0;i < cnt;i++ {
			node := que[0]
			que = que[1:]

			s = append(s, node.Val)
			if node.Left != nil {
				que = append(que,node.Left)
			}
			if node.Right != nil {
				que = append(que,node.Right)
			}
		}

		// add s to the start of res
		t = append(t,s)
		res = append(t,res...)
	}

	return res
}