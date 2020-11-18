package leet_108

/*
	Given an array where elements are sorted in ascending order, convert it to a height balanced BST.
	For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of
	every node never differ by more than 1.
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func helper(nums []int, low int, high int) *TreeNode {
	if low  > high {
		return nil
	}

	mid := low + (high-low)/2
	left := helper(nums,low,mid-1)
	right := helper(nums,mid+1,high)

	return &TreeNode{
		Val:   nums[mid],
		Left:  left,
		Right: right,
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums,0,len(nums)-1)
}
