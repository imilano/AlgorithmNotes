package leet_450

/*
	Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.
	Basically, the deletion can be divided into two stages:
    	Search for a node to remove.
    	If the node is found, delete the node.
	Follow up: Can you solve it with time complexity O(height of tree)?
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BST 中删除一个节点，那么就用它左子树的最右节点或者右子树的最左节点来代替这个节点
// 对于BST删除节点，可以把问题分解为以下几种情形（如果节点存在）：
//  - 待删除节点是叶节点，则直接删除该叶子节点
//  - 待删除节点左节点为空，则直接将父节点接到该节点的右子树
//  - 待删除节点右节点为空，则直接将父节点接到该节点的左子树
//  - 待删除节点左子树右子树均不为空，则将该节点的值设置为右子树中的最左节点的值，然后在右子树中递归删除该最左节点即可
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		// if left node is null, then just return its right node
		if root.Left == nil {
			return root.Right
		}

		// if right node is null, then just return its left node
		if root.Right == nil {
			return root.Left
		}

		// if both right node and left node is not null, then replace the value with min value of right subtree,
		// then delete min value of right tree iteratively
		min := root.Right
		for min.Left != nil {
			min = min.Left
		}

		root.Val = min.Val
		root.Right = deleteNode(root.Right, min.Val)
	}

	return root
}
