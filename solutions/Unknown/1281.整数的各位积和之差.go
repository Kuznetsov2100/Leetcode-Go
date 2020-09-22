/*
 * @lc app=leetcode.cn id=1281 lang=golang
 *
 * [1281] 整数的各位积和之差
 */

// @lc code=start
func subtractProductAndSum(n int) int {
	sum, product := 0, 1
	for n > 0 {
		digit := n % 10
		sum += digit
		product *= digit
		n /= 10
	}
	return product-sum
}
// @lc code=end

