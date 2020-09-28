/*
 * @lc app=leetcode.cn id=1374 lang=golang
 *
 * [1374] 生成每种字符都是奇数个的字符串
 */

// @lc code=start
func generateTheString(n int) string {
	if n % 2 == 0 {
		return helper(n)
	}
	return helper(n-1) + "c"
}

func helper(n int) string {
	if n == 0 {
		return ""
	}
	return "a" + strings.Repeat("b", n-1)
}
// @lc code=end

