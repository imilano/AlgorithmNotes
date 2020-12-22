package leet_23

/*
	You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
	Merge all the linked-lists into one sorted linked-list and return it.
 */


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

//---------------------
// 方法1：复用mergeTwoSortedList，逐一合并，效率较低
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

        dummy = dummy.Next
    }

    if l1 != nil {
        dummy.Next = l1
    } else if l2 != nil {
        dummy.Next = l2
    }

    return head
}


func mergeKLists(lists []*ListNode) *ListNode {
    length := len(lists)
    if length == 0 {
        return nil
    }

    var head *ListNode
    for  _,v := range lists {
        head = mergeTwoLists(head,v)
    }

    return head
}


//-----------------------------------------------
// 方法2：使用分治法
// time complexity: O(Nlgk), N 是所有的元素数量，k是数组长度
// space complexity: O(1)
func mergeKListsDivide(lists []*ListNode) *ListNode {
    if lists == nil {
        return nil
    }

    n := len(lists)
    for n > 1{
        // 加1是为了处理奇数个数的情况，加一保证仍然能从尾部开始向前进行合并
        // 如果个数是偶数，那么加1无影响
        k := (n+1)/2
        for i := 0; i< n/2;i++ {
            lists[i] = mergeTwoLists(lists[i],lists[i+k])
        }
        n = k
    }

    return lists[0]
}

//--------------------------------
// 方法3：use min-heap
