/*
 * @lc app=leetcode.cn id=1221 lang=golang
 *
 * [1221] 分割平衡字符串
 */

// @lc code=start
func balancedStringSplit(s string) int {
	var res int
	countL, countR := 0, 0
	for i := range s {
		if s[i] == 'L' {
			countL++
		} else {
			countR++
		}
		if countL == countR {
			res++		
		}
	}
	return res
}
// @lc code=end

