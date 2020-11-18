package leet_236

/*
	Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
	According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

	Constraints:
			The number of nodes in the tree is in the range [2, 105].
			-109 <= Node.val <= 109
			All Node.val are unique.
			p != q
			p and q will exist in the tree.
*/

// Definition for TreeNode.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// add myself
type ListNode struct {

}

// https://www.youtube.com/watch?v=13m9ZCB8gjw&feature=emb_logo
// If root == nil, then we just return nil
// If root == either q or p, then we just return root.
// Then every node on the tree has two value returned from its left and right, we check the return value left and right,
// if left and right are both nil, which we do not find the ancestor, then we just return null, returning to its parent;
// if left and right both are not nil, then we just find the two node, root node will be the lca, we just need to return root;
// else either left or right is nil, if left is not nil, we return left, else ,we return right.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil  {
        return nil
    }

    if root.Val == p.Val || root.Val == q.Val {
        return root
    }

    left := lowestCommonAncestor(root.Left,p,q)
    right := lowestCommonAncestor(root.Right,p,q)

    if left != nil && right != nil {
        return root
    }

    if left == nil && right == nil {
        return nil
    }

    if left != nil  {
        return left
    }

    return right
}