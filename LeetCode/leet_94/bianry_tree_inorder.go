package leet_94

import "lightsinger.life/algorithm/leetcode/utils"

/*
	Given the root of a binary tree, return the inorder traversal of its nodes' values.
*/

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result = make([]int, 0)

func inorder(root *TreeNode) {
	if root == nil {
		return
	}

	inorder(root.Left)
	result = append(result, root.Val)
	inorder(root.Right)
}

// Inorder with recursion
func inorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)

	if root == nil {
		return arr
	}

	arr = append(inorderTraversal(root.Left), root.Val)
	arr = append(arr, inorderTraversal(root.Right)...)
	return arr
	//inorder(root)
	//return result
}

// Inorder with stack
func inorderWithStack(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := utils.NewStack()

	cur := root
	for cur != nil || stack.Len() != 0 { // or, not and
		// go to left most ,and push every element to stack
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}

		//// access root node's value
		//result = append(result,cur.Val) // WRONG!!!!, nil pointer reference, cause previous for loop would make cur equal to nil
		//stack.Pop()  // delete current node
		//cur = cur.Right // go through right node

		// get top of stack
		cur = stack.Pop().(*TreeNode) // why should i use type assert when pop actually returns a int?
		// access root value
		result = append(result, cur.Val)
		// go through right node
		cur = cur.Right
	}
	return result
}

// Implement stack with go,
// see util/stack
