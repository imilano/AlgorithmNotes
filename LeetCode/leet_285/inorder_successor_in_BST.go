package leet_285

import (
	"lightsinger.life/algorithm/leetcode/utils"
	"reflect"
)

/*
Given a binary search tree and a node in it, find the in-order successor of that node in the BST.
The successor of a node p is the node with the smallest key greater than p.val.

Note:
If the given node has no in-order successor in the tree, return null.
It's guaranteed that the values of the tree are unique.
*/

type TreeNode struct {
	val         int
	left, right *TreeNode
}

// 简单的办法，就是先把树中序遍历到数组中，然后在数组中找到target的后一个数
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var nums []*TreeNode
	inOrder(root, &nums)

	n := len(nums)
	var i int
	for i < n {
		if nums[i].val == p.val {
			break
		}
	}

	if i+1 < n {
		return nums[i+1]
	} else {
		return nil
	}
}

func inOrder(root *TreeNode, nums *[]*TreeNode) {
	if root == nil {
		return
	}

	inOrder(root.left, nums)
	*nums = append(*nums, root)
	inOrder(root.right, nums)
}

// 中序遍历的方法，使用一个布尔变量表示当前是否已经遍历过target，如果遍历到当前节点时已经遍历过target，那么返回当前节点；如果遍历到当前节点
//没有遍历过target，那么查看当前节点是否是target，如果是，那么将该布尔变量标记为true，继续遍历，下一个节点就是返回节点；如果不是，那么继续遍历
func inorderSuccessor2(root *TreeNode, p *TreeNode) *TreeNode {
	stack := utils.NewStack()
	visited := false
	t := root

	for !stack.Empty() || t != nil {
		for t != nil {
			stack.Push(t)
			t = t.left
		}

		t = stack.Pop().(*TreeNode)
		if visited {
			return t
		}
		if reflect.DeepEqual(t, p) {
			visited = true
		}
		t = t.right
	}

	return nil
}
