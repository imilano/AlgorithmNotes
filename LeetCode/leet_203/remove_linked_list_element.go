package leet_203

/*
Remove all elements from a linked list of integers that have value val.

Example:

Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := new(ListNode)
	pre := dummy
	dummy.Next = head
	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}

	return dummy.Next
}

func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	next := removeElements2(head.Next, val)
	if head.Val == val {
		return next
	} else {
		head.Next = next
		return head
	}
}
