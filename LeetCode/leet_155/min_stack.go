package leet_155

/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.
 */


// -----------------------------
// 使用双栈
type MinStack struct {
	stack []int
	minStack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:nil,
		minStack: nil,
	}
}


func (this *MinStack) Push(x int)  {
	length := len(this.minStack)
	if length == 0 {
		this.minStack = append(this.minStack,x)
	}
	this.stack = append(this.stack,x)

	top := this.minStack[length-1]
	if x < top {
		this.minStack = append(this.minStack,x)
	}

	//length := len(this.stack)
	//if length == 0 {
	//	this.mu.Lock()
	//	this.stack = append(this.stack,x)
	//	this.mu.Unlock()
	//	return
	//}
	//
	//var tmp []int
	//var minIndex int
	//for minIndex := length-1; minIndex >=0;minIndex-- {
	//	if this.stack[minIndex] > x {
	//		break
	//	}
	//	tmp = append(tmp,this.stack[minIndex])
	//}
	//
	//if minIndex == -1 {
	//	this.mu.Lock()
	//	this.stack = append([]int{x},this.stack...)
	//	this.mu.Unlock()
	//	return
	//} else if minIndex == length-1 {
	//	this.mu.Lock()
	//	this.stack = append(this.stack,x)
	//	this.mu.Unlock()
	//	return
	//} else {
	//	this.mu.Lock()
	//	this.stack[minIndex+1] = x
	//	this.stack = this.stack[:minIndex+2]
	//	this.stack = append(this.stack,tmp...)
	//	this.mu.Unlock()
	//	return
	//}
}


func (this *MinStack) Pop()  {
	len1,len2 := len(this.stack),len(this.minStack)
	t1,t2 := this.stack[len1-1],this.stack[len2-1]

	if t1 == t2 {
		this.minStack = this.minStack[:len2-1]
	}
	this.stack = this.stack[:len1-1]
	//if len(this.stack) == 0 {
	//	return
	//}
	//this.mu.Lock()
	//this.stack = this.stack[:len(this.stack)-1]
	//this.mu.Unlock()
}


func (this *MinStack) Top() int {
	length := len(this.stack)
	return this.stack[length-1]
	//return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	length := len(this.minStack)
	return this.minStack[length-1]
	//return this.stack[len(this.stack)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */