package leet_99

import "sort"

/*
	You are given the root of a binary search tree (BST), where exactly two nodes of the tree were swapped by mistake.
	Recover the tree without changing its structure.

	Follow up: A solution using O(n) space is pretty straight forward; could you devise a constant space solution?
 */

// Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

 // 采用中序遍历的方式，遍历所有节点，然后保存该节点以及该节点的值。遍历结束后，对节点值进行排序，然后将排序好的值分别赋值给节点即可
func recoverN(root *TreeNode, list *[]*TreeNode, val *[]int) {
    if root == nil {
        return
    }
    recoverN(root.Left,list,val)
    *list = append(*list,root)
    *val = append(*val,root.Val)
    recoverN(root.Right,list,val)
}

func recoverTree(root *TreeNode)  {
    // use inorder and O(n) complexity
    var list []*TreeNode
    var val []int
    recoverN(root,&list,&val)
    sort.Ints(val)
    for i := range list {
        list[i].Val = val[i]
    }
}


// O(1)解法，使用线索二叉树
