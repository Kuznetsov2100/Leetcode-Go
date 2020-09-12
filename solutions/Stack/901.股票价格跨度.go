/*
 * @lc app=leetcode.cn id=901 lang=golang
 *
 * [901] 股票价格跨度
 */

// @lc code=start
type StockSpanner struct {
	prices []int // 单调递减栈，栈中保存的是price
	span []int // 常规栈， 栈中保存的是prices中每个股票的价格跨度
	
}


func Constructor() StockSpanner {
	return StockSpanner{}
}


func (this *StockSpanner) Next(price int) int {
	res := 1
	for len(this.prices) > 0 && this.prices[len(this.prices)-1] <= price {
		this.prices = this.prices[:len(this.prices)-1]
		res += this.span[len(this.span)-1]
		this.span = this.span[:len(this.span)-1]
	}
	this.prices = append(this.prices, price)
	this.span = append(this.span, res)
	return res
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
// @lc code=end

