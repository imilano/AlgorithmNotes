package leet_234

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l := reverseList(l1)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}

	l2 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  0,
			Next: nil,
		},
	}

	fmt.Println(isPalindrome2(l2))

}
