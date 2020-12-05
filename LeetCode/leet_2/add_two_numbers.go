package leet_02

import "fmt"

/**
 * You are given two non-empty linked lists representing two non-negative integers.
 * The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 **/

 // Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }

 func getCarryAndSum(a int) (int,int) {
    c := a /10
    return c,a - 10*c
 }

 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
     // 一味的使用边界检查并不会加快程序运行速度
     // if l1 == nil {
     //     return l2
     // }
     // if l2 == nil {
     //     return l1
     // }

     var carry int
     // 使用new的开销还是比较大，不如直接构造
     //cur := new(ListNode)
     cur := &ListNode{}
     head := cur

     for l1 != nil || l2 != nil || carry != 0 {
         if l1 != nil {
             carry += l1.Val
             l1 = l1.Next
         }

         if l2 != nil {
             carry += l2.Val
             l2 = l2.Next
         }

         cur.Next = &ListNode{
             Val:  carry%10,
             Next: nil,
         }
         cur = cur.Next
         carry = carry/10
     }

     return head.Next
 }


func addTwoNumbersOriginal(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }

    if l2 ==  nil {
        return l1
    }

    var sum, carry int
    head := &ListNode{}
    result := head

    for l1 != nil && l2 != nil {
        tmp := l1.Val + l2.Val + carry
        carry, sum = getCarryAndSum(tmp)

        node := &ListNode{
            Val:  sum,
            Next: nil,
        }
        head.Next = node
        head = node

        l1 = l1.Next
        l2 = l2.Next
    }

    for l1 != nil {
        node := &ListNode{
            Val:  l1.Val + carry,
            Next: nil,
        }

        head.Next = node
        head = node
        l1 = l1.Next
        carry = 0
    }

    for ; l2 != nil; {
    	node := &ListNode{
            Val:  l2.Val + carry,
            Next: nil,
        }

        head.Next = node
        head = node
        l2 = l2.Next
        carry = 0
    }

    if carry != 0 && l1 == nil && l2 == nil {
        node := &ListNode{
            Val:  carry,
            Next: nil,
        }
        head.Next = node
    }

    return result.Next
}
 
 func main() {
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

     r := addTwoNumbers(n1,n3)

     for ;r != nil; {
         fmt.Print(r.Val)
         r = r.Next
     }
 }