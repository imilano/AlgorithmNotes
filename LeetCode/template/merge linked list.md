合并两个链表，需要一个dummyHead，新链表的头。还需要一个cur指针，记录当前新链表中的最新节点。
```go
func merge(L1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	cur := new(ListNode)
	head := cur 
	
	for l1 != nil || l2 != nil {
		if l1 != nil {
			// do something
			carry += l1.Val
			l1  = l1.Next
}
	if k2 != nil {
		// do something
		carry += l2.Val
		l2 = l2.Next
}
	cur.Next = &ListNode {
		Val: carry % 10,
		Next: nil
}
	cur = cur.Next
	carry /= 10
}
}
```