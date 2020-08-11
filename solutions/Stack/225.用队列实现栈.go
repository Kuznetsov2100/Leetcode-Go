/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
type MyStack struct {
	item []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.item = append(this.item,x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	element := this.item[len(this.item)-1]
	this.item = this.item[:len(this.item)-1]
	return element
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.item[len(this.item)-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.item) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

