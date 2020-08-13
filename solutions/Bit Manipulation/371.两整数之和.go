/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 */

// @lc code=start
func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	var lower, carrier int
	for {
		lower = a ^ b // 计算低位
		carrier = a & b // 计算进位
		if carrier == 0 { // 进位为0，计算结束
			break
		}
		a = lower
		b = carrier << 1
	}
	return lower
}
// @lc code=end

