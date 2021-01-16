package utils

import "sync"

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
