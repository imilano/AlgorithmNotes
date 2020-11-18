package leet_199

/*
	Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.
	Example:
		Input: [1,2,3,null,5,null,4]
		Output: [1, 3, 4]
		Explanation:
		   1            <---
		 /   \
		2     3         <---
		 \     \
		  5     4       <---
*/


// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
    var res []int
    var cnt int
    var cur *TreeNode
    var que []*TreeNode

    if root == nil {
        return res
    }
    cur = root
    que = append(que,root)
    for len(que) != 0 {
        cnt = len(que)
        for i := 0; i <cnt; i++ {
            cur = que[0]
            que = que[1:]
            if i == 0 {
                res = append(res,cur.Val)
            }

            if cur.Right != nil {
                que = append(que,cur.Right)
            }

            if cur.Left != nil {
                que = append(que, cur.Left)
            }
        }
    }
    return res
}