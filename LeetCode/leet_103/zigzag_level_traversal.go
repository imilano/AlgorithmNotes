package leet_103

/*
	Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

	For example:
	Given binary tree [3,9,20,null,null,15,7],

		3
	   / \
	  9  20
		/  \
	   15   7
	return its zigzag level order traversal as:
	[
	  [3],
	  [20,9],
	  [15,7]
	]
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	var cnt int
	var que []*TreeNode

	if root == nil {
		return res
	}
	que = append(que, root)
	order := true
	for len(que) != 0 {
		var s []int
		cnt = len(que)

		for i := 0; i < cnt ; i++ {
			node := que[0]
			que = que[1:]

			// 如果order，正序添加，如果!order，则逆序添加
			if order {
				s = append(s,node.Val)
			} else {
				s = append([]int{node.Val},s...)
			}

			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}


			//if order {
			//	if node.Left != nil {
			//		que = append(que,node.Left)
			//	}
			//	if node.Right != nil {
			//		que = append(que,node.Right)
			//	}
			//} else {
			//	if node.Right != nil {
			//		que = append(que,node.Right)
			//	}
			//
			//	if node.Left != nil {
			//		que = append(que,node.Left)
			//	}
			//}
		}

		order = !order
		res = append(res, s)
	}

	return res
}
