package leet_437

/*
	437. Path Sum III
	You are given a binary tree in which each node contains an integer value.
	Find the number of paths that sum to a given value.
	The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).
	The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.
 */
//  Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

// brutal force
// time complexity: best O(nlogn), worst O(n^2)
// space complexity: O(n)
// 采用双重递归的方法
func pathSumFrom(root *TreeNode, sum int) int {
    var s int
    if root == nil {
        return 0
    }

    sum -= root.Val
    if sum == 0 {
        s = 1
    } else {
        s = 0
    }
    return s + pathSumFrom(root.Left,sum) + pathSumFrom(root.Right,sum)
}

 // contain many repeated calculation
func pathSum(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }

    return pathSumFrom(root,sum) + pathSum(root.Left, sum) + pathSum(root.Right,sum)
}