package leet_100

/*
	Given two binary trees, write a function to check if they are the same or not.
	Two binary trees are considered the same if they are structurally identical and the nodes have the same value.
 */


 // Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

 // Recursive way
func isSameTreeRec(p *TreeNode, q *TreeNode) bool {
    if p == nil && q== nil {
        return true
    }

    if q == nil || p== nil {
        return false
    }

    if p.Val != q.Val {
        return false
    }
    return isSameTreeRec(p.Right,q.Right) && isSameTreeRec(p.Left, q.Left)
}

// Iterative. Use a queue, both q and p enque, then deque twice to check if node are same.
func check(p,q *TreeNode) bool {
    if p == nil && q== nil {
        return true
    }

    if q == nil || p== nil {
        return false
    }

    if p.Val != q.Val {
        return false
    }

    return true
}

func isSameTreeIte(p *TreeNode, q *TreeNode) bool {
    var que []*TreeNode

    que = append(que,p,q)
    for len(que) != 0 {
        first,second := que[0],que[1]
        que = que[2:]

        if !check(first,second) {
            return false
        }

        // If first and second are both nil, we can't dereference a nil pointer
        if first != nil && second != nil {
            que = append(que,first.Left,second.Left)
            que = append(que,first.Right,second.Right)
        }

    }

    return true
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    return isSameTreeIte(p,q)
}