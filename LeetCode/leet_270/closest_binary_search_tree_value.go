package leet_270

import (
	"math"
)

/*
Given a non-empty binary search tree and a target value, find the value in the BST that is closest to the target.

Note:

Given target value is a floating point.
You are guaranteed to have only one unique value in the BST that is closest to the target.
Example:

Input: root = [4,2,5,1,3], target = 3.714286

    4
   / \
  2   5
 / \
1   3

Output: 4
*/

type TreeNode struct {
	val int
	left,right *TreeNode
}

func closestValue(root *TreeNode, target float32) int {
	lower,upper := math.MinInt32,math.MaxInt32
	for root != nil {
		if float32(root.val) >= target {
			upper = min(upper,root.val)
			root = root.left
		} else {
			lower = max(lower,root.val)
			root = root.right
		}
	}

	if abs(float32(upper)-target) <= abs(float32(lower)-target) {
		return upper
	} else {
		return lower
	}
}

func abs(a float32) float32 {
	if a < 0 {
		a = -a
	}

	return a
}

func max(a,b int) int {
	if a < b {
		return  b
	}

	return a
}

func min(a,b int) int {
	if a < b {
		return a
	}

	return  b
}
