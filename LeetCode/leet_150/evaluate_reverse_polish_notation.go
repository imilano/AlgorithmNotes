package leet_150

/*
Evaluate the value of an arithmetic expression in Reverse Polish Notation.

Valid operators are +, -, *, /. Each operand may be an integer or another expression.

Note:

Division between two integers should truncate toward zero.
The given RPN expression is always valid. That means the expression would always evaluate to a result and there won't be any divide by zero operation.
Example 1:

Input: ["2", "1", "+", "3", "*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9
*/

//------------------------------------------
// utils
import (
	"strconv"
	"sync"
)

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

//-------------------------------------
func evalRPN(tokens []string) int {
	stack := NewStack()

	for len(tokens) != 0 {
		cur := tokens[0]
		tokens = tokens[1:]

		token, err := strconv.Atoi(cur)
		if err == nil {
			stack.Push(token)
		} else {
			num1 := stack.Top().(int)
			stack.Pop()
			num2 := stack.Top().(int)
			stack.Pop()

			// be careful about the sequence of number
			if cur == "+" {
				stack.Push(num2 + num1)
			} else if cur == "-" {
				stack.Push(num2 - num1)
			} else if cur == "*" {
				stack.Push(num2 * num1)
			} else if cur == "/" {
				stack.Push(num2 / num1)
			}
		}
	}

	return stack.Top().(int)
}
