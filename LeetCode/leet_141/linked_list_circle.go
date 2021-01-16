package leet_141

/*
Given `head`, the head of a linked list, determine if the linked list has a cycle in it. There is a cycle in a linked list
if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is
used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
Return true if there is a cycle in the linked list. Otherwise, return false.
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	// WARNING ,起点应该是fast=slow=head
	slow, fast := head, head.Next
	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	return false
}
