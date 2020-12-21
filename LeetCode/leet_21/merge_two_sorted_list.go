package leet_21

/*
    Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.
 */

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

//------------------------------
// iterative
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var head,dummy *ListNode

    if l1 == nil {
        return l2
    }

    if l2 == nil {
        return l1
    }

    if l1.Val <= l2.Val {
        head = l1
        l1 = l1.Next
    } else {
        head = l2
        l2 = l2.Next
    }

    dummy = head
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            dummy.Next = l1
            l1 = l1.Next
        } else {
            dummy.Next = l2
            l2 = l2.Next
        }

        // don't forget this
        dummy = dummy.Next
    }

    if l1 != nil {
        dummy.Next = l1
    } else if l2 != nil {
        dummy.Next = l2
    }

    return head
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }

    cur := new(ListNode)
    head := cur
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            cur.Next = l1
            cur = cur.Next
            l1 = l1.Next
        } else {
            cur.Next = l2
            cur = cur.Next
            l2 = l2.Next
        }
    }

    if l1 != nil {
        cur.Next = l1
    }

    if l2 != nil {
        cur.Next = l2
    }

    return head.Next
}


//---------------------------
// recursive
func mergeTwoLists3(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }

    if l2 == nil {
        return l1
    }

    if l1.Val < l2.Val {
        l1.Next = mergeTwoLists3(l1.Next,l2)
        return l1
    } else {
        l2.Next = mergeTwoLists3(l1,l2.Next)
        return l2
    }
}
