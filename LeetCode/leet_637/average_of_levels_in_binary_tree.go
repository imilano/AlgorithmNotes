package leet_637

/*
	Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.

	Example 1:
	Input:
		3
	   / \
	  9  20
		/  \
	   15   7
	Output: [3, 14.5, 11]
	Explanation:
	The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].

	Note:
		The range of node's value is in the range of 32-bit signed integer.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	var que []*TreeNode
	var cnt int

	if root == nil {
		return res
	}

	que = append(que, root)
	for len(que) != 0 {
		var sum float64
		var node *TreeNode

		cnt = len(que)
		for i := 0; i < cnt; i++ {
			node = que[0]
			que = que[1:]
			sum += float64(node.Val)

			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}

		res = append(res, sum/float64(cnt))
	}

	return res
}
