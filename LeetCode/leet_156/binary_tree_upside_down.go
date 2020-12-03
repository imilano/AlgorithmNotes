package leet_156

/*
Given a binary tree where all the right nodes are either leaf nodes with a sibling (a left node that shares the same parent node) or empty, flip it upside down and turn it into a tree where the original right nodes turned into left leaf nodes. Return the new root.

Example:

Input: [1,2,3,4,5]

    1
   / \
  2   3
 / \
4   5

Output: return the root of the binary tree [4,5,2,#,#,3,1]

   4
  / \
 5   2
    / \
   3   1
*/
//  Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// ------------------------------------------------------
// 递归
// 将原树根节点的左节点作为新的根节点，将原根节点的右节点作为新根节点左节点，将原根节点作为新根节点右节点。主要要将新根节点的右节点的两个子节点设为空
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	left,right := root.Left,root.Right
	res := upsideDownBinaryTree(left)
	left.Left = right
	left.Right = root
	root.Left = nil
	root.Right = nil
	return res
}

// -----------------------
// 迭代版，遵循上述思路
func upsideDownBinaryTreeIte(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	cur := root
	var pre,next,tmp *TreeNode
	for cur != nil {
		next = cur.Left
		cur.Left = tmp
		tmp = cur.Right
		cur.Right = pre
		pre = cur
		cur = next
	}
	return pre
}