package leet_145

import (
    "lightsinger.life/algorithm/leetcode/utils"
)

/*
	Given the root of a binary tree, return the postorder traversal of its nodes' values.
	Follow up:
		Recursive solution is trivial, could you do it iteratively?


    方法1：使用前序遍历。
            前序遍历的顺序是「根-左-右」而后序遍历的顺序是『左-右-根』，我们只需要修改一下前序遍历的顺序，将其改为「根-右-左」，再对返回的结果做一个reverse即可得后序遍历的结果。
    方法2：使用栈
            同样采用前序遍历的方式，但是在插入的时候选择头插法而不是尾插法

 */


//  Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func postWithRec(root *TreeNode) []int {
    var res []int

    if root == nil {
        return res
    }

    res = append(postWithRec(root.Left),postWithRec(root.Right)...)
    res = append(res,root.Val)

    return res
}

func postWithStack(root *TreeNode) []int {
    var result []int
    if root == nil {
        return result
    }

    stack := utils.NewStack()
    stack.Push(root)
    for !stack.Empty() {  // 按照「根-右-左」的顺序遍历，但是在入栈时，由于后遍历左节点，所以需要先将左节点入栈
        cur := stack.Pop().(*TreeNode)
        result =  append([]int{cur.Val},result...)  // 逆序头插

        if cur.Left != nil {
            stack.Push(cur.Left)
        }

        if cur.Right != nil {
            stack.Push(cur.Right)
        }
    }

    return result
}


 // Modify preorder result to get postorder result
func postUsePre(root *TreeNode) []int {
    var result []int

    if root == nil {
        return result
    }

    stack := utils.NewStack()
    stack.Push(root)
    if !stack.Empty() {
        cur := stack.Pop().(*TreeNode)

        result =  append(result,cur.Val)
        if cur.Left != nil {
            stack.Push(cur.Left)
        }

        if cur.Right != nil {
            stack.Push(cur.Right)
        }
    }

    return reverse(result)
}

 func reverse(arr []int) []int {
    if len(arr) <=1 {
        return arr
    }

    left,right := 0,len(arr)-1
    for left < right{
        arr[left],arr[right] = arr[right],arr[left]

        left++;
        right--;
     }

     return arr
 }

