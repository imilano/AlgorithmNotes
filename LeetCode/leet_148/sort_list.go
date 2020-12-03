package leet_148

import "sort"

/*
Given the head of a linked list, return the list after sorting it in ascending order.

Follow up: Can you sort the linked list in O(n logn) time and O(1) memory (i.e. constant space)?
*/



// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

//---------------------------------
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    var nums []int
    cur := head
    for cur != nil {
        nums = append(nums,cur.Val)
        cur = cur.Next
    }

    sort.Ints(nums)
    cur = head
    for cur != nil {
        if len(nums) != 0 {
            cur.Val = nums[0]
            nums = nums[1:]
        }
        cur = cur.Next
    }

    return head
}

//-------------------------------------
// Merge sort
func sortList2(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    pre,slow,fast := head,head,head
    for fast != nil && fast.Next != nil {
        pre = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    pre.Next = nil
    return merge(sortList2(head),sortList2(slow))
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }

    if l2 == nil {
        return l1
    }

    if l1.Val < l2.Val {
        l1.Next = merge(l1.Next,l2)
        return l1
    } else {
        l2.Next = merge(l1,l2.Next)
        return l2
    }
}