package inorder

/*
	Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.
	Calling next() will return the next smallest number in the BST.
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	Sli []int
}

func helper(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	res = append(helper(root.Left), root.Val)
	res = append(res, helper(root.Right)...)

	return res
}

func Constructor(root *TreeNode) BSTIterator {
	res := helper(root)
	ite := BSTIterator{Sli: res}
	return ite
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	res := this.Sli[0]
	this.Sli = this.Sli[1:]

	return res
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if len(this.Sli) == 0 {
		return false
	}

	return true
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
