/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
func isPowerOfFour(num int) bool {
	if num == 0 {
		return false
	}
	var count0 int
	copynum := num
	for num != 0 {
		if num % 2 == 0 {
			count0++
		}
		num = num / 2
	}
	if count0 % 2 != 0 {
		return false
	}
	copynum = copynum & (copynum-1)
	if copynum == 0 {
		return true
	}
	return false
}
// @lc code=end

