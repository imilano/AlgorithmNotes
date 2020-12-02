package leet_143

import "testing"

func TestReorderList(t *testing.T) {
	n3 := &ListNode{Val:3, Next:nil}
	n2 := &ListNode{Val:2, Next:n3}
	n1 := &ListNode{Val:1,Next:n2}
	reorderList(n1)
}