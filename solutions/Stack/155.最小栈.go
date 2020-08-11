/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	item []int
	minvalue []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{minvalue: []int{(1<<63-1)}}
}


func (this *MinStack) Push(x int)  {
	this.item = append(this.item, x)
	if x <= this.minvalue[len(this.minvalue)-1] {
		this.minvalue = append(this.minvalue, x)
	}
}


func (this *MinStack) Pop()  {
	element := this.item[len(this.item)-1]
	this.item = this.item[:len(this.item)-1]
	if element <= this.minvalue[len(this.minvalue)-1] {
		this.minvalue = this.minvalue[:len(this.minvalue)-1]
	}
}


func (this *MinStack) Top() int {
	return this.item[len(this.item)-1]
}


func (this *MinStack) GetMin() int {
	return this.minvalue[len(this.minvalue)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

