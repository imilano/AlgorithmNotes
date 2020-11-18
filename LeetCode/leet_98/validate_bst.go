package leet_98

import "lightsinger.life/algorithm/leetcode/utils"
/*
    BST的一大特性在于根节点的的值大于左节点的值，同时小于右节点的值。如果只是单纯检查每个节点的值，然后用它和左右节点分别进行比较的话，对于一部分树是没有什么问题的。
    但是，需要注意的是，仅仅如此还不够。BST要求每个节点都要有上述的性质，所以也就意味着根节点的值不仅需要大于左节点的值，而且要比左子树的最大值都大，但是要比右节点的最小值要小。
    对于右节点同理，这就需要我们维持两个指针，一个指向该节点的上界，一个指向该节点的下界。每个节点需要跟该节点的上下界进行对比，一旦出现不在界限内的情况，即可返回false。

    根据以上分析，这里有三种解法：
        - 基于递归
            即一般的递归解法。
        - 基于栈
            本质上是一种迭代方式。
        - 基于中序遍历
            这也是一种基于栈的解法，但是使用了中序遍历的性质。对于一棵BST，中序遍历得到的结果必然是有序序列。我们没有必要等遍历完了再进行比较，只需要一个跟踪上一个遍历节点的指针，当前遍历节点的值与该节点
            所指向的值做对比即可。
 */

// Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func isValidWithInorder(root *TreeNode) bool {
    var pre *TreeNode
    if root == nil {
        return true
    }

    stack := utils.NewStack()
    cur := root
    for cur != nil || !stack.Empty() {
        for cur != nil {
            stack.Push(cur)
            cur = cur.Left
        }

        cur = stack.Pop().(*TreeNode)
        if pre != nil && pre.Val >= cur.Val {
            return false
        }

        pre,cur = cur,cur.Right
    }

    return true
}

func isValidRec(node *TreeNode,min *TreeNode,max *TreeNode) bool {
    if node == nil {
        return true
    }

    // 如果当前节点大于上界
    if max != nil && node.Val >= max.Val {
        return false
    }

    // 如果当前节点小于下界
    if min != nil && node.Val <= min.Val {
        return false
    }

    // 如果当前节点的左节点大于上界，小于下界
    if !isValidRec(node.Left,min,node) {
        return false
    }

    // 如果当前节点的右节点大于上界，小于下界
    if !isValidRec(node.Right,node,max) {
        return false
    }

    return true
}



func isValidBST(root *TreeNode) bool {
    // Inorder traversal
    return isValidWithInorder(root)

    // Recursive way
    // return isValidRec(root,nil,nil)
}

