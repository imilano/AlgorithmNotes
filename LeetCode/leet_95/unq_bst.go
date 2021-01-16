package leet_95

/*
	Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.
*/

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generate(start int, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil} // 结果是不同的树，所以需要把空节点nil也添加进去
	}

	result := []*TreeNode{}

	// 枚举可行根节点
	for i := start; i <= end; i++ {
		// 获取所有可行的左子树集合
		left := generate(start, i-1)
		// 获取所有可行的右子树集合
		right := generate(i+1, end)
		// 从左子树集合中选出一棵左子树，再从右子树集合中选出一棵右子树，拼接到根节点上。
		for _, l := range left {
			for _, r := range right {
				result = append(result, &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return result
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return generate(1, n)
}
