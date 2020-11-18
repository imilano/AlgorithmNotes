package leet_82

/*
Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.
Return the linked list sorted as well.

Example 1:
	Input: 1->2->3->3->4->4->5
	Output: 1->2->5

Example 2:
	Input: 1->1->1->2->3
	Output: 2->3
*/


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

// 注意题目，一旦一个值出现了重复，那么就要将所有的重复节点都删除。

// ----------------------------------------------------
// Recursive way
//
func deleteDuplicatesRec(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    if head.Val != head.Next.Val {
        head.Next = deleteDuplicatesRec(head.Next)
        return head
    } else {
        for head.Next != nil && head.Val == head.Next.Val {
            head = head.Next
        }

        return deleteDuplicatesRec(head.Next)
    }
}
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    //first,second := head,head.Next
    //for second != nil {
    //    if first.Val == second.Val {
    //        second = second.Next
    //    } else {
    //        first.Next = second
    //        first = second
    //        second = second.Next
    //    }
    //}

    cur := head
    var que []*ListNode
    que = append(que,cur)
    for cur != nil && len(que) != 0  {
        cur = cur.Next
        tail := que[len(que)-1]
        if cur.Val == tail.Val {
            tmp := cur
            for tmp != nil && tmp.Val == cur.Val {
                tmp = tmp.Next
            }

            cur = tmp
            que = que[:len(que)-1]
        }

        if cur != nil {
            que = append(que, cur)
        }
    }

    if len(que) == 0 {
        return nil
    }

    for i := 0; i< len(que)-1;i++ {
        que[i].Next = que[i+1].Next
    }
    que[len(que)-1].Next = nil
    return que[0]
}