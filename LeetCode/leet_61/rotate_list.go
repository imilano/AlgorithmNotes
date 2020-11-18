package leet_61

/*
Given a linked list, rotate the list to the right by k places, where k is non-negative.
 */



// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    // if list is nil or k equals to 0, then we just need to return headd directly
    if head == nil || k == 0{
        return head
    }

    // count length
    length := 1
    dummy := head
    for dummy.Next != nil {
        length++
        dummy = dummy.Next
    }

    // if k is multiple of length, return head directly
    if k%length == 0 {
        return head
    }

    // if k is larger than length,  then do modular operation.
    // using two pointer to find kth node counting from right
    k %= length
    fast,slow := head,head
    for k > 0 {
        fast = fast.Next
        k--
    }

    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }

    dummy = slow.Next
    slow.Next = nil
    fast.Next = head

    return dummy
}

