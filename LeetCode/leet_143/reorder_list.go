package leet_143

/*
Given a singly linked list L: L0→L1→…→Ln-1→Ln,
reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example 1:

Given 1->2->3->4, reorder it to 1->4->2->3.
Example 2:

Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
*/


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil || head.Next.Next == nil{
        return
    }

    var stack []*ListNode
    cur := head
    for cur != nil {
       stack = append([]*ListNode{cur},stack...)
       cur = cur.Next
   }

   length := (len(stack)-1)/2
    cur = head
    for length > 0 {
        top := stack[0]
        stack = stack[1:]
        next := cur.Next

        cur.Next = top
        top.Next = next
        cur = next
        length--
    }

    if len(stack) != 0 {
        t := stack[0]
        stack = stack[1:]
        t.Next = nil
    }
}

// replace this with stack
func constructList(head *ListNode) (*ListNode,*ListNode) {
    if head.Next == nil {
        return head,head
    }

    first,newHead := constructList(head.Next)
    newHead.Next = head
    ret := newHead.Next
    return first,ret
}
