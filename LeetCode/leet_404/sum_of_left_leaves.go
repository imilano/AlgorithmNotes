package leet_404


/*
Find the sum of all left leaves in a given binary tree.
Example:

    3
   / \
  9  20
    /  \
   15   7
There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
 */


// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// original wrong solution
func sumOfLeftLeavesRec(root *TreeNode) int {
    var sum,cnt int
    if root == nil || (root.Left == nil && root.Right == nil) {
        return sum
    }

    var que []*TreeNode
    var cur *TreeNode
    que = append(que,root)

    for len(que) != 0{
        cnt = len(que)
        for i := 0 ; i< cnt; i++ {
            cur = que[0]
            que = que[1:]
            if i == 0 && cur.Left == nil && cur.Right == nil {
                sum += cur.Val
            }

            if cur.Left != nil {
                que = append(que, cur.Left)
            }

            if cur.Right != nil {
                que = append(que,cur.Right)
            }
        }
    }

    return sum
}

// right solution using recursive
func helper(root *TreeNode, isLeft bool) int {
    if root == nil {
        return 0
    }

    if root.Left == nil && root.Right == nil && isLeft {
        return root.Val
    }

    return helper(root.Left,true) + helper(root.Right, false)
}
func sumOfLeftLeaves(root *TreeNode) int {
    return helper(root,false)
}
