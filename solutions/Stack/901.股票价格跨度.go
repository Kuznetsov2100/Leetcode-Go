/*
 * @lc app=leetcode.cn id=901 lang=golang
 *
 * [901] 股票价格跨度
 */

// @lc code=start
type pair [2]int
type StockSpanner struct {
	prices []pair // 单调递减栈，栈内元素为pair， pair[0]为价格，pair[1]对应是价格跨度
}


func Constructor() StockSpanner {
	return StockSpanner{}
}


func (this *StockSpanner) Next(price int) int {
	res := 1 // 价格跨度包含当天，所以初始化为1
	for len(this.prices) > 0 && this.prices[len(this.prices)-1][0] <= price {
		res += this.prices[len(this.prices)-1][1] // pair[1]为跨度
		this.prices = this.prices[:len(this.prices)-1]
		
	}
	this.prices = append(this.prices, pair{price, res})
	return res
}



/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
// @lc code=end

