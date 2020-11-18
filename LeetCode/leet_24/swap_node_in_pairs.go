package leet_24

/*
	Given a linked list, swap every two adjacent nodes and return its head.
	You may not modify the values in the list's nodes. Only nodes itself may be changed.
 */

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    dummy := new(ListNode)
    dummy.Next = head
    current := dummy
    for current.Next != nil && current.Next.Next != nil {
        first := current.Next
        second := current.Next.Next

        current.Next  = second
        first.Next = second.Next
        second.Next = first
        current = current.Next.Next
    }

    return dummy.Next
}
