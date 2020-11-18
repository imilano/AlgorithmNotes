package leet_102

/*
	Binary Tree Level Order Traversal
	Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

	 For example:
	Given binary tree [3,9,20,null,null,15,7],
		3
	   / \
	  9  20
		/  \
	   15   7

	return its level order traversal as:
	[
	  [3],
	  [9,20],
	  [15,7]
	]
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// LevelOrder with Queue
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var que []*TreeNode

	if root == nil {
		return res
	}

	que = append(que,root)
	for len(que) != 0 {
		var s []int

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

		res = append(res,s)
	}

	return res
}