/*
 * @lc app=leetcode.cn id=796 lang=golang
 *
 * [796] 旋转字符串
 */

// @lc code=start
func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	// B+B 包含了所有可通过旋转操作从B得到的字符串
	return strings.Contains(B+B, A)
}
// @lc code=end

