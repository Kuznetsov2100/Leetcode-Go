/*
 * @lc app=leetcode.cn id=1464 lang=golang
 *
 * [1464] 数组中两元素的最大乘积
 */

// @lc code=start
func maxProduct(nums []int) int {
	first, second := math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num > first {
			first, second = num, first
		} else if num > second {
			second = num
		}
	}
	return (first-1)*(second-1)
}
// @lc code=end

