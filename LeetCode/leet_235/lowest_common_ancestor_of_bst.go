package leet_235

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time complexity: O(n)
// space complexity: O(n),cause the height of a skewed BST might by n,which lead to n level stack
// BST 的一个特点就是，大小顺序按照左根右的顺序排列。那么，对于当前根节点root，如果该节点的值大于p和q中的较大者，那么说明p和q的最小公共祖先在跟节点的左子树，那么进入左子树查找；
// 反之，如果根节点的值小于q和p中的较小者，那么说明二者的最下公共祖先在右节点，那么进入右节点进行查找；如果都不是，说明根节点夹在左右节点之间，此处即为我们所需要的节点。
func lowestCommonAncestorRec(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode
	parentVal, pVal, qVal := root.Val, p.Val, q.Val

	if pVal > parentVal && qVal > parentVal {
		res = lowestCommonAncestor(root.Right, p, q)
	} else if pVal < parentVal && qVal < parentVal {
		res = lowestCommonAncestor(root.Left, p, q)
	} else {
		res = root
		return res
	}
	return res
}

// time complexity: O(n)
// space complexity: O(1)
func lowestCommonAncestorIte(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode
	pVal, qVal := p.Val, q.Val

	node := root
	for node != nil { // Traversal the tree
		parentVal := node.Val
		if pVal > parentVal && qVal > parentVal { // traverse to right
			node = node.Right
		} else if pVal < parentVal && qVal < parentVal { // traverse to left
			node = node.Left
		} else {
			res = node
		}
	}

	return res
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// Recursive way
	// return lowestCommonAncestorRec(root,p,q)
	return lowestCommonAncestorIte(root, p, q)
}
