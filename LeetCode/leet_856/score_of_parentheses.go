package leet_856

import "sync"

/*
Given a balanced parentheses string `S`, compute the score of the string based on the following rule:
() has score 1
AB has score A + B, where A and B are balanced parentheses strings.
(A) has score 2 * A, where A is a balanced parentheses string.
Example 1:

Input: "()"
Output: 1
Example 2:

Input: "(())"
Output: 2
Example 3:

Input: "()()"
Output: 2
Example 4:

Input: "(()(()))"
Output: 6
Note:

S is a balanced parentheses string, containing only ( and ).
2 <= S.length <= 50
*/

// 一共包含两种关系，包含和并列。一个简单的()值一份，包含的话乘以2，并列的话相加

//--------------------------------
// util function
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

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

//---------------------------------
// recursive solution

func scoreOfParentheses(S string) int {
	var res int
	length := len(S)

	for i := 0; i < length; i++ {
		if S[i] == ')' {
			continue
		}

		// 找到内括号的合法括号范围,只有当cnt为0时，才找到一串合法的内（）的范围
		pos, cnt := i+1, 1
		for cnt != 0 {
			if S[pos] == '(' {
				cnt++
			} else {
				cnt--
			}

			pos++
		}

		innerScore := scoreOfParentheses(S[i+1 : pos])
		res += max(1, 2*innerScore) // 对于()这样的单独的完整括号而言，递归得分为0，而最少得分应该是1，故而应该至少为1
		i = pos - 1
	}

	return res
}

//------------------------------
// iterative using stack
func scoreOfParentheses2(S string) int {
	var res int
	stack := NewStack()
	length := len(S)
	for i := 0; i < length; i++ {
		if S[i] == '(' {
			stack.Push(res)
			res = 0
		} else {
			res = stack.Top().(int) + max(res*2, 1) // top是并列的括号的分数值，而top是当前子字符串的值
			stack.Pop()
		}
	}

	return res
}
