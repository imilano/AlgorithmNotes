package leet_20

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
*/


import "sync"

type (
	Stack struct {
		top *Node
		length int
		lock *sync.RWMutex
	}

	Node struct {
		value interface{}
		pre *Node
	}
)

// Create new stack
func NewStack() *Stack{
	return &Stack{
		top:    nil,
		length: 0,
		lock:   &sync.RWMutex{},
	}
}

// Length of stack
func (s *Stack) Len() int{
	return s.length
}

// Top element  of stack
func (s *Stack) Top()  interface{} {
	if s.length ==0 {
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
func (s *Stack) Empty()  bool{
	return s.Len() == 0
}



// wrong idea
func isValidOriginal(s string) bool {
	length := len(s)
	if length == 0 || length%2 == 1{
		return false
	}

	left,right := length/2,length/2+1
	dict := map[uint8]uint8{
		'(':')',
		'[':']',
		'{':'}',
	}
	for left >= 0 && right <= length-1 {
		if s[left] != dict[s[left]] {
			return false
		}
		left--
		right++
	}

	return true
}

func isValid(s string) bool {
	length := len(s)
	if length == 0 || length%2 == 1{
		return false
	}

	stack := NewStack()

	var index int
	dict := map[uint8]uint8{
		'(':')',
		'[':']',
		'{':'}',
	}

	for index != length {
		cur := s[index]
		if cur == '(' || cur == '[' || cur == '{' {
			stack.Push(s[index])
			index++
			continue
		}
		index++

		if !stack.Empty() {
			mat := stack.Pop().(uint8)
			if cur != dict[mat] {
				return false
			}
		} else {
			return false
		}


	}

	return stack.Empty()
}