package leet_429

/*
	Given an n-ary tree, return the level order traversal of its nodes' values.
	Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).
*/

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var que []*Node
	var cnt int
	var res [][]int

	if root == nil {
		return res
	}

	que = append(que, root)
	for len(que) != 0 {
		var lev []int
		var cur *Node
		cnt = len(que)
		for i := 0; i < cnt; i++ {
			cur = que[0]
			que = que[1:]

			lev = append(lev, cur.Val)
			for _, node := range cur.Children {
				que = append(que, node)
			}
		}

		res = append(res, lev)
	}

	return res
}
