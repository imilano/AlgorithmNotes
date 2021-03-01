package leet_42

/*
	Given n non-negative integers representing an elevation map where the width of each bar is 1,
	compute how much water it can trap after raining.
*/

// stack
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

// tool function

func count(height []int, left, right int) int {
	width := right - left - 1
	area := height[left] * width
	for i := left + 1; i < right; i++ {
		area -= height[i]
	}

	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// ---------------------------------------------------------------------
// dynamic programming use two array to reduce time complexity, this is a quite clever solution
// time complexity:O(n)
// space complexity: O(n)
func trapWith2Scan(height []int) int {
	var area int
	length := len(height)
	if length <= 1 {
		return area
	}

	leftMax, rightMax := make([]int, length), make([]int, length)
	// get max from left to right
	leftMax[0] = height[0]
	for i := 1; i < length; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	// get max from right to left
	rightMax[length-1] = height[length-1]
	for i := length - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}

	// take care, here we use min, not max
	for i := 0; i < length; i++ {
		area += min(leftMax[i], rightMax[i]) - height[i]
	}

	return area
}

//---------------------------------------------
// use two pointer
// time complexity: O(N)
// space complexity: O(1)
func trapWith2Pointer(height []int) int {
	var res int
	if len(height) <= 1 {
		return res
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := height[0], height[len(height)-1]

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}

			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}

			right--
		}
	}

	return res
}

// ------------------------------------------------
// Monotone Stack
// 使用单调递减栈的思想。本处的目的在于形成一个水坑，假设当前元素为e，如果当前元素大于栈顶元素，那么我们可能会形成一个水坑：如果此时栈里面有两个元素，就可以形成一个水坑，
// 如果栈里面只有一个元素，就无法形成水坑，此时继续扫描。如果当前元素小于栈顶元素或者栈为空，那么将当前元素压入栈中。
// 需要注意的是，我们并不是直接把高度压入栈，为了方便计算宽度，我们把当前的坐标压入栈（因而实际的栈中的元素是逐渐递增的）。
func trapWithMonotoneStack(height []int) int {
	var res int
	length := len(height)
	if length == 0 {
		return res
	}

	stack := NewStack()
	var index int
	for index < length {
		// 注意等号
		if stack.Empty() || height[index] <= height[stack.Top().(int)] {
			stack.Push(index)
			index++
		} else {
			bottom := stack.Top().(int)
			stack.Pop()
			if stack.Empty() {
				continue
			}
			border := stack.Top().(int)
			res += (min(height[index], height[border]) - height[bottom]) * (index - border - 1)
		}
	}

	return res
}

// ------------------------------------------------
// Original wrong solution
func trap(height []int) int {
	var area int
	length := len(height)
	if length <= 1 {
		return area
	}

	maxHeight := 0
	for i := range height {
		if height[i] > height[maxHeight] {
			maxHeight = i
		}
	}

	lmin, lmax := maxHeight, maxHeight
	for lmin >= 0 {
		if lmin-1 >= 0 && !(height[lmin] >= height[lmin-1] && height[lmin] >= height[lmin+1]) {
			lmin--
			continue
		}
		area += count(height, lmin, lmax)
		lmax = lmin
		lmin--
	}

	rmin, rmax := maxHeight, maxHeight
	for rmax < length {
		if rmax+1 < length && height[rmax] >= height[rmax+1] && height[rmax] >= height[rmax-1] {
			rmax++
			continue
		}

		area += count(height, rmin, rmax)
		rmin = rmax
		rmax++
	}

	return area
}
