package leet_255

import (
	"lightsinger.life/algorithm/leetcode/utils"
	"math"
)

/*
Given an array of numbers, verify whether it is the correct preorder traversal sequence of a binary search tree.

You may assume each number in the sequence is unique.

Consider the following binary search tree:

     5
    / \
   2   6
  / \
 1   3
Example 1:

Input: [5,2,6,1,3]
Output: false
Example 2:

Input: [5,2,1,3,6]
Output: true
Follow up:
Could you do it using only constant space complexity?
*/

//前序遍历有一个性质，就是对于当前值val，左子树的值都比val小，而右子树的任何值都比val大。
//使用分治法，对于当前数字val，保存两个数字lower和upper分别代表下界和上界，任何一个节点值val都必须在lower和upper之间，如果不在，则返回false；
//如果在，那么遍历前序数组，找到第一个比当前值val大的值i（也就是当前节点的右节点），以此为分界，对左右两部分分别进行递归调用；需要注意的是，需要将
//左递归的upper更新为当前值val，将右递归的lower值更新为当前值val，只有左右两递归都返回true，才返回true
func verifyPreorder(preOrder []int) bool {
	return helper(preOrder, 0, len(preOrder)-1, math.MaxInt32, math.MinInt32)
}

func helper(nums []int, start, end int, upper, lower int) bool {
	if start > end { // 整个数组遍历完
		return true
	}

	val, i := nums[start], 0 // val表示当前树的跟节点
	if val <= lower || val >= upper {
		return false
	}
	for i = start + 1; i <= end; i++ { // 找到当前节点的右子节点
		if nums[i] >= val {
			break
		}
	}

	// 根节点的左子树都必须小于val，根节点的右子树都必须大于val
	return helper(nums, start+1, i-1, val, lower) && helper(nums, i, end, upper, val)
}

// TODO solve it again
// 采用单调栈的方式，还是根据之前的性质，当前节点的值是左子树的最大值，是右子树的最小值；遍历数组，使用min表示当前子树的根节点（也就是右子树
//的最小值），如果当前遍历到的元素值大于栈顶值，那么意味着一棵子树的所有左节点已经遍历完了，当前遍历到的节点是某个节点的右子树，那么我们需要
//找到当前遍历到的节点是哪个节点的右节点，很明显，如果当前遍历到的节点一直比栈顶元素大，那么栈顶元素仍然是某个节点的左子树的节点，于是我们将
//min设置为栈顶节点，然后将栈顶节点出栈，一直到栈为空或者当前栈顶节点比已经大于当前遍历到的元素了，那么我们最后出栈的元素（也就是min）就是
//当前遍历节点的根节点
func verifyPreorder2(preOrder []int) bool {
	s := utils.NewStack()
	min := math.MinInt32
	for _, v := range preOrder {
		if v < min {
			return false
		}

		for !s.Empty() && s.Top().(int) < v {
			min = s.Pop().(int)
		}
		s.Push(v)
	}

	return true
}
