/*
 * @lc app=leetcode.cn id=1446 lang=golang
 *
 * [1446] 连续字符
 */

// @lc code=start
func maxPower(s string) int {
	res, count := 1, 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			if count > res {
				res = count
			}
			count = 1
			continue
		} 
		count++
	}
	if count > res {
		res = count
	}
	return res
}
// @lc code=end

