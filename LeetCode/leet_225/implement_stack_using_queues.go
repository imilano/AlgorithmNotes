package leet_225

/*
Implement a last in first out (LIFO) stack using only two queues. The implemented stack should support all the functions of a normal queue (push, top, pop, and empty).

Implement the MyStack class:

void push(int x) Pushes element x to the top of the stack.
int pop() Removes the element on the top of the stack and returns it.
int top() Returns the element on the top of the stack.
boolean empty() Returns true if the stack is empty, false otherwise.
Notes:

You must use only standard operations of a queue, which means only push to back, peek/pop from front, size, and is empty operations are valid.
Depending on your language, the queue may not be supported natively. You may simulate a queue using a list or deque (double-ended queue), as long as you use only a queue's standard operations.


Example 1:

Input
["MyStack", "push", "push", "top", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 2, 2, false]

Explanation
MyStack myStack = new MyStack();
myStack.push(1);
myStack.push(2);
myStack.top(); // return 2
myStack.pop(); // return 2
myStack.empty(); // return False


Constraints:

1 <= x <= 9
At most 100 calls will be made to push, pop, top, and empty.
All the calls to pop and top are valid.


Follow-up: Can you implement the stack such that each operation is amortized O(1) time complexity? In other words, performing n operations will take overall O(n) time even if one of those operations may take longer. You can use more than two queues.
*/


type MyStack struct {
	que []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{que: nil}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.que = append(this.que,x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	n := len(this.que)
	res := this.que[n-1]
	this.que = this.que[:n-1]
	return res
}


/** Get the top element. */
func (this *MyStack) Top() int {
	if this.Empty() {
		return -1
	}

	return this.que[len(this.que)-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.que) == 0 {
		return true
	}

	return false
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
