package leet_142

/*
Given a linked list, return the node where the cycle begins. If there is no cycle, return null. There is a cycle in a linked
list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos
is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Notice that you should not modify the linked list.
*/


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

//假设链表起点为X，环与链第一个相交的节点为Y，快慢指针第一次相遇的节点为Z，X与Y之间举例a，Y与Z之间距离b，Z与Y（顺时针方向）距离为c。快指针一次走两步，慢指针走一步。
//快慢指针第一次相遇于Z点，慢指针走过a+b，快指针走过a+b+c+b，则有公式2(a+b) = a+b+c+b，得出a=c。故而让一个指针从Z点开始走，一个指针从X点开始走，第一次相遇的点就是环的入口。
//详见：https://www.cnblogs.com/hiddenfox/p/3408931.html
func detectCycle(head *ListNode) *ListNode {
	// WARNING, 快慢指针应该是从同一个其实节点开始走的，故而是fast = slow=head，而不是slow = head，fast = head.head
    slow,fast := head,head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if slow == fast {
            break
        }
    }

    if fast == nil || fast.Next == nil {
        return nil
    }

    slow = head
    for fast != slow {
        fast = fast.Next
        slow = slow.Next
    }

    return slow
}
