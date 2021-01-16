package leet_768

import "sync"

/*
Given an array arr of integers (not necessarily distinct), we split the array into some number of "chunks" (partitions), and individually sort each chunk.  After concatenating them, the result equals the sorted array.

What is the most number of chunks we could have made?

Example 1:

Input: arr = [5,4,3,2,1]
Output: 1
Explanation:
Splitting into two or more chunks will not return the required result.
For example, splitting into [5, 4], [3, 2, 1] will result in [4, 5, 1, 2, 3], which isn't sorted.

Example 2:

Input: arr = [2,1,3,4,4]
Output: 4
Explanation:
We can split into two chunks, such as [2, 1], [3, 4, 4].
However, splitting into [2, 1], [3], [4], [4] is the highest number of chunks possible.

*/

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

// # Monotone stack
// 单调栈
func maxChunksToSorted(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	stack := NewStack()
	for i := 0; i < len(arr); i++ {
		if stack.Empty() || arr[i] >= stack.Top().(int) {
			stack.Push(arr[i])
		} else {
			// 这里我们单调栈的元素个数实际上是遍历到当前数字之前可以拆分成的块儿的个数，我们遇到一个大于栈顶的元素，就将其压入栈，
			// suppose其是一个新块儿的开头，但是一旦后面遇到小的数字了，我们就要反过来检查前面的数字，有可能我们之前认为的可以拆分成块儿的地方，现在就不能拆了
			cur := stack.Top().(int)
			stack.Pop()
			for !stack.Empty() && stack.Top().(int) > arr[i] {
				stack.Pop()
			}

			stack.Push(cur)
		}
	}

	return stack.Len()
}
