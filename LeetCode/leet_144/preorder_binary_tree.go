package leet_144

import (
    "lightsinger.life/algorithm/leetcode/utils"
)

/*
	Given the root of a binary tree, return the preorder traversal of its nodes' values.

	Follow up:
		Recursive solution is trivial, could you do it iteratively?
 */


 // Definition for a binary tree node.
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

func preWithRec(root *TreeNode) []int {
    s := []int{}
    if root == nil {
        return s
    }

    s = append([]int{root.Val},preWithRec(root.Left)...)
    s = append(s,preWithRec(root.Right)...)

    return s
}


// preorder traversal with stack
func  preWithStack(root *TreeNode) []int {
    result := make([]int,0)
    stack := utils.NewStack()

    cur := root
    for cur != nil || !stack.Empty() {
        result = append(result,cur.Val)

        if cur.Right != nil {
            stack.Push(cur.Right)
        }

        if cur.Left != nil {
            stack.Push(cur.Left)
        }

        if !stack.Empty() {
            cur = stack.Pop().(*TreeNode)
        } else {
            cur = nil
        }
    }

    return result
}



func preorderTraversal(root *TreeNode) []int {
    // recursive
    //return preWithRec(root)

    // stack
    return preWithStack(root)
}
