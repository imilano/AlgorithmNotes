package leet_86

/*
Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

Example:

Input: head = 1->4->3->2->5->2, x = 3
Output: 1->2->2->4->3->5
*/


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

// 空间换时间
// 注意，关键在于，将小于x的数都放在左边，大于x的数都放在右边。大于x的数还需要保证稳定性
func partition(head *ListNode, x int) *ListNode {
	// edge case
	if head == nil || head.Next == nil {
		return head
	}

	// scan list to divide number into two group
	var less,greater []int
	cur := head
	for cur != nil {
		if cur.Val < x {
			less = append(less,cur.Val)
		} else {
			greater = append(greater,cur.Val)
		}

		cur = cur.Next
	}

	cur = head
	for len(less) != 0 {
		cur.Val = less[0]
		cur =cur.Next
		less = less[1:]
	}

	for len(greater) != 0 {
		cur.Val = greater[0]
		greater = greater[1:]
		cur = cur.Next
	}

	return head
}

// 使用指针，相比上一步，可以节约很多空间，并且解法更精妙
func partitionConcise(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return  head
	}

	less,greater := new(ListNode),new(ListNode)
	dummyLess,dummyGreater := less,greater
	for head != nil {
		node := head
		// 使用尾插法
		if node.Val < x {
			less.Next = node
			less = node
		} else {
			greater.Next = node
			greater= node
		}

		head = head.Next
	}

	greater.Next = nil
	less.Next = dummyGreater.Next
	return dummyLess.Next
}
