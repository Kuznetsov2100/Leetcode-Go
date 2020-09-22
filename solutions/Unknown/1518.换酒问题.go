/*
 * @lc app=leetcode.cn id=1518 lang=golang
 *
 * [1518] 换酒问题
 */

// @lc code=start
func numWaterBottles(numBottles int, numExchange int) int {
	res := numBottles
	for numBottles >= numExchange {
		exchangedBottles := numBottles/numExchange
		res += exchangedBottles
		numBottles = exchangedBottles + numBottles % numExchange
	}
	return res
}
// @lc code=end

