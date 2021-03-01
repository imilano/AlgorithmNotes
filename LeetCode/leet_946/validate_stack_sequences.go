package leet_946

import "lightsinger.life/algorithm/leetcode/utils"

/*
Given two sequences pushed and popped with distinct values, return true if and only if this could have been the result of a sequence of push and pop operations on an initially empty stack.

Example 1:

Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
Output: true
Explanation: We might do the following sequence:
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
Example 2:

Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
Output: false
Explanation: 1 cannot be popped before 2.
Note:

0 <= pushed.length == popped.length <= 1000
0 <= pushed[i], popped[i] < 1000
pushed is a permutation of popped.
pushed and popped have distinct values.
*/

// 使用一个栈模拟压入弹出顺序，使用一个变量i来记录弹出序列的当前位置。将pushed中的元素依次压入栈，此时看弹出序列的当前数字是否和栈顶元素相同，
//如果相同，则弹出栈顶元素，并且i自增1，若下一个栈顶元素还跟新位置上的数字相同，还要进行相同的操作。最后遍历完push序列后，如果栈为空，则返回true，
//否则返回false
func validateStackSequences(pushed []int, popped []int) bool {
	var idx int
	stack := utils.NewStack()

	for _, v := range pushed {
		stack.Push(v)
		for !stack.Empty() && stack.Top().(int) == popped[idx] {
			stack.Pop()
			idx++
		}
	}

	return stack.Empty()
}