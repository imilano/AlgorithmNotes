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
	Val   int
	Left  *TreeNode
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
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil && right == nil {
		return nil
	}

	if left != nil {
		return left
	}

	return right
}


// --------------------------
// 与上述相同p和q的分布只会有三种情况，要么是在root的左子树中，要么是在root的右子树中，要么左右子树都存在。分类讨论：
// 	- 如果根节点为空或者根节点等于p或q，那么直接返回根节点；否则对根节点的左右节点分别进行遍历；
//	- 如果如果子树返回非空，右子树也返回非空，那么直接返回根节点
//	- 如果左子树为空，右子树非空，那么返回右子树。（右子树非空包含两种情况，一种情况是返回的结果是p和q二者中的较高者，一种是返回q和p的最小公共祖先）
//	- 如果右子树为空，左子树非空，那么直接返回左子树。（左子树非空包含两种情况，一种是返回的结果是p和q二者中的较高者，一种是返回p和q的最小公共祖先）
func lowestCommonAncestor2(root,p,q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor2(root.Left,p,q)
	right := lowestCommonAncestor2(root.Right,p,q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	} else {
		return right
	}
}
