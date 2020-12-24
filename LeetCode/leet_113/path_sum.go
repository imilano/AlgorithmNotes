package leet_113

/*
	Path Sum II
	Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
	Note: A leaf is a node with no children.

	Example:
	Given the below binary tree and sum = 22,
		  5
		 / \
		4   8
	   /   / \
	  11  13  4
	 /  \    / \
	7    2  5   1

	Return:
	[
	   [5,4,11,2],
	   [5,8,4,5]
	]
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// Path Sum using DFS
func dfs(root *TreeNode,sum int,path []int,res *[][]int) {
	path = append(path,root.Val)
	if root.Left == nil && root.Right == nil && sum == root.Val {
		tmp := make([]int, len(path))
		copy(tmp,path)
		*res = append(*res, tmp)
	}

	if root.Left != nil {
		dfs(root.Left,sum-root.Val,path,res)
	}

	if root.Right != nil {
		dfs(root.Right,sum-root.Val,path,res)
	}
}

func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	dfs(root,sum,nil,&res)

	return res
}

//--------------------------------
func helper(root *TreeNode, sum int, cur []int, res *[][]int) {
	cur = append(cur,root.Val)
	sum -= root.Val
	if root.Left == nil && root.Right == nil && sum == 0 {
		t := make([]int,len(cur))
		copy(t,cur)
		*res = append(*res, t)
		return
	}

	if root.Left != nil {
		helper(root.Left,sum, cur,res)
	}
	if root.Right != nil {
		helper(root.Right, sum, cur,res)
	}
}