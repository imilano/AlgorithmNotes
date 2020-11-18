package leet_105

/*
	Given preorder and inorder traversal of a tree, construct the binary tree.
	For example, given
		preorder = [3,9,20,15,7]
		inorder = [9,3,15,20,7]

	Return the following binary tree:
		3
	   / \
	  9  20
		/  \
	   15   7
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func helper(preIndex int,inStart int, inEnd int, pre []int, in []int) *TreeNode {
	var root,rootIndex int
	if preIndex > len(pre)-1 || inStart > inEnd {
		return nil
	}

	// seek root index in preOrder array
	root = pre[preIndex]
	for i := inStart; i<= inEnd ; i++ {
		if in[i] == root {
			rootIndex = i
			break // avoid meaningless loop
		}
	}

	// left subtree and right subtree
	left := helper(preIndex+1,inStart,rootIndex-1,pre,in)
	right := helper(preIndex + rootIndex-inStart+1,rootIndex+1,inEnd,pre,in)  // rootIndex - inStart means the subtree length of root node in InOrder array, plus 1 means to add length of root node.

	return &TreeNode{
		Val:   root,
		Left:  left,
		Right: right,
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return helper(0,0,len(inorder)-1,preorder,inorder)
}
