package leet_160

/*
Write a program to find the node at which the intersection of two singly linked lists begins.
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//var res *ListNode
	var lenA, lenB int

	if headA == nil || headB == nil {
		return nil
	}
	cur := headA
	for cur != nil {
		lenA++
		cur = cur.Next
	}

	cur = headB
	for cur != nil {
		lenB++
		cur = cur.Next
	}

	startA, startB := headA, headB
	if lenA > lenB {
		i := 0
		for i < lenA-lenB {
			startA = startA.Next
			i++
		}
	} else {
		i := 0
		for i < lenB-lenA {
			startB = startB.Next
			i++
		}
	}

	for startA != nil && startB != nil && startA != startB {
		startA = startA.Next
		startB = startB.Next
		//res = startB
	}

	// 可能不会交织，故而会出现二者都为空的情况，所以需要判断startA的状态
	if startB != nil && startA != nil {
		return startA
	}
	return nil
}

// -----------------------------
// 使用环的思路
// a 第一次遍历完，我们跳到B的开头进行遍历；同理，b第一次遍历完，我们跳到A的开头进行遍历，最后A和B必然会相遇，因为二者都走过了相同的路程.
// 相遇时，要么是在交叉节点，要么是在各自的尾部空节点，故而直接返回即可。
func getIntersectionNodeUseCycle(headA, headB *ListNode) *ListNode {
	if headB == nil || headA == nil {
		return nil
	}

	a, b := headA, headB
	for a != b {
		//a = a ? a.Next : headB
		//b = b ? b.Next : headA
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}

		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}

	return a
}
