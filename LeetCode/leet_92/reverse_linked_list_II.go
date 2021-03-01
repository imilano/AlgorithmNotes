package leet_92

/*
Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := new(ListNode)
	pre := dummy
	dummy.Next = head
	var cur *ListNode

	// 走到需交换节点的前一个节点
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}

	cur = pre.Next
	// 以转置1、 2、3、4、5中的2、3、4为例
	for i := 0; i < n-m; i++ {
		// 将节点3放在节点1后面
		tmp := pre.Next
		pre.Next = cur.Next
		// 断开节点2和节点3，将节点2的next连到节点4
		cur.Next = cur.Next.Next
		// 将节点1和节点2断开，将节点1连到节点3
		pre.Next.Next = tmp
	}

	return dummy.Next
}

//count := 1
//var stack []*ListNode
//cur := head
//var slow *ListNode
//for count != n+1 {
//    if count == m-1 {
//        slow = cur
//    }
//    if count >= m && count <= n {
//        stack = append([]*ListNode{cur},stack...)
//    }
//
//    cur = cur.Next
//}
//
//for len(stack) != 0 {
//    node := stack[0]
//    stack = stack[1:]
//    slow.Next = node
//    slow = node
//}
//
//slow.Next = cur
//
////slow,fast := head,head
////count := 1
////var dummyHead,dummyTail *ListNode
////for count <= n {
////    if count == m -1 {
////        slow = fast
////    }
////    if count == m {
////        dummyHead = fast
////        dummyTail = fast
////    }
////
////    if count >= m && count <= n {
////        tmp := fast
////        tmp.Next = dummyHead
////        dummyHead = tmp
////    }
////
////    if count > n {
////        break
////    }
////
////    fast = fast.Next
////    count++
////}
////
////slow.Next = dummyHead
////dummyTail.Next = fast
//
//return head
