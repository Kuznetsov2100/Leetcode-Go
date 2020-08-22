/*
 * @lc app=leetcode.cn id=788 lang=golang
 *
 * [788] 旋转数字
 */

// @lc code=start
func rotatedDigits(N int) int {
	var res int
	for i := 1; i <= N; i++ {
		if isGood(i, false) {
			res++
		}
	}
	return res
}

func isGood(n int, flag bool) bool {
	if n == 0 {
		return flag
	}
	d := n % 10
	if d == 3 || d == 4 || d == 7 {
		return false
	}
	if d == 0 || d == 1 || d == 8 {
		return isGood(n/10, flag)
	}
	return isGood(n/10, true)
}


// @lc code=end

