package leet_02

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	n1 := &ListNode{
		Val:  1,
		Next: nil,
	}

	n2 := &ListNode{
		Val:  8,
		Next: nil,
	}
	n3 := &ListNode{
		Val:  0,
		Next: nil,
	}

	n1.Next = n2

	r := addTwoNumbers(n1, n3)

	for r != nil {
		fmt.Print(r.Val)
		r = r.Next
	}
}

func TestNillPointer(t *testing.T) {
	var l1 *ListNode

	// default value of pointer in go is nil, deference a nil pointer wil lead to panic
	fmt.Println(l1.Val)
	fmt.Println(l1.Next)
}
