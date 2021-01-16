package leet_83

/*
Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:

Input: 1->1->2
Output: 1->2

Example 2:

Input: 1->1->2->3->3
Output: 1->2->3
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Recursive way
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val != head.Next.Val {
		head.Next = deleteDuplicates(head.Next)
		return head
	} else {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}

		return deleteDuplicates(head)
	}
}

// Iterative way
func deleteDuplicatesIte(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	dummy := new(ListNode)
	res := dummy
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}

		dummy.Next = cur
		cur = cur.Next
		dummy = dummy.Next
	}

	return res.Next
}
