package leet_106

/*
	Given inorder and postorder traversal of a tree, construct the binary tree.
	Note:
		You may assume that duplicates do not exist in the tree.
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func helper(postStart int, inStart int, inEnd int, in []int, post []int) *TreeNode {
	var root,rootIndex int
	if postStart < 0 || inStart > inEnd {
		return nil
	}

	root = post[postStart]
	for i := inStart; i<= inEnd ;i++ {
		if in[i] == root {
			rootIndex = i
			break  // avoid meaningless loop
		}
	}
	// inEnd - rootIndex + 1 means the length of right subtree in inorder , postStart - (inEnd - rootIndex + 1) means the root position in postOrder of subtree
	left := helper(postStart-(inEnd-rootIndex+1),inStart,rootIndex-1,in,post)
	right := helper(postStart-1,rootIndex+1,inEnd,in,post)

	return &TreeNode{
		Val:   root,
		Left:  left,
		Right: right,
	}
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	return helper(len(postorder)-1,0,len(inorder)-1,inorder,postorder)
}