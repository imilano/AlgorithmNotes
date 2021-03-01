package leet_173

import "sync"

/*
	Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.
	Calling next() will return the next smallest number in the BST.

Note:
    next() and hasNext() should run in average O(1) time and uses O(h) memory, where h is the height of the tree.
    You may assume that next() call will always be valid, that is, there will be at least a next smallest number in the BST when next() is called.

*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type (
	Stack struct {
		top    *Node
		length int
		lock   *sync.RWMutex
	}

	Node struct {
		value interface{}
		pre   *Node
	}
)

// Create new stack
func NewStack() *Stack {
	return &Stack{
		top:    nil,
		length: 0,
		lock:   &sync.RWMutex{},
	}
}

// Length of stack
func (s *Stack) Len() int {
	return s.length
}

// Top element  of stack
func (s *Stack) Top() interface{} {
	if s.length == 0 {
		return nil
	}

	return s.top.value
}

// Pop pop out top node and return its value
func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.length == 0 {
		return nil
	}

	r := s.top
	s.length--
	s.top = r.pre
	return r.value
}

// Push push a node into stack
func (s *Stack) Push(value interface{}) {
	node := &Node{
		value: value,
		pre:   s.top,
	}

	s.length++
	s.top = node
}

// Empty check if it is an empty stack
func (s *Stack) Empty() bool {
	return s.Len() == 0
}

// --------------------------------------------
// Using Stack, when building iterator, traverse all the way to the leftmost node, then return stack. When call next, just pop the top node,
// return its value, and push its right node all the way to the leftmost.

type BSTIterator struct {
	stack *Stack
}

func (this *BSTIterator) helper(root *TreeNode) {
	var cur *TreeNode
	if root == nil {
		return
	}

	cur = root
	for cur != nil { // go to the leftmost
		this.stack.Push(cur)
		cur = cur.Left
	}
}

func Constructor(root *TreeNode) BSTIterator {
	stack := NewStack()
	ite := BSTIterator{stack: stack}
	ite.helper(root)
	return ite
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	cur := this.stack.Pop().(*TreeNode) // return the stack top element, and from its right node, push all the left node all the way to the leftmost node.
	this.helper(cur.Right)
	return cur.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.stack.Empty()
}
