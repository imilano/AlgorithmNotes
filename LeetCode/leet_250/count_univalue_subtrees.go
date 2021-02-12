package leet_250

/*
Given a binary tree, count the number of uni-value subtrees.

A Uni-value subtree means all nodes of the subtree have the same value.

Example :

Input:  root = [5,1,5,5,5,null,5]

              5
             / \
            1   5
           / \   \
          5   5   5

Output: 4
*/

type TreeNode struct {
	Val int
	Left,Right *TreeNode
}


// 注意这里的unival subtree，这里指的是一棵子树的根节点也属于该子树
var res int
func countUnivalSubtrees(root *TreeNode) int {
	if root == nil {
		return res
	}

	if isUnival(root,root.Val) {
		res++
	}

	countUnivalSubtrees(root.Left)
	countUnivalSubtrees(root.Right)
	return res
}

func isUnival(root *TreeNode,val int) bool {
	if root ==nil {
		return true
	}

	return root.Val == val && isUnival(root.Right,val) && isUnival(root.Left,val)
}
