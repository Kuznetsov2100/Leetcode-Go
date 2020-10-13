/*
 * @lc app=leetcode.cn id=1154 lang=golang
 *
 * [1154] 一年中的第几天
 */

// @lc code=start
func dayOfYear(date string) int {
	t, _ := time.Parse("2006-01-02", date)
	return t.YearDay()
}
// @lc code=end

