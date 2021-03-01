package leet_272

/*
Given a non-empty binary search tree and a target value, find k values in the BST that are closest to the target.

Note:

Given target value is a floating point.
You may assume k is always valid, that is: k≤ total nodes.
You are guaranteed to have only one unique set of k values in the BST that are closest to the target.
Example:

Input: root = [4,2,5,1,3], target = 3.714286, and k = 2

    4
   / \
  2   5
 / \
1   3

Output: [4,3]
Follow up:
Assume that the BST is balanced, could you solve it in less than O(n) runtime (where n = total nodes)?
*/
type TreeNode struct {
	val         int
	left, right *TreeNode
}

// 使用中序遍历，建立一个有序数组，然后从距离target最近的点开始，维持一个窗口，然后从窗口开始不断向两侧扩充
func closestKValues(root *TreeNode, target float32, k int) []int {
	var nums []int

	inOrder(root, &nums)
	closest := findClosest(nums, target)

	// maintain a window boundary
	n := len(nums)
	left, right := closest, closest
	for left > 0 && right < n && right-left+1 < k {
		if abs(float32(nums[left])-target) <= abs(float32(nums[nums[right]])-target) {
			left--
		} else {
			right++
		}
	}

	return nums[left : right+1]
}

func findClosest(nums []int, target float32) int {
	n := len(nums)
	left, right := 0, n

	for left < right {
		mid := left + (right-left)/2
		if float32(nums[mid]) < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if abs(float32(nums[left])-target) < abs(float32(nums[right])-target) {
		return left
	} else {
		return right
	}
}

func abs(a float32) float32 {
	if a < 0 {
		a = -a
	}

	return a
}

func inOrder(root *TreeNode, val *[]int) {
	if root == nil {
		return
	}

	inOrder(root.left, val)
	*val = append(*val, root.val)
	inOrder(root.right, val)
}
