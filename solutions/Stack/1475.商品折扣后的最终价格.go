/*
 * @lc app=leetcode.cn id=1475 lang=golang
 *
 * [1475] 商品折扣后的最终价格
 */

// @lc code=start
func finalPrices(prices []int) []int {
	var stack []int // 单调递增栈，栈内保存的是prices数组中的元素索引
	for i, price := range prices {
		for len(stack) > 0 && prices[stack[len(stack)-1]] >= price {
			prices[stack[len(stack)-1]] -= price
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return prices
}
// @lc code=end

