/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
func singleNumber(nums []int) int {
	// 真值表 --> 逻辑表达式
	X, Y := 0, 0
	for _, Z := range nums {
		Y = Y ^ Z & ^X
		X = X ^ Z & ^Y
	}
	return Y
}
// @lc code=end

