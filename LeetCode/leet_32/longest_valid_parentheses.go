package leet_32

import "math"
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

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}


/*
	Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.
*/

//-------------------------------
// For test
func LongestValidParentheses(s string) int {
	return longestValidParentheses(s)
}



//-------------------------------------
// brutal force
// time complexity: O(n^2)
// space complexity: O(1)

// The workflow of the solution is as below.
//
//    Scan the string from beginning to end.
//    If current character is '(',
//    push its index to the stack. If current character is ')' and the
//    character at the index of the top of stack is '(', we just find a
//    matching pair so pop from the stack. Otherwise, we push the index of
//    ')' to the stack.
//    After the scan is done, the stack will only
//    contain the indices of characters which cannot be matched. Then
//    let's use the opposite side - substring between adjacent indices
//    should be valid parentheses.
//    If the stack is empty, the whole input
//    string is valid. Otherwise, we can scan the stack to get longest
//    valid substring as described in step 3.
func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}

	length := len(s)
	var balanced int
	res := math.MinInt8
	for i := 0; i <length; i++ {
		balanced = 0
		for j := i; j < length;j++ {
			if s[j] == '(' {
				balanced++
			} else if s[j] == ')'{
				balanced--
			}

			if balanced < 0 {
				break
			}

			if balanced == 0 {
				res = max(j-i+1,res)
			}

		}
	}

	if res < 0 {
		return 0
	}
	return res
}

// using stack
// time complexity: O(n)
// space complexity: O(n)
func longestValidParenthesesWithStack(s string) int {
	var ans int
	if len(s) == 0 {
		return ans
	}

	stack := NewStack()
	length := len(s)
	// This is very tricky
	stack.Push(-1)
	for i := 0; i< length;i++ {
		if s[i] == '(' {
			stack.Push(i)
		} else {
			stack.Pop()
			if stack.Empty() {
				stack.Push(i)
			} else {
				ans = max(ans,i-stack.Top().(int))
			}
		}
	}

	return ans
}

// use dynamic programming
// time complexity: O(n)
// space complexity: O(n)


// The main idea is as follows: construct a array longest[], for any longest[i], it stores the longest length of valid
// parentheses which is end at i.

// And the DP idea is :
// 	If s[i] is '(', set longest[i] to 0,because any string end with '(' cannot be a valid one.
//	Else if s[i] is ')'
//	If s[i-1] is '(', longest[i] = longest[i-2] + 2
// 	Else if s[i-1] is ')' and s[i-longest[i-1]-1] == '(', longest[i] = longest[i-1] + 2 + longest[i-longest[i-1]-2]
// 	For example, input "()(())", at i = 5, longest array is [0,2,0,0,2,0], longest[5] = longest[4] + 2 + longest[1] = 6.
func longestValidParenthesesDP(s string) int {
	var res int
	length := len(s)
	if length  == 0 {
		return res
	}
	dp := make([]int,length)
	for i := 1; i< length; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i-2 > 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}

				res = max(res, dp[i])
			} else if s[i-1] == ')' && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
					if i - dp[i-1]-2 > 0 {
						dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
					} else {
						dp[i] = dp[i-1] + 2
					}

					res = max(res,dp[i])
			}
		}
	}

	return res
}

// Another version of dp
func longestValidParenthesesNewDP(s string) int {
	var res int
	length := len(s)
	if length == 0 {
		return res
	}

	dp := make([]int,length)
	open := 0
	for i := 0; i< length; i++ {
		if s[i] == '(' {
			open++
		}

		if s[i] == ')' && open > 0 {
			dp[i] = dp[i-1] + 2

			if i - dp[i] > 0 {
				dp[i] += dp[i-dp[i]]
			}

			open--
		}

		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}