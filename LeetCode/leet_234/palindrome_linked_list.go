package leet_234

import (
	"lightsinger.life/algorithm/leetcode/utils"
)

/*
Given a singly linked list, determine if it is a palindrome.

Example 1:

Input: 1->2
Output: false
Example 2:

Input: 1->2->2->1
Output: true
Follow up:
Could you do it in O(n) time and O(1) space?
*/


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

//------------------------------------------------------------
// 使用栈
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next ==  nil {
		return true
	}

	stack := utils.NewStack()
	t := head
	for t != nil {
		stack.Push(t.Val)
		t = t.Next
	}

	t = head
	for t != nil {
		v := stack.Pop().(int)
		if v != t.Val {
			return false
		}
		t = t.Next
	}

	return true
}

//--------------------------------------------------------------
// 使用栈和双指针方法。先采用双指针方法，慢指针走一步，快指针走两步，快指针走到末尾，慢指针走到中点，在慢指针走的时候，将遍历到的值入栈。
// 指针遍历完之后，将慢指针元素和栈中的元素进行比较，如果相等则为回文，否则不是

//---------------------------------------------------------------------------
// 先遍历得到元素数量，然后对链表后半部分进行翻转，然后使用双指针方法，从原头部和后半部分的头部进行一一比较
func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil{
		return true
	}

	slow,fast := head,head
	for fast.Next != nil && fast.Next.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}

	newHead := reverseList(slow.Next)
	s,l := head,newHead
	for l != nil {
		if s.Val != l.Val {
			return false
		}
		s = s.Next
		l = l.Next
	}


	return true
}


func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}



