package leet_235


// Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// time complexity: O(n)
// space complexity: O(n),cause the height of a skewed BST might by n,which lead to n level stack
func lowestCommonAncestorRec(root,p,q *TreeNode) *TreeNode {
    var res *TreeNode
    parentVal,pVal,qVal := root.Val,p.Val,q.Val

    if pVal > parentVal && qVal > parentVal {
        res = lowestCommonAncestor(root.Right,p,q)
    } else if pVal < parentVal && qVal < parentVal{
        res = lowestCommonAncestor(root.Left,p,q)
    } else {
        res = root
        return res
    }
    return res
}

// time complexity: O(n)
// space complexity: O(1)
func lowestCommonAncestorIte(root,p,q *TreeNode) *TreeNode  {
    var res *TreeNode
    pVal,qVal := p.Val,q.Val

    node := root
    for node != nil {  // Traversal the tree
        parentVal := node.Val
        if pVal > parentVal && qVal > parentVal {  // traverse to right
            node = node.Right
        } else if pVal < parentVal && qVal < parentVal {  // traverse to left
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
    return lowestCommonAncestorIte(root,p,q)
}
