package leet_19


/*
	Given the head of a linked list, remove the nth node from the end of the list and return its head.
	Follow up: Could you do this in one pass?
 */


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

// time complexity: O(1)
// space complexity: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var fast,slow *ListNode
    if head == nil {
        return nil
    }


    fast,slow = head,head
    for i := 0; i< n;i++ {
        fast = fast.Next
    }
    if fast == nil {
        return head.Next
    }

    for fast.Next != nil {
        fast=fast.Next
        slow = slow.Next
    }

    slow.Next = slow.Next.Next
    return head
}

// time complexity: O(n)
// space complexity: O(n)
func helperRec(root *ListNode, depth *int,n int) *ListNode {
    if root == nil {
        return nil
    }

    child := helperRec(root.Next,depth,n)
    *depth++
    if *depth == n {
        return child
    }

    root.Next = child
    return root
}


func removeNthFromEndRec(head *ListNode,n int) *ListNode {
    if head == nil {
        return nil
    }

    var depth int
    return helperRec(head,&depth,n)
}