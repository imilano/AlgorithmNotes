package leet_147

import "sort"

/*
Sort a linked list using insertion sort.
 */
// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	var nums []int
	cur := head
	for cur != nil {
		nums = append(nums,cur.Val)
		cur = cur.Next
	}

	sort.Ints(nums)
	cur = head
	for cur != nil {
		cur.Val = nums[0]
		nums = nums[1:]
		cur = cur.Next
	}

	return head
}
