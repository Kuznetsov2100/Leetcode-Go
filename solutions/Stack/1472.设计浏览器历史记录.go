/*
 * @lc app=leetcode.cn id=1472 lang=golang
 *
 * [1472] 设计浏览器历史记录
 */

// @lc code=start
type BrowserHistory struct {
	stack []string
	cur int
	N int // stack length
}


func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{stack: []string{homepage}, cur: 0, N: 1}
}


func (this *BrowserHistory) Visit(url string)  {
	
	if this.cur < this.N-1 {
		this.stack = this.stack[:this.cur+1]
		this.N = this.cur+1
	}
	this.stack = append(this.stack, url)
	this.cur++
	this.N++
}


func (this *BrowserHistory) Back(steps int) string {
	this.cur -= steps
	if this.cur < 0 {
		this.cur = 0
	}
	return this.stack[this.cur]
}


func (this *BrowserHistory) Forward(steps int) string {
	this.cur += steps
	if this.cur >= this.N {
		this.cur = this.N-1
	}
	return this.stack[this.cur]
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
// @lc code=end

