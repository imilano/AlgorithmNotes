package leet_237

/*
Write a function to delete a node in a singly-linked list. You will not be given access to the head of the list, instead you will be given access to the node to be deleted directly.

It is guaranteed that the node to be deleted is not a tail node in the list.
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 很简单，只需要将待删除节点的后续节点的值都往前移动，删除最后一个节点即可
func deleteNode(node *ListNode) {
	cur, next := node, node.Next
	for next.Next != nil {
		cur.Val = next.Val
		cur = cur.Next
		next = next.Next
	}

	cur.Val = next.Val
	cur.Next = nil
}
