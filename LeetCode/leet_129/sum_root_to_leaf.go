package leet_129

/*
	Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.
	An example is the root-to-leaf path 1->2->3 which represents the number 123.
	Find the total sum of all root-to-leaf numbers.
	Note: A leaf is a node with no children.
 */


//  Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

 var res []int
 func helper(root *TreeNode, sum int)  {
     if root == nil {
         res = append(res,sum)
         return
     }
     //if root.Left == nil && root.Right == nil {
     //    sum = sum + root.Val
     //    res = append(res,sum)
     //    return
     //}

     sum = sum * 10 + root.Val
     helper(root.Left,sum)
     helper(root.Right,sum)
 }

func sumNumbers(root *TreeNode) int {
    //var sum int
    //if root == nil {
    //   return sum
    //}
    //if root.Left == nil && root.Right ==nil {
    //   return root.Val
    //}
    //
    //helper(root,root.Val)
    //for _,v := range res {
    //   sum += v
    //}
    //
    //return sum

    return getSum(root,0)
}

func getSum(root *TreeNode, sum int) int  {
    if root == nil {
        return 0
    }

    sum = sum * 10 + root.Val
    if root.Left == nil && root.Right == nil {
        return sum
    }

    return getSum(root.Left,sum) + getSum(root.Right,sum)
}