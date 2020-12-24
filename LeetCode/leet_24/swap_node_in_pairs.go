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
    cur := dummy
    for cur.Next != nil && cur.Next.Next != nil {
        first := cur.Next
        second := cur.Next.Next

        cur.Next  = second
        first.Next = second.Next
        second.Next = first
        cur = cur.Next.Next
    }

    return dummy.Next
}

//----------------------
func swapPairs2(head *ListNode) *ListNode {
    var arr []int
    cur := head
    for cur != nil {
        arr = append(arr, cur.Val)
        cur = cur.Next
    }

    for i := 0; i < len(arr)-1; i+=2 {
        arr[i],arr[i+1] = arr[i+1],arr[i]
    }

    cur = head
    for _,v := range arr {
        cur.Val = v
        cur = cur.Next
    }


    return head
}

//-----------------
func swapPairs3(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    cur := new(ListNode)
    dummy := cur
    first,second := head,head.Next
    for first != nil && second != nil{
        t := second.Next

        cur.Next = second
        cur.Next.Next = first
        cur = cur.Next.Next

        first = t
        if t != nil {
            second = t.Next
        } else {
            second = nil
        }
    }

    // deal with odd length
    cur.Next = first
    return dummy.Next
}
