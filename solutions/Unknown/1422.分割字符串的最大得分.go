/*
 * @lc app=leetcode.cn id=1422 lang=golang
 *
 * [1422] 分割字符串的最大得分
 */

// @lc code=start
func maxScore(s string) int {
	var res int
	count1 := strings.Count(s, "1")
	count0 := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '1' {
			count1--
		} else {
			count0++
		}
		res = max(res, count0+count1)
	}
	return res
}



func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

